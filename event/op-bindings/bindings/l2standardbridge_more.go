// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"
	"github.com/cornerstone-labs/acorus/event/op-bindings/solc"
)

const L2StandardBridgeStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_0_0_20\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_address\"},{\"astId\":1001,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_1_0_20\",\"offset\":0,\"slot\":\"1\",\"type\":\"t_address\"},{\"astId\":1002,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"spacer_2_0_20\",\"offset\":0,\"slot\":\"2\",\"type\":\"t_address\"},{\"astId\":1003,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"deposits\",\"offset\":0,\"slot\":\"3\",\"type\":\"t_mapping(t_address,t_mapping(t_address,t_uint256))\"},{\"astId\":1004,\"contract\":\"contracts/L2/L2StandardBridge.sol:L2StandardBridge\",\"label\":\"__gap\",\"offset\":0,\"slot\":\"4\",\"type\":\"t_array(t_uint256)1005_storage\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_array(t_uint256)1005_storage\":{\"encoding\":\"inplace\",\"label\":\"uint256[47]\",\"numberOfBytes\":\"1504\"},\"t_mapping(t_address,t_mapping(t_address,t_uint256))\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e mapping(address =\u003e uint256))\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_mapping(t_address,t_uint256)\"},\"t_mapping(t_address,t_uint256)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e uint256)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_uint256\"},\"t_uint256\":{\"encoding\":\"inplace\",\"label\":\"uint256\",\"numberOfBytes\":\"32\"}}}"

var L2StandardBridgeStorageLayout = new(solc.StorageLayout)

var L2StandardBridgeDeployedBin = "0x6080604052600436106101485760003560e01c806354fd4d50116100c05780638f601f6611610074578063a3a7954811610059578063a3a7954814610448578063ac6986c51461045b578063f407a99e1461048f57600080fd5b80638f601f66146103ce578063927ede2d1461041457600080fd5b8063662a633a116100a5578063662a633a146103675780637f46ddb21461037a57806387087623146103ae57600080fd5b806354fd4d501461033257806357eccc341461035457600080fd5b806332b7006d116101175780633cb747bf116100fc5780633cb747bf146102cc578063540abf73146102ff578063548e0a5c1461031f57600080fd5b806332b7006d1461026057806336c717c11461027357600080fd5b80630166a07a14610207578063162f1686146102275780631635f5fd1461023a57806324ca60181461024d57600080fd5b3661020257333b156101e1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084015b60405180910390fd5b61020033333462030d40604051806020016040528060008152506104a2565b005b600080fd5b34801561021357600080fd5b50610200610222366004612fe8565b61068b565b610200610235366004613099565b610a81565b6102006102483660046130f3565b610b59565b61020061025b366004613166565b610f26565b61020061026e3660046131ba565b610f69565b34801561027f57600080fd5b507f00000000000000000000000000000000000000000000000000000000000000005b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020015b60405180910390f35b3480156102d857600080fd5b507f00000000000000000000000000000000000000000000000000000000000000006102a2565b34801561030b57600080fd5b5061020061031a3660046131f2565b61103c565b61020061032d366004613269565b611081565b34801561033e57600080fd5b506103476110c4565b6040516102c3919061330e565b610200610362366004613321565b611167565b610200610375366004612fe8565b6111af565b34801561038657600080fd5b506102a27f000000000000000000000000000000000000000000000000000000000000000081565b3480156103ba57600080fd5b506102006103c9366004613374565b6112b8565b3480156103da57600080fd5b506104066103e93660046133f7565b600360209081526000928352604080842090915290825290205481565b6040519081526020016102c3565b34801561042057600080fd5b506102a27f000000000000000000000000000000000000000000000000000000000000000081565b610200610456366004613374565b611394565b34801561046757600080fd5b506102a27f000000000000000000000000000000000000000000000000000000000000000081565b61020061049d3660046130f3565b6113d8565b823414610531576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f5374616e646172644272696467653a206272696467696e67204d4e54206d757360448201527f7420696e636c7564652073756666696369656e74204d4e542076616c7565000060648201526084016101d8565b61053d858585846118cc565b60007f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b575e30034837f000000000000000000000000000000000000000000000000000000000000000063f407a99e60e01b8b8b8b8a6040516024016105bd9493929190613430565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e087901b90921682526106519392918a90600401613479565b6000604051808303818588803b15801561066a57600080fd5b505af115801561067e573d6000803e3d6000fd5b5050505050505050505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156107a957507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa15801561076d573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061079191906134c5565b73ffffffffffffffffffffffffffffffffffffffff16145b61085b576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101d8565b61086487611979565b156109b25761087387876119db565b610925576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101d8565b6040517f40c10f1900000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8581166004830152602482018590528816906340c10f1990604401600060405180830381600087803b15801561099557600080fd5b505af11580156109a9573d6000803e3d6000fd5b50505050610a34565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600360209081526040808320938a16835292905220546109f0908490613511565b73ffffffffffffffffffffffffffffffffffffffff8089166000818152600360209081526040808320948c1683529390529190912091909155610a34908585611afb565b610a78878787878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611bcf92505050565b50505050505050565b333b15610b10576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101d8565b610b533333868686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c5d92505050565b50505050565b3373ffffffffffffffffffffffffffffffffffffffff7f000000000000000000000000000000000000000000000000000000000000000016148015610c7757507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa158015610c3b573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c5f91906134c5565b73ffffffffffffffffffffffffffffffffffffffff16145b610d29576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101d8565b3073ffffffffffffffffffffffffffffffffffffffff851603610dce576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101d8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1603610ea9576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101d8565b610edd73deaddeaddeaddeaddeaddeaddeaddeaddead11117342000000000000000000000000000000000000078686611e7e565b610f1f85858585858080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611edc92505050565b5050505050565b610f1f3385878686868080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611c5d92505050565b333b15610ff8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101d8565b610f1f853333878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611f7d92505050565b610a7887873388888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061207292505050565b610b533385348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104a292505050565b60606110ef7f0000000000000000000000000000000000000000000000000000000000000000612573565b6111187f0000000000000000000000000000000000000000000000000000000000000000612573565b6111417f0000000000000000000000000000000000000000000000000000000000000000612573565b60405160200161115393929190613528565b604051602081830303815290604052905090565b6111aa3333348686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506104a292505050565b505050565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff16148015611233575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead0000145b1561124a5761124585858585856113d8565b610a78565b73ffffffffffffffffffffffffffffffffffffffff8716158015611297575073ffffffffffffffffffffffffffffffffffffffff861673deaddeaddeaddeaddeaddeaddeaddeaddead1111145b156112a9576112458585858585610b59565b610a788688878787878761068b565b333b15611347576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603760248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20616e20454f4100000000000000000060648201526084016101d8565b61138c86863333888888888080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525061207292505050565b505050505050565b61138c863387878787878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250611f7d92505050565b3373ffffffffffffffffffffffffffffffffffffffff7f0000000000000000000000000000000000000000000000000000000000000000161480156114f657507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff16636e296e456040518163ffffffff1660e01b8152600401602060405180830381865afa1580156114ba573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906114de91906134c5565b73ffffffffffffffffffffffffffffffffffffffff16145b6115a8576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604160248201527f5374616e646172644272696467653a2066756e6374696f6e2063616e206f6e6c60448201527f792062652063616c6c65642066726f6d20746865206f7468657220627269646760648201527f6500000000000000000000000000000000000000000000000000000000000000608482015260a4016101d8565b823414611637576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f5374616e646172644272696467653a20616d6f756e742073656e7420646f657360448201527f206e6f74206d6174636820616d6f756e7420726571756972656400000000000060648201526084016101d8565b3073ffffffffffffffffffffffffffffffffffffffff8516036116dc576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f207360448201527f656c66000000000000000000000000000000000000000000000000000000000060648201526084016101d8565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff16036117b7576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602860248201527f5374616e646172644272696467653a2063616e6e6f742073656e6420746f206d60448201527f657373656e67657200000000000000000000000000000000000000000000000060648201526084016101d8565b60006117fb855a8686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506126b092505050565b90508061188a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602360248201527f5374616e646172644272696467653a204d4e54207472616e736665722066616960448201527f6c6564000000000000000000000000000000000000000000000000000000000060648201526084016101d8565b61138c86868686868080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920191909152506126ca92505050565b8373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e8686866040516119659392919061359e565b60405180910390a4610b5384848484612777565b60006119a5827f1d1d8b63000000000000000000000000000000000000000000000000000000006127e4565b806119d557506119d5827fec4fc8e3000000000000000000000000000000000000000000000000000000006127e4565b92915050565b6000611a07837f1d1d8b63000000000000000000000000000000000000000000000000000000006127e4565b15611ab0578273ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a57573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611a7b91906134c5565b73ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff161490506119d5565b8273ffffffffffffffffffffffffffffffffffffffff1663d6c0b2c46040518163ffffffff1660e01b8152600401602060405180830381865afa158015611a57573d6000803e3d6000fd5b60405173ffffffffffffffffffffffffffffffffffffffff83166024820152604481018290526111aa9084907fa9059cbb00000000000000000000000000000000000000000000000000000000906064015b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529190526020810180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff0000000000000000000000000000000000000000000000000000000090931692909217909152612807565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611c479392919061359e565b60405180910390a461138c868686868686612913565b611c7d73deaddeaddeaddeaddeaddeaddeaddeaddead1111333086611e7e565b6040517f095ea7b300000000000000000000000000000000000000000000000000000000815273420000000000000000000000000000000000000760048201526024810184905273deaddeaddeaddeaddeaddeaddeaddeaddead11119063095ea7b3906044016020604051808303816000875af1158015611d02573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190611d2691906135dc565b50611d338585858461299b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b575e30034857f0000000000000000000000000000000000000000000000000000000000000000631635f5fd60e01b8a8a8a89604051602401611db19493929190613430565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e087901b9092168252611e459392918990600401613479565b6000604051808303818588803b158015611e5e57600080fd5b505af1158015611e72573d6000803e3d6000fd5b50505050505050505050565b60405173ffffffffffffffffffffffffffffffffffffffff80851660248301528316604482015260648101829052610b539085907f23b872dd0000000000000000000000000000000000000000000000000000000090608401611b4d565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead111173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd89868686604051611f699392919061359e565b60405180910390a4610b5384848484612a3c565b7fffffffffffffffffffffffff215221522152215221522152215221522152eeef73ffffffffffffffffffffffffffffffffffffffff871601611fcc57611fc78585858585611c5d565b61138c565b73ffffffffffffffffffffffffffffffffffffffff8616611ff457611fc785858585856104a2565b60008673ffffffffffffffffffffffffffffffffffffffff1663c01e1bd66040518163ffffffff1660e01b8152600401602060405180830381865afa158015612041573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061206591906134c5565b9050610a78878288888888885b73ffffffffffffffffffffffffffffffffffffffff871673deaddeaddeaddeaddeaddeaddeaddeaddead1111148015906120c1575073ffffffffffffffffffffffffffffffffffffffff861615155b61214d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c325374616e646172644272696467653a20427269646765455243323020646f60448201527f206e6f7420737570706f727420455448206272696467696e672e00000000000060648201526084016101d8565b73ffffffffffffffffffffffffffffffffffffffff8716158015906121be57507f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff1614155b61224a576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4c325374616e646172644272696467653a20427269646765455243323020646f60448201527f206e6f7420737570706f7274204d4e54206272696467696e672e00000000000060648201526084016101d8565b61225387611979565b156123a15761226287876119db565b612314576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152604a60248201527f5374616e646172644272696467653a2077726f6e672072656d6f746520746f6b60448201527f656e20666f72204f7074696d69736d204d696e7461626c65204552433230206c60648201527f6f63616c20746f6b656e00000000000000000000000000000000000000000000608482015260a4016101d8565b6040517f9dc29fac00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff868116600483015260248201859052881690639dc29fac90604401600060405180830381600087803b15801561238457600080fd5b505af1158015612398573d6000803e3d6000fd5b50505050612435565b6123c373ffffffffffffffffffffffffffffffffffffffff8816863086611e7e565b73ffffffffffffffffffffffffffffffffffffffff8088166000908152600360209081526040808320938a16835292905220546124019084906135fe565b73ffffffffffffffffffffffffffffffffffffffff8089166000908152600360209081526040808320938b16835292905220555b612443878787878786612a9b565b7f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff1663b575e30060007f0000000000000000000000000000000000000000000000000000000000000000630166a07a60e01b8a8c8b8b8b8a6040516024016124c596959493929190613616565b604080517fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe08184030181529181526020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167fffffffff000000000000000000000000000000000000000000000000000000009485161790525160e086901b90921682526125599392918890600401613479565b600060405180830381600087803b15801561066a57600080fd5b6060816000036125b657505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156125e057806125ca81613671565b91506125d99050600a836136d8565b91506125ba565b60008167ffffffffffffffff8111156125fb576125fb6136ec565b6040519080825280601f01601f191660200182016040528015612625576020820181803683370190505b5090505b84156126a85761263a600183613511565b9150612647600a8661371b565b6126529060306135fe565b60f81b8183815181106126675761266761372f565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506126a1600a866136d8565b9450612629565b949350505050565b600080600080845160208601878a8af19695505050505050565b8373ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f000000000000000000000000000000000000000000000000000000000000000073ffffffffffffffffffffffffffffffffffffffff167fb0444523268717a02698be47d0803aa7468c00acbed2f8bd93a0459cde61dd898686866040516127639392919061359e565b60405180910390a4610b5384848484612b29565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f74bbfec0d26a17c2367408038090a9a4e1cd1671129dc8fdf57f146a499fe3d584846040516127d692919061375e565b60405180910390a350505050565b60006127ef83612b88565b801561280057506128008383612bec565b9392505050565b6000612869826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c65648152508573ffffffffffffffffffffffffffffffffffffffff16612cbb9092919063ffffffff16565b8051909150156111aa578080602001905181019061288791906135dc565b6111aa576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e60448201527f6f7420737563636565640000000000000000000000000000000000000000000060648201526084016101d8565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167fd59c65b35445225835c83f50b6ede06a7be047d22e357073e250d9af537518cd86868660405161298b9392919061359e565b60405180910390a4505050505050565b8373ffffffffffffffffffffffffffffffffffffffff1673deaddeaddeaddeaddeaddeaddeaddeaddead111173ffffffffffffffffffffffffffffffffffffffff16600073ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051612a289392919061359e565b60405180910390a4610b5384848484612cca565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f31b2166ff604fc5672ea5df08a78081d2bc6d746cadce880747f3643d819e83d84846040516127d692919061375e565b8373ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff168673ffffffffffffffffffffffffffffffffffffffff167f73d170910aba9e6d50b102db522b1dbcd796216f5128b445aa2135272886497e868686604051612b139392919061359e565b60405180910390a461138c868686868686612d29565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167fef2dd684d0d947aa195ea84c18e3b5c457d3462c09eb29b20aac4f7d4d4f003584846040516127d692919061375e565b6000612bb4827f01ffc9a700000000000000000000000000000000000000000000000000000000612bec565b80156119d55750612be5827fffffffff00000000000000000000000000000000000000000000000000000000612bec565b1592915050565b604080517fffffffff000000000000000000000000000000000000000000000000000000008316602480830191909152825180830390910181526044909101909152602080820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff167f01ffc9a700000000000000000000000000000000000000000000000000000000178152825160009392849283928392918391908a617530fa92503d91506000519050828015612ca4575060208210155b8015612cb05750600081115b979650505050505050565b60606126a88484600085612da1565b8273ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff167f2849b43074093a05396b6f2a937dee8565b15a48a7b3d4bffb732a5017380af584846040516127d692919061375e565b8373ffffffffffffffffffffffffffffffffffffffff168573ffffffffffffffffffffffffffffffffffffffff168773ffffffffffffffffffffffffffffffffffffffff167f7ff126db8024424bbfd9826e8ab82ff59136289ea440b04b39a0df1b03b9cabf86868660405161298b9392919061359e565b606082471015612e33576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f60448201527f722063616c6c000000000000000000000000000000000000000000000000000060648201526084016101d8565b73ffffffffffffffffffffffffffffffffffffffff85163b612eb1576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016101d8565b6000808673ffffffffffffffffffffffffffffffffffffffff168587604051612eda9190613777565b60006040518083038185875af1925050503d8060008114612f17576040519150601f19603f3d011682016040523d82523d6000602084013e612f1c565b606091505b5091509150612cb082828660608315612f36575081612800565b825115612f465782518084602001fd5b816040517f08c379a00000000000000000000000000000000000000000000000000000000081526004016101d8919061330e565b73ffffffffffffffffffffffffffffffffffffffff81168114612f9c57600080fd5b50565b60008083601f840112612fb157600080fd5b50813567ffffffffffffffff811115612fc957600080fd5b602083019150836020828501011115612fe157600080fd5b9250929050565b600080600080600080600060c0888a03121561300357600080fd5b873561300e81612f7a565b9650602088013561301e81612f7a565b9550604088013561302e81612f7a565b9450606088013561303e81612f7a565b93506080880135925060a088013567ffffffffffffffff81111561306157600080fd5b61306d8a828b01612f9f565b989b979a50959850939692959293505050565b803563ffffffff8116811461309457600080fd5b919050565b600080600080606085870312156130af57600080fd5b843593506130bf60208601613080565b9250604085013567ffffffffffffffff8111156130db57600080fd5b6130e787828801612f9f565b95989497509550505050565b60008060008060006080868803121561310b57600080fd5b853561311681612f7a565b9450602086013561312681612f7a565b935060408601359250606086013567ffffffffffffffff81111561314957600080fd5b61315588828901612f9f565b969995985093965092949392505050565b60008060008060006080868803121561317e57600080fd5b85359450602086013561319081612f7a565b935061319e60408701613080565b9250606086013567ffffffffffffffff81111561314957600080fd5b6000806000806000608086880312156131d257600080fd5b85356131dd81612f7a565b94506020860135935061319e60408701613080565b600080600080600080600060c0888a03121561320d57600080fd5b873561321881612f7a565b9650602088013561322881612f7a565b9550604088013561323881612f7a565b94506060880135935061324d60808901613080565b925060a088013567ffffffffffffffff81111561306157600080fd5b6000806000806060858703121561327f57600080fd5b843561328a81612f7a565b93506130bf60208601613080565b60005b838110156132b357818101518382015260200161329b565b83811115610b535750506000910152565b600081518084526132dc816020860160208601613298565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b60208152600061280060208301846132c4565b60008060006040848603121561333657600080fd5b61333f84613080565b9250602084013567ffffffffffffffff81111561335b57600080fd5b61336786828701612f9f565b9497909650939450505050565b60008060008060008060a0878903121561338d57600080fd5b863561339881612f7a565b955060208701356133a881612f7a565b9450604087013593506133bd60608801613080565b9250608087013567ffffffffffffffff8111156133d957600080fd5b6133e589828a01612f9f565b979a9699509497509295939492505050565b6000806040838503121561340a57600080fd5b823561341581612f7a565b9150602083013561342581612f7a565b809150509250929050565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261346f60808301846132c4565b9695505050505050565b84815273ffffffffffffffffffffffffffffffffffffffff841660208201526080604082015260006134ae60808301856132c4565b905063ffffffff8316606083015295945050505050565b6000602082840312156134d757600080fd5b815161280081612f7a565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b600082821015613523576135236134e2565b500390565b6000845161353a818460208901613298565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551613576816001850160208a01613298565b60019201918201528351613591816002840160208801613298565b0160020195945050505050565b73ffffffffffffffffffffffffffffffffffffffff841681528260208201526060604082015260006135d360608301846132c4565b95945050505050565b6000602082840312156135ee57600080fd5b8151801515811461280057600080fd5b60008219821115613611576136116134e2565b500190565b600073ffffffffffffffffffffffffffffffffffffffff80891683528088166020840152808716604084015280861660608401525083608083015260c060a083015261366560c08301846132c4565b98975050505050505050565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036136a2576136a26134e2565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826136e7576136e76136a9565b500490565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b60008261372a5761372a6136a9565b500690565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b8281526040602082015260006126a860408301846132c4565b60008251613789818460208701613298565b919091019291505056fea164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(L2StandardBridgeStorageLayoutJSON), L2StandardBridgeStorageLayout); err != nil {
		panic(err)
	}

	layouts["L2StandardBridge"] = L2StandardBridgeStorageLayout
	deployedBytecodes["L2StandardBridge"] = L2StandardBridgeDeployedBin
}
