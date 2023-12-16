package acorus

import (
	"context"
	"fmt"
	"github.com/cornerstone-labs/acorus/event"
	"github.com/cornerstone-labs/acorus/event/processor"
	"math/big"
	"net"
	"runtime/debug"
	"strconv"
	"sync"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-redis/redis/v8"
)

var ZeroAddr = common.Address{}

type Acorus struct {
	log         log.Logger
	db          *database.DB
	redis       *redis.Client
	httpConfig  config.ServerConfig
	L1Sync      *synchronizer.L1Sync
	L2Sync      *synchronizer.L2Sync
	Chain       config.ChainConfig
	ChainBridge string
}

func NewAcorus(
	log log.Logger,
	db *database.DB,
	redis *redis.Client,
	config *config.Config,
	chainBridge string,
) (*Acorus, error) {
	l1EthClient, err := node.DialEthClient(config.Chain.L1RPC)
	if err != nil {
		log.Error("Dail eth client err", "err", err)
		return nil, err
	}
	l1Cfg := synchronizer.Config{
		LoopIntervalMsec:  config.Chain.L1PollingInterval,
		HeaderBufferSize:  config.Chain.L1HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(config.Chain.L1ConfirmationDepth)),
		StartHeight:       big.NewInt(int64(config.Chain.L1StartHeight)),
	}

	//l1Contracts, l2Contracts, _, err := nil,nil,nil,nil;ChainContractsSelect(chainBridge, config)
	if err != nil {
		return nil, err
	}
	log.Info("chain select success", "chain", chainBridge)

	l1Syncer, err := synchronizer.NewL1Sync(l1Cfg, log, db, l1EthClient, config.Chain.L1Contracts, config.Chain)
	if err != nil {
		return nil, err
	}
	log.Info("l1 syncer new success", "rpc", config.Chain.L1RPC)

	l2EthClient, err := node.DialEthClient(config.Chain.L2RPC)
	if err != nil {
		return nil, err
	}
	l2Cfg := synchronizer.Config{
		LoopIntervalMsec:  config.Chain.L2PollingInterval,
		HeaderBufferSize:  config.Chain.L2HeaderBufferSize,
		ConfirmationDepth: big.NewInt(int64(config.Chain.L2ConfirmationDepth)),
	}

	l2Syncer, err := synchronizer.NewL2Sync(l2Cfg, log, db, l2EthClient, config.Chain.L2Contracts, config.Chain)
	if err != nil {
		return nil, err
	}

	//dispatcher, err := event.NewEventDispatcher(log, db, l1Syncer, config.Chain, chainBridge, resultContracts)
	//if err != nil {
	//	return nil, err
	//}

	acorus := &Acorus{
		log:         log,
		db:          db,
		redis:       redis,
		httpConfig:  config.HTTPServer,
		L1Sync:      l1Syncer,
		L2Sync:      l2Syncer,
		Chain:       config.Chain,
		ChainBridge: chainBridge,
	}
	return acorus, nil
}

func (i *Acorus) startHttpServer(ctx context.Context) error {
	i.log.Debug("starting http server...", "port", i.httpConfig.Host)

	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))

	addr := net.JoinHostPort(i.httpConfig.Host, strconv.Itoa(i.httpConfig.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	i.log.Info("http server started", "addr", srv.Addr())
	<-ctx.Done()
	defer i.log.Info("http server stopped")
	return srv.Stop(context.Background())
}

func (i *Acorus) startMetricsServer(ctx context.Context) error {
	return nil
}

func (i *Acorus) Run(ctx context.Context) error {
	var wg sync.WaitGroup
	errCh := make(chan error, 5)

	// if any goroutine halts, we stop the entire Acorus
	processCtx, processCancel := context.WithCancel(ctx)
	runProcess := func(start func(ctx context.Context) error) {
		wg.Add(1)
		go func() {
			defer func() {
				if err := recover(); err != nil {
					i.log.Error("halting acorus on panic", "err", err)
					debug.PrintStack()
					errCh <- fmt.Errorf("panic: %v", err)
				}
				processCancel()
				wg.Done()
			}()
			errCh <- start(processCtx)
		}()
	}

	// synchronizer engine
	runProcess(i.L1Sync.Start)
	runProcess(i.L2Sync.Start)

	// event engine
	factoryParams := processor.GetEventFactoryParams{
		ChainBridge: i.ChainBridge,
		Log:         i.log,
		Db:          i.db,
		L1Syncer:    i.L1Sync,
		ChainConfig: i.Chain,
	}
	factory, factoryErr := event.GetEventFactory(factoryParams)
	if factoryErr != nil {
		return factoryErr
	}
	runProcess(factory.GetChainUnpackService(factoryParams).Start)

	// service server
	runProcess(i.startHttpServer)

	wg.Wait()

	err := <-errCh
	if err != nil {
		i.log.Error("acorus stopped", "err", err)
	} else {
		i.log.Info("acorus stopped")
	}
	return err
}

//func ChainContractsSelect(chainBridge string, config2 *config.Config) (l1Contracts []common.Address, l2Contracts []common.Address, contracts interface{}, err error) {
//	var resultContracts interface{}
//	var l1ResultContracts []common.Address
//	var l2ResultContracts []common.Address
//	if chainBridge == common2.Op {
//		if err := config2.OpContracts.L2Contracts.ForEach(func(name string, addr common.Address) error {
//			if addr == ZeroAddr {
//				log.Error("address not configured", "name", name)
//				return errors.New("all L2Contracts must be configured")
//			}
//			log.Info("configured contract", "name", name, "addr", addr)
//			l2ResultContracts = append(l2ResultContracts, addr)
//			return nil
//		}); err != nil {
//			return nil, nil, nil, err
//		}
//		if err := config2.OpContracts.L1Contracts.ForEach(func(name string, addr common.Address) error {
//			if addr == ZeroAddr && !strings.HasPrefix(name, "Legacy") {
//				log.Error("address not configured", "name", name)
//				return errors.New("all L1Contracts must be configured")
//			}
//			log.Info("configured contract", "name", name, "addr", addr)
//			l1ResultContracts = append(l1ResultContracts, addr)
//			return nil
//		}); err != nil {
//			return nil, nil, nil, err
//		}
//
//		resultContracts = config2.OpContracts
//
//	} else if chainBridge == common2.Polygon {
//		// todo: handle polygon logic
//	} else if chainBridge == common2.Scroll {
//		if err := config2.SclContracts.L2Contracts.ForEach(func(name string, addr common.Address) error {
//			if addr == ZeroAddr {
//				log.Error("address not configured", "name", name)
//				return errors.New("all L2Contracts must be configured")
//			}
//			log.Info("configured contract", "name", name, "addr", addr)
//			l2ResultContracts = append(l2ResultContracts, addr)
//			return nil
//		}); err != nil {
//			return nil, nil, nil, err
//		}
//		if err := config2.SclContracts.L1Contracts.ForEach(func(name string, addr common.Address) error {
//			if addr == ZeroAddr && !strings.HasPrefix(name, "Legacy") {
//				log.Error("address not configured", "name", name)
//				return errors.New("all L1Contracts must be configured")
//			}
//			log.Info("configured contract", "name", name, "addr", addr)
//			l1ResultContracts = append(l1ResultContracts, addr)
//			return nil
//		}); err != nil {
//			return nil, nil, nil, err
//		}
//
//		resultContracts = config2.SclContracts
//	}
//	return l1ResultContracts, l2ResultContracts, resultContracts, nil
//}
