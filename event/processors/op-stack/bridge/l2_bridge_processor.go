package bridge

import (
	"errors"
	"fmt"
	contracts2 "github.com/cornerstone-labs/acorus/event/processors/op-stack/contracts"
	"github.com/cornerstone-labs/acorus/event/processors/op-stack/mantle/op-bindings/predeploys"
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/log"

	"github.com/cornerstone-labs/acorus/config"
	"github.com/cornerstone-labs/acorus/database"
	"github.com/cornerstone-labs/acorus/database/event/l1-l2"
	"github.com/cornerstone-labs/acorus/database/worker"
)

// L2ProcessInitiatedBridgeEvents will query the database for bridge events that have been initiated between
// the specified block range. This covers every part of the multi-layered stack:
//  1. OptimismPortal
//  2. L2CrossDomainMessenger
//  3. L2StandardBridge
func L2ProcessInitiatedBridgeEvents(log log.Logger, db *database.DB, l2Contracts config.L2Contracts, fromHeight, toHeight *big.Int) error {
	// (1) L2ToL1MessagePasser
	l2ToL1MPMessagesPassed, err := contracts2.L2ToL1MessagePasserMessagePassedEvents(l2Contracts.L2ToL1MessagePasser, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(l2ToL1MPMessagesPassed) > 0 {
		log.Info("detected transaction withdrawals", "size", len(l2ToL1MPMessagesPassed))
	}
	messagesPassed := make(map[logKey]*contracts2.L2ToL1MessagePasserMessagePassed, len(l2ToL1MPMessagesPassed))
	transactionWithdrawals := make([]l1_l2.L2TransactionWithdrawal, len(l2ToL1MPMessagesPassed))
	l2ToL1s := make([]worker.L2ToL1, len(l2ToL1MPMessagesPassed))
	for i := range l2ToL1MPMessagesPassed {
		messagePassed := l2ToL1MPMessagesPassed[i]
		messagesPassed[logKey{messagePassed.Event.BlockHash, messagePassed.Event.LogIndex}] = &messagePassed
		transactionWithdrawals[i] = l1_l2.L2TransactionWithdrawal{
			WithdrawalHash:       messagePassed.WithdrawalHash,
			InitiatedL2EventGUID: messagePassed.Event.GUID,
			Nonce:                messagePassed.Nonce,
			GasLimit:             messagePassed.GasLimit,
			Tx:                   messagePassed.Tx,
		}
		l2ToL1s[i] = worker.L2ToL1{
			L2BlockNumber:             messagePassed.Event.BlockHash,
			MsgNonce:                  messagePassed.Nonce,
			L2TransactionHash:         messagePassed.Event.TransactionHash,
			L2WithdrawTransactionHash: messagePassed.WithdrawalHash,
			L1ProveTxHash:             common.Hash{},
			L1FinalizeTxHash:          common.Hash{},
			Status:                    0,
			FromAddress:               messagePassed.Tx.FromAddress,
			ETHAmount:                 messagePassed.Tx.ETHAmount,
			ERC20Amount:               messagePassed.Tx.ERC20Amount,
			GasLimit:                  messagePassed.GasLimit,
			TimeLeft:                  new(big.Int).SetUint64(0),
			ToAddress:                 messagePassed.Tx.ToAddress,
			L1TokenAddress:            common.Address{},
			L2TokenAddress:            common.Address{},
			Timestamp:                 int64(messagePassed.Event.Timestamp),
		}
	}
	if len(messagesPassed) > 0 {
		if err := db.BridgeTransactions.StoreL2TransactionWithdrawals(transactionWithdrawals); err != nil {
			return err
		}
		if err := db.L2ToL1.StoreL2ToL1Transactions(l2ToL1s); err != nil {
			return err
		}
	}

	// (2) L2CrossDomainMessenger
	crossDomainSentMessages, err := contracts2.CrossDomainMessengerSentMessageEvents("l2", l2Contracts.L2CrossDomainMessenger, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainSentMessages) > 0 {
		log.Info("detected sent messages", "size", len(crossDomainSentMessages))
	}

	sentMessages := make(map[logKey]*contracts2.CrossDomainMessengerSentMessageEvent, len(crossDomainSentMessages))
	bridgeMessages := make([]l1_l2.L2BridgeMessage, len(crossDomainSentMessages))
	for i := range crossDomainSentMessages {
		sentMessage := crossDomainSentMessages[i]
		sentMessages[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex}] = &sentMessage

		// extract the withdrawal hash from the previous MessagePassed event
		messagePassed, ok := messagesPassed[logKey{sentMessage.Event.BlockHash, sentMessage.Event.LogIndex - 1}]
		if !ok {
			log.Error("expected MessagePassedEvent preceding SentMessage", "tx_hash", sentMessage.Event.TransactionHash.String())
			return fmt.Errorf("expected MessagePassedEvent preceding SentMessage. tx_hash = %s", sentMessage.Event.TransactionHash.String())
		}

		bridgeMessages[i] = l1_l2.L2BridgeMessage{TransactionWithdrawalHash: messagePassed.WithdrawalHash, BridgeMessage: sentMessage.BridgeMessage}
	}
	if len(bridgeMessages) > 0 {
		if err := db.BridgeMessages.StoreL2BridgeMessages(bridgeMessages); err != nil {
			return err
		}
	}

	// (3) L2StandardBridge
	initiatedBridges, err := contracts2.StandardBridgeInitiatedEvents("l2", l2Contracts.L2StandardBridge, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(initiatedBridges) > 0 {
		log.Info("detected bridge withdrawals", "size", len(initiatedBridges))
	}

	bridgedTokens := make(map[common.Address]int)
	bridgeWithdrawals := make([]l1_l2.L2BridgeWithdrawal, len(initiatedBridges))
	l2ToL1s2 := make([]worker.L2ToL1, len(initiatedBridges))
	for i := range initiatedBridges {
		initiatedBridge := initiatedBridges[i]
		// extract the cross domain message hash & deposit source hash from the following events
		if initiatedBridge.Event.EventSignature.String() == predeploys.ETHWithdrawEventSignature {
			messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 5}]
			if !ok {
				log.Error("expected MessagePassed following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 6}]
			if !ok {
				log.Error("expected SentMessage following MessagePassed event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}

			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeWithdrawals[i] = l1_l2.L2BridgeWithdrawal{
				TransactionWithdrawalHash: messagePassed.WithdrawalHash,
				BridgeTransfer:            initiatedBridge.BridgeTransfer,
			}
			l2ToL1s2[i].L2WithdrawTransactionHash = messagePassed.WithdrawalHash
			l2ToL1s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l2ToL1s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
		} else if initiatedBridge.Event.EventSignature.String() == predeploys.MNTWithdrawEventSignature {
			messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Error("expected MessagePassed following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Error("expected SentMessage following MessagePassed event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}

			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeWithdrawals[i] = l1_l2.L2BridgeWithdrawal{
				TransactionWithdrawalHash: messagePassed.WithdrawalHash,
				BridgeTransfer:            initiatedBridge.BridgeTransfer,
			}
			l2ToL1s2[i].L2WithdrawTransactionHash = messagePassed.WithdrawalHash
			l2ToL1s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l2ToL1s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
		} else if initiatedBridge.Event.EventSignature.String() == predeploys.ERC20WithdrawEventSignature {
			messagePassed, ok := messagesPassed[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 1}]
			if !ok {
				log.Error("expected MessagePassed following BridgeInitiated event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected MessagePassed following BridgeInitiated event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}
			sentMessage, ok := sentMessages[logKey{initiatedBridge.Event.BlockHash, initiatedBridge.Event.LogIndex + 2}]
			if !ok {
				log.Error("expected SentMessage following MessagePassed event", "tx_hash", initiatedBridge.Event.TransactionHash.String())
				return fmt.Errorf("expected SentMessage following MessagePassed event. tx_hash = %s", initiatedBridge.Event.TransactionHash.String())
			}

			initiatedBridge.BridgeTransfer.CrossDomainMessageHash = &sentMessage.BridgeMessage.MessageHash
			bridgedTokens[initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
			bridgeWithdrawals[i] = l1_l2.L2BridgeWithdrawal{
				TransactionWithdrawalHash: messagePassed.WithdrawalHash,
				BridgeTransfer:            initiatedBridge.BridgeTransfer,
			}
			l2ToL1s2[i].L2WithdrawTransactionHash = messagePassed.WithdrawalHash
			l2ToL1s2[i].L1TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.LocalTokenAddress
			l2ToL1s2[i].L2TokenAddress = initiatedBridge.BridgeTransfer.TokenPair.RemoteTokenAddress
		}
	}

	if len(bridgeWithdrawals) > 0 {
		if err := db.BridgeTransfers.StoreL2BridgeWithdrawals(bridgeWithdrawals); err != nil {
			return err
		}
		if err := db.L2ToL1.UpdateTokenPairs(l2ToL1s2); err != nil {
			return err
		}

	}

	// a-ok!
	return nil
}

// L2ProcessFinalizedBridgeEvent will query the database for all the finalization markers for all initiated
// bridge events. This covers every part of the multi-layered stack:
//  1. L2CrossDomainMessenger (relayMessage marker)
//  2. L2StandardBridge (no-op, since this is simply a wrapper over the L2CrossDomainMEssenger)
//
// NOTE: Unlike L1, there's no L2ToL1MessagePasser stage since transaction deposits are apart of the block derivation process.
func L2ProcessFinalizedBridgeEvents(log log.Logger, db *database.DB, l2Contracts config.L2Contracts, fromHeight, toHeight *big.Int) error {
	// (1) L2CrossDomainMessenger
	crossDomainRelayedMessages, err := contracts2.CrossDomainMessengerRelayedMessageEvents("l2", l2Contracts.L2CrossDomainMessenger, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(crossDomainRelayedMessages) > 0 {
		log.Info("detected relayed messages", "size", len(crossDomainRelayedMessages))
	}

	relayedMessages := make(map[logKey]*contracts2.CrossDomainMessengerRelayedMessageEvent, len(crossDomainRelayedMessages))
	for i := range crossDomainRelayedMessages {
		relayed := crossDomainRelayedMessages[i]
		relayedMessages[logKey{BlockHash: relayed.Event.BlockHash, LogIndex: relayed.Event.LogIndex}] = &relayed
		message, err := db.BridgeMessages.L1BridgeMessage(relayed.MessageHash)
		if err != nil {
			return err
		} else if message == nil {
			log.Error("missing indexed L1CrossDomainMessenger message", "tx_hash", relayed.Event.TransactionHash.String())
			return fmt.Errorf("missing indexed L1CrossDomainMessager message. tx_hash = %s", relayed.Event.TransactionHash.String())
		}

		if err := db.BridgeMessages.MarkRelayedL1BridgeMessage(relayed.MessageHash, relayed.Event.GUID); err != nil {
			log.Error("failed to relay cross domain message", "err", err, "tx_hash", relayed.Event.TransactionHash.String())
			return err
		}
	}

	// (2) L2StandardBridge
	finalizedBridges, err := contracts2.StandardBridgeFinalizedEvents("l2", l2Contracts.L2StandardBridge, db, fromHeight, toHeight)
	if err != nil {
		return err
	}
	if len(finalizedBridges) > 0 {
		log.Info("detected finalized bridge deposits", "size", len(finalizedBridges))
	}

	finalizedTokens := make(map[common.Address]int)
	for i := range finalizedBridges {
		// Nothing actionable on the database. However, we can treat the relayed message
		// as an invariant by ensuring we can query for a deposit by the same hash
		finalizedBridge := finalizedBridges[i]
		relayedMessage, ok := relayedMessages[logKey{finalizedBridge.Event.BlockHash, finalizedBridge.Event.LogIndex + 1}]
		if !ok {
			log.Error("expected RelayedMessage following BridgeFinalized event", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return fmt.Errorf("expected RelayedMessage following BridgeFinalized event. tx_hash = %s", finalizedBridge.Event.TransactionHash.String())
		}

		// Since the message hash is computed from the relayed message, this ensures the withdrawal fields must match
		deposit, err := db.BridgeTransfers.L1BridgeDepositWithFilter(l1_l2.BridgeTransfer{CrossDomainMessageHash: &relayedMessage.MessageHash})
		if err != nil {
			return err
		} else if deposit == nil {
			log.Error("missing L1StandardBridge deposit on L2 finalization", "tx_hash", finalizedBridge.Event.TransactionHash.String())
			return errors.New("missing L1StandardBridge deposit on L2 finalization")
		}

		finalizedTokens[finalizedBridge.BridgeTransfer.TokenPair.LocalTokenAddress]++
	}

	// a-ok!
	return nil
}