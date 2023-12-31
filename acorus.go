package acorus

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"net"
	"strconv"
	"sync/atomic"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/prometheus/client_golang/prometheus"

	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/service/common/httputil"
	"github.com/cornerstone-labs/acorus/synchronizer"
	"github.com/cornerstone-labs/acorus/synchronizer/node"
)

type Acorus struct {
	DB              *database.DB
	ethClient       map[uint64]node.EthClient
	apiServer       *httputil.HTTPServer
	metricsServer   *httputil.HTTPServer
	metricsRegistry *prometheus.Registry
	Synchronizer    map[uint64]*synchronizer.Synchronizer
	shutdown        context.CancelCauseFunc
	stopped         atomic.Bool
	chainIdList     []uint64
}

func NewAcorus(ctx context.Context, cfg *config.Config, shutdown context.CancelCauseFunc) (*Acorus, error) {
	out := &Acorus{
		shutdown: shutdown,
	}
	if err := out.initFromConfig(ctx, cfg); err != nil {
		return nil, errors.Join(err, out.Stop(ctx))
	}
	return out, nil
}

func (as *Acorus) Start(ctx context.Context) error {
	for i := range as.chainIdList {
		if err := as.Synchronizer[as.chainIdList[i]].Start(); err != nil {
			return fmt.Errorf("failed to start L1 Sync: %w", err)
		}
	}
	return nil
}

func (as *Acorus) Stop(ctx context.Context) error {
	var result error
	for i := range as.chainIdList {
		if as.Synchronizer[as.chainIdList[i]] != nil {
			if err := as.Synchronizer[as.chainIdList[i]].Close(); err != nil {
				result = errors.Join(result, fmt.Errorf("failed to close L2 Synchronizer: %w", err))
			}
		}
		if as.ethClient[as.chainIdList[i]] != nil {
			as.ethClient[as.chainIdList[i]].Close()
		}
	}

	if as.apiServer != nil {
		if err := as.apiServer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close acorus API server: %w", err))
		}
	}

	if as.DB != nil {
		if err := as.DB.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close DB: %w", err))
		}
	}

	if as.metricsServer != nil {
		if err := as.metricsServer.Close(); err != nil {
			result = errors.Join(result, fmt.Errorf("failed to close metrics server: %w", err))
		}
	}

	as.stopped.Store(true)

	log.Info("acorus stopped")

	return result
}

func (as *Acorus) Stopped() bool {
	return as.stopped.Load()
}

func (as *Acorus) initFromConfig(ctx context.Context, cfg *config.Config) error {
	if err := as.initRPCClients(ctx, cfg); err != nil {
		return fmt.Errorf("failed to start RPC clients: %w", err)
	}
	if err := as.initDB(ctx, cfg.MasterDb); err != nil {
		return fmt.Errorf("failed to init DB: %w", err)
	}
	if err := as.initSynchronizer(cfg); err != nil {
		return fmt.Errorf("failed to init L1 Sync: %w", err)
	}

	return nil
}

func (as *Acorus) initRPCClients(ctx context.Context, conf *config.Config) error {
	for i := range conf.RPCs {
		rpc := conf.RPCs[i]
		ethClient, err := node.DialEthClient(ctx, rpc.RpcUrl)
		if err != nil {
			return fmt.Errorf("failed to dial L1 client: %w", err)
		}
		as.ethClient[rpc.ChainId] = ethClient
		as.chainIdList = append(as.chainIdList, rpc.ChainId)
	}
	return nil
}

func (as *Acorus) initDB(ctx context.Context, cfg config.Database) error {
	db, err := database.NewDB(ctx, cfg)
	if err != nil {
		return fmt.Errorf("failed to connect to database: %w", err)
	}
	as.DB = db
	return nil
}

func (as *Acorus) initSynchronizer(config *config.Config) error {
	for i := range config.RPCs {
		rpcItem := config.RPCs[i]
		cfg := synchronizer.Config{
			LoopIntervalMsec:  5,
			HeaderBufferSize:  500,
			ConfirmationDepth: big.NewInt(int64(1)),
			StartHeight:       big.NewInt(int64(rpcItem.StartBlock)),
			ChainId:           uint(rpcItem.ChainId),
		}
		ethClient, err := node.DialEthClient(context.Background(), rpcItem.RpcUrl)
		if err != nil {
			return fmt.Errorf("failed to dial L1 client: %w", err)
		}
		synchronizerTemp, err := synchronizer.NewSynchronizer(&cfg, as.DB, ethClient, as.shutdown)
		if err != nil {
			return err
		}
		as.Synchronizer[rpcItem.ChainId] = synchronizerTemp
	}
	return nil
}

func (as *Acorus) initBridgeProcessor(config config.Config) error {
	return nil
}

func (as *Acorus) initBusinessProcessor(cfg config.Config) error {
	return nil
}

func (as *Acorus) startHttpServer(ctx context.Context, cfg config.Server) error {
	log.Debug("starting http server...", "port", cfg.Port)
	r := chi.NewRouter()
	r.Use(middleware.Heartbeat("/healthz"))
	addr := net.JoinHostPort(cfg.Host, strconv.Itoa(cfg.Port))
	srv, err := httputil.StartHTTPServer(addr, r)
	if err != nil {
		return fmt.Errorf("http server failed to start: %w", err)
	}
	as.apiServer = srv
	log.Info("http server started", "addr", srv.Addr())
	return nil
}

func (as *Acorus) startMetricsServer(ctx context.Context, cfg config.Server) error {
	return nil
}
