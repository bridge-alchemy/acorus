// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"
	"github.com/cornerstone-labs/acorus/event/op-bindings/solc"
)

const DeployerWhitelistStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/legacy/DeployerWhitelist.sol:DeployerWhitelist\",\"label\":\"owner\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/legacy/DeployerWhitelist.sol:DeployerWhitelist\",\"label\":\"whitelist\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var DeployerWhitelistStorageLayout = new(solc.StorageLayout)

var DeployerWhitelistDeployedBin = "0x608060405234801561001057600080fd5b506004361061007d5760003560e01c80638da5cb5b1161005b5780638da5cb5b146100c85780639b19251a1461010d578063b1540a0114610140578063bdc7b54f1461015357600080fd5b806308fd63221461008257806313af40351461009757806354fd4d50146100aa575b600080fd5b61009561009036600461088a565b61015b565b005b6100956100a53660046108c6565b6102bb565b6100b26104ec565b6040516100bf9190610918565b60405180910390f35b6000546100e89073ffffffffffffffffffffffffffffffffffffffff1681565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016100bf565b61013061011b3660046108c6565b60016020526000908152604090205460ff1681565b60405190151581526020016100bf565b61013061014e3660046108c6565b61058f565b6100956105e0565b60005473ffffffffffffffffffffffffffffffffffffffff16331461022d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f4465706c6f79657257686974656c6973743a2066756e6374696f6e2063616e2060448201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460648201527f68697320636f6e74726163740000000000000000000000000000000000000000608482015260a4015b60405180910390fd5b73ffffffffffffffffffffffffffffffffffffffff821660008181526001602090815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00168515159081179091558251938452908301527f8daaf060c3306c38e068a75c054bf96ecd85a3db1252712c4d93632744c42e0d910160405180910390a15050565b60005473ffffffffffffffffffffffffffffffffffffffff163314610388576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f4465706c6f79657257686974656c6973743a2066756e6374696f6e2063616e2060448201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460648201527f68697320636f6e74726163740000000000000000000000000000000000000000608482015260a401610224565b73ffffffffffffffffffffffffffffffffffffffff8116610451576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604d60248201527f4465706c6f79657257686974656c6973743a2063616e206f6e6c79206265206460448201527f697361626c65642076696120656e61626c65417262697472617279436f6e747260648201527f6163744465706c6f796d656e7400000000000000000000000000000000000000608482015260a401610224565b6000546040805173ffffffffffffffffffffffffffffffffffffffff928316815291831660208301527fb532073b38c83145e3e5135377a08bf9aab55bc0fd7c1179cd4fb995d2a5159c910160405180910390a1600080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff92909216919091179055565b60606105177f0000000000000000000000000000000000000000000000000000000000000000610724565b6105407f0000000000000000000000000000000000000000000000000000000000000000610724565b6105697f0000000000000000000000000000000000000000000000000000000000000000610724565b60405160200161057b93929190610969565b604051602081830303815290604052905090565b6000805473ffffffffffffffffffffffffffffffffffffffff1615806105da575073ffffffffffffffffffffffffffffffffffffffff821660009081526001602052604090205460ff165b92915050565b60005473ffffffffffffffffffffffffffffffffffffffff1633146106ad576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604c60248201527f4465706c6f79657257686974656c6973743a2066756e6374696f6e2063616e2060448201527f6f6e6c792062652063616c6c656420627920746865206f776e6572206f66207460648201527f68697320636f6e74726163740000000000000000000000000000000000000000608482015260a401610224565b60005460405173ffffffffffffffffffffffffffffffffffffffff90911681527fc0e106cf568e50698fdbde1eff56f5a5c966cc7958e37e276918e9e4ccdf8cd49060200160405180910390a1600080547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60608160000361076757505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b8115610791578061077b81610a0e565b915061078a9050600a83610a75565b915061076b565b60008167ffffffffffffffff8111156107ac576107ac610a89565b6040519080825280601f01601f1916602001820160405280156107d6576020820181803683370190505b5090505b8415610859576107eb600183610ab8565b91506107f8600a86610acf565b610803906030610ae3565b60f81b81838151811061081857610818610afb565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350610852600a86610a75565b94506107da565b949350505050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461088557600080fd5b919050565b6000806040838503121561089d57600080fd5b6108a683610861565b9150602083013580151581146108bb57600080fd5b809150509250929050565b6000602082840312156108d857600080fd5b6108e182610861565b9392505050565b60005b838110156109035781810151838201526020016108eb565b83811115610912576000848401525b50505050565b60208152600082518060208401526109378160408501602087016108e8565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169190910160400192915050565b6000845161097b8184602089016108e8565b80830190507f2e0000000000000000000000000000000000000000000000000000000000000080825285516109b7816001850160208a016108e8565b600192019182015283516109d28160028401602088016108e8565b0160020195945050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203610a3f57610a3f6109df565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082610a8457610a84610a46565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082821015610aca57610aca6109df565b500390565b600082610ade57610ade610a46565b500690565b60008219821115610af657610af66109df565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(DeployerWhitelistStorageLayoutJSON), DeployerWhitelistStorageLayout); err != nil {
		panic(err)
	}

	layouts["DeployerWhitelist"] = DeployerWhitelistStorageLayout
	deployedBytecodes["DeployerWhitelist"] = DeployerWhitelistDeployedBin
}
