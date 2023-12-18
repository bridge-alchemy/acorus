// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"
	"github.com/mantlenetworkio/lithosphere/event/op-bindings/solc"
)

const L2ToL1MessagePasserStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"sentMessages\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_bytes32,t_bool)\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2ToL1MessagePasser.sol:L2ToL1MessagePasser\",\"label\":\"msgNonce\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_uint240\"}],\"types\":{\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_bytes32\":{\"encoding\":\"inplace\",\"label\":\"bytes32\",\"numberOfBytes\":\"32\"},\"t_mapping(t_bytes32,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(bytes32 =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_bytes32\",\"value\":\"t_bool\"},\"t_uint240\":{\"encoding\":\"inplace\",\"label\":\"uint240\",\"numberOfBytes\":\"30\"}}}"

var L2ToL1MessagePasserStorageLayout = new(solc.StorageLayout)

var L2ToL1MessagePasserDeployedBin = "0x6080604052600436106100695760003560e01c806354fd4d501161004357806354fd4d50146100e957806382e3702d1461010b578063ecc704281461014b57600080fd5b806339fd0090146100945780633f827a5a146100a757806344df8e70146100d457600080fd5b3661008f5761008d600033620186a0604051806020016040528060008152506101b0565b005b600080fd5b61008d6100a23660046106dd565b6101b0565b3480156100b357600080fd5b506100bc600181565b60405161ffff90911681526020015b60405180910390f35b3480156100e057600080fd5b5061008d610409565b3480156100f557600080fd5b506100fe610441565b6040516100cb9190610865565b34801561011757600080fd5b5061013b61012636600461087f565b60006020819052908152604090205460ff1681565b60405190151581526020016100cb565b34801561015757600080fd5b506101a26001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b6040519081526020016100cb565b831561023b576040517f9dc29fac0000000000000000000000000000000000000000000000000000000081523360048201526024810185905273deaddeaddeaddeaddeaddeaddeaddeaddead111190639dc29fac90604401600060405180830381600087803b15801561022257600080fd5b505af1158015610236573d6000803e3d6000fd5b505050505b60006102d86040518060e001604052806102956001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b815233602082015273ffffffffffffffffffffffffffffffffffffffff871660408201523460608201526080810188905260a0810186905260c0018490526104e4565b600081815260208190526040902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00166001179055905073ffffffffffffffffffffffffffffffffffffffff8416336103736001547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff167e010000000000000000000000000000000000000000000000000000000000001790565b7f5da382596b838a63b4248e533d8e399b3b0f13ba6c6679f670489d44716cb17334898888886040516103aa959493929190610898565b60405180910390a45050600180547dffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8082168301167fffff000000000000000000000000000000000000000000000000000000000000909116179055505050565b4761041381610536565b60405181907f7967de617a5ac1cc7eba2d6f37570a0135afa950d8bb77cdd35f0d0b4e85a16f90600090a250565b606061046c7f0000000000000000000000000000000000000000000000000000000000000000610565565b6104957f0000000000000000000000000000000000000000000000000000000000000000610565565b6104be7f0000000000000000000000000000000000000000000000000000000000000000610565565b6040516020016104d0939291906108cf565b604051602081830303815290604052905090565b80516020808301516040808501516060860151608087015160a088015160c08901519451600098610519989097969101610945565b604051602081830303815290604052805190602001209050919050565b80604051610543906106a2565b6040518091039082f0905080158015610560573d6000803e3d6000fd5b505050565b6060816000036105a857505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156105d257806105bc816109d2565b91506105cb9050600a83610a39565b91506105ac565b60008167ffffffffffffffff8111156105ed576105ed6106ae565b6040519080825280601f01601f191660200182016040528015610617576020820181803683370190505b5090505b841561069a5761062c600183610a4d565b9150610639600a86610a64565b610644906030610a78565b60f81b81838151811061065957610659610a90565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610693600a86610a39565b945061061b565b949350505050565b600880610ac083390190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600080600080608085870312156106f357600080fd5b84359350602085013573ffffffffffffffffffffffffffffffffffffffff8116811461071e57600080fd5b925060408501359150606085013567ffffffffffffffff8082111561074257600080fd5b818701915087601f83011261075657600080fd5b813581811115610768576107686106ae565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156107ae576107ae6106ae565b816040528281528a60208487010111156107c757600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b60005b838110156108065781810151838201526020016107ee565b83811115610815576000848401525b50505050565b600081518084526108338160208601602086016107eb565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000610878602083018461081b565b9392505050565b60006020828403121561089157600080fd5b5035919050565b85815284602082015283604082015260a0606082015260006108bd60a083018561081b565b90508260808301529695505050505050565b600084516108e18184602089016107eb565b80830190507f2e00000000000000000000000000000000000000000000000000000000000000808252855161091d816001850160208a016107eb565b600192019182015283516109388160028401602088016107eb565b0160020195945050505050565b878152600073ffffffffffffffffffffffffffffffffffffffff80891660208401528088166040840152508560608301528460808301528360a083015260e060c083015261099660e083018461081b565b9998505050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a0357610a036109a3565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610a4857610a48610a0a565b500490565b600082821015610a5f57610a5f6109a3565b500390565b600082610a7357610a73610a0a565b500690565b60008219821115610a8b57610a8b6109a3565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe608060405230fffea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2ToL1MessagePasserStorageLayoutJSON), L2ToL1MessagePasserStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2ToL1MessagePasser"] = L2ToL1MessagePasserStorageLayout
	deployedBytecodes["L2ToL1MessagePasser"] = L2ToL1MessagePasserDeployedBin
}
