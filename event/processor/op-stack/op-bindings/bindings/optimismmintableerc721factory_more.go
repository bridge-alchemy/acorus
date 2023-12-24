// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"encoding/json"
	"github.com/cornerstone-labs/acorus/event/processor/op-stack/op-bindings/solc"
)

const OptimismMintableERC721FactoryStorageLayoutJSON = "{\"storage\":[{\"astId\":1000,\"contract\":\"contracts/universal/OptimismMintableERC721Factory.sol:OptimismMintableERC721Factory\",\"label\":\"isOptimismMintableERC721\",\"offset\":0,\"slot\":\"0\",\"type\":\"t_mapping(t_address,t_bool)\"}],\"types\":{\"t_address\":{\"encoding\":\"inplace\",\"label\":\"address\",\"numberOfBytes\":\"20\"},\"t_bool\":{\"encoding\":\"inplace\",\"label\":\"bool\",\"numberOfBytes\":\"1\"},\"t_mapping(t_address,t_bool)\":{\"encoding\":\"mapping\",\"label\":\"mapping(address =\u003e bool)\",\"numberOfBytes\":\"32\",\"key\":\"t_address\",\"value\":\"t_bool\"}}}"

var OptimismMintableERC721FactoryStorageLayout = new(solc.StorageLayout)

var OptimismMintableERC721FactoryDeployedBin = "0x60806040523480156200001157600080fd5b50600436106200006f5760003560e01c80637d1d0c5b11620000565780637d1d0c5b14620000cd578063d97df6521462000104578063ee9a31a2146200014157600080fd5b806354fd4d5014620000745780635572acae1462000096575b600080fd5b6200007e62000169565b6040516200008d9190620005dc565b60405180910390f35b620000bc620000a736600462000622565b60006020819052908152604090205460ff1681565b60405190151581526020016200008d565b620000f57f000000000000000000000000000000000000000000000000000000000000000081565b6040519081526020016200008d565b6200011b6200011536600462000722565b62000214565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016200008d565b6200011b7f000000000000000000000000000000000000000000000000000000000000000081565b6060620001967f0000000000000000000000000000000000000000000000000000000000000000620003fa565b620001c17f0000000000000000000000000000000000000000000000000000000000000000620003fa565b620001ec7f0000000000000000000000000000000000000000000000000000000000000000620003fa565b60405160200162000200939291906200079f565b604051602081830303815290604052905090565b600073ffffffffffffffffffffffffffffffffffffffff8416620002e5576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526044602482018190527f4f7074696d69736d4d696e7461626c65455243373231466163746f72793a204c908201527f3120746f6b656e20616464726573732063616e6e6f742062652061646472657360648201527f7328302900000000000000000000000000000000000000000000000000000000608482015260a40160405180910390fd5b60007f00000000000000000000000000000000000000000000000000000000000000007f00000000000000000000000000000000000000000000000000000000000000008686866040516200033a906200054f565b6200034a9594939291906200081b565b604051809103906000f08015801562000367573d6000803e3d6000fd5b5073ffffffffffffffffffffffffffffffffffffffff8181166000818152602081815260409182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0016600117905590513381529394509188169290917fe72783bb8e0ca31286b85278da59684dd814df9762a52f0837f89edd1483b299910160405180910390a3949350505050565b6060816000036200043e57505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b81156200046e57806200045581620008ab565b9150620004669050600a8362000915565b915062000442565b60008167ffffffffffffffff8111156200048c576200048c62000640565b6040519080825280601f01601f191660200182016040528015620004b7576020820181803683370190505b5090505b84156200054757620004cf6001836200092c565b9150620004de600a8662000946565b620004eb9060306200095d565b60f81b81838151811062000503576200050362000978565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506200053f600a8662000915565b9450620004bb565b949350505050565b6132d780620009a883390190565b60005b838110156200057a57818101518382015260200162000560565b838111156200058a576000848401525b50505050565b60008151808452620005aa8160208601602086016200055d565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b602081526000620005f1602083018462000590565b9392505050565b803573ffffffffffffffffffffffffffffffffffffffff811681146200061d57600080fd5b919050565b6000602082840312156200063557600080fd5b620005f182620005f8565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b600082601f8301126200068157600080fd5b813567ffffffffffffffff808211156200069f576200069f62000640565b604051601f83017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f01168101908282118183101715620006e857620006e862000640565b816040528381528660208588010111156200070257600080fd5b836020870160208301376000602085830101528094505050505092915050565b6000806000606084860312156200073857600080fd5b6200074384620005f8565b9250602084013567ffffffffffffffff808211156200076157600080fd5b6200076f878388016200066f565b935060408601359150808211156200078657600080fd5b5062000795868287016200066f565b9150509250925092565b60008451620007b38184602089016200055d565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551620007f1816001850160208a016200055d565b600192019182015283516200080e8160028401602088016200055d565b0160020195945050505050565b600073ffffffffffffffffffffffffffffffffffffffff808816835286602084015280861660408401525060a060608301526200085c60a083018562000590565b828103608084015262000870818562000590565b98975050505050505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff8203620008df57620008df6200087c565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b600082620009275762000927620008e6565b500490565b6000828210156200094157620009416200087c565b500390565b600082620009585762000958620008e6565b500690565b600082198211156200097357620009736200087c565b500190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fdfe6101406040523480156200001257600080fd5b50604051620032d7380380620032d7833981016040819052620000359162000640565b600180600084848262000049838262000769565b50600162000058828262000769565b50505060809290925260a05260c0526001600160a01b038516620000e95760405162461bcd60e51b815260206004820152603360248201527f4f7074696d69736d4d696e7461626c654552433732313a20627269646765206360448201527f616e6e6f7420626520616464726573732830290000000000000000000000000060648201526084015b60405180910390fd5b83600003620001615760405162461bcd60e51b815260206004820152603660248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465206360448201527f6861696e2069642063616e6e6f74206265207a65726f000000000000000000006064820152608401620000e0565b6001600160a01b038316620001df5760405162461bcd60e51b815260206004820152603960248201527f4f7074696d69736d4d696e7461626c654552433732313a2072656d6f7465207460448201527f6f6b656e2063616e6e6f742062652061646472657373283029000000000000006064820152608401620000e0565b60e08490526001600160a01b03838116610100819052908616610120526200021590601462000269602090811b62000f5c17901c565b6200022b856200042960201b6200119f1760201c565b6040516020016200023e92919062000835565b604051602081830303815290604052600a90816200025d919062000769565b505050505050620009a6565b606060006200027a836002620008bf565b62000287906002620008e1565b6001600160401b03811115620002a157620002a162000566565b6040519080825280601f01601f191660200182016040528015620002cc576020820181803683370190505b509050600360fc1b81600081518110620002ea57620002ea620008fc565b60200101906001600160f81b031916908160001a905350600f60fb1b816001815181106200031c576200031c620008fc565b60200101906001600160f81b031916908160001a905350600062000342846002620008bf565b6200034f906001620008e1565b90505b6001811115620003d1576f181899199a1a9b1b9c1cb0b131b232b360811b85600f1660108110620003875762000387620008fc565b1a60f81b828281518110620003a057620003a0620008fc565b60200101906001600160f81b031916908160001a90535060049490941c93620003c98162000912565b905062000352565b508315620004225760405162461bcd60e51b815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401620000e0565b9392505050565b606081600003620004515750506040805180820190915260018152600360fc1b602082015290565b8160005b811562000481578062000468816200092c565b9150620004799050600a836200095e565b915062000455565b6000816001600160401b038111156200049e576200049e62000566565b6040519080825280601f01601f191660200182016040528015620004c9576020820181803683370190505b5090505b84156200054157620004e160018362000975565b9150620004f0600a866200098f565b620004fd906030620008e1565b60f81b818381518110620005155762000515620008fc565b60200101906001600160f81b031916908160001a90535062000539600a866200095e565b9450620004cd565b949350505050565b80516001600160a01b03811681146200056157600080fd5b919050565b634e487b7160e01b600052604160045260246000fd5b60005b83811015620005995781810151838201526020016200057f565b83811115620005a9576000848401525b50505050565b600082601f830112620005c157600080fd5b81516001600160401b0380821115620005de57620005de62000566565b604051601f8301601f19908116603f0116810190828211818310171562000609576200060962000566565b816040528381528660208588010111156200062357600080fd5b620006368460208301602089016200057c565b9695505050505050565b600080600080600060a086880312156200065957600080fd5b620006648662000549565b9450602086015193506200067b6040870162000549565b60608701519093506001600160401b03808211156200069957600080fd5b620006a789838a01620005af565b93506080880151915080821115620006be57600080fd5b50620006cd88828901620005af565b9150509295509295909350565b600181811c90821680620006ef57607f821691505b6020821081036200071057634e487b7160e01b600052602260045260246000fd5b50919050565b601f8211156200076457600081815260208120601f850160051c810160208610156200073f5750805b601f850160051c820191505b8181101562000760578281556001016200074b565b5050505b505050565b81516001600160401b0381111562000785576200078562000566565b6200079d81620007968454620006da565b8462000716565b602080601f831160018114620007d55760008415620007bc5750858301515b600019600386901b1c1916600185901b17855562000760565b600085815260208120601f198616915b828110156200080657888601518255948401946001909101908401620007e5565b5085821015620008255787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b6832ba3432b932bab69d60b91b8152600083516200085b8160098501602088016200057c565b600160fe1b60099184019182015283516200087e81600a8401602088016200057c565b712f746f6b656e5552493f75696e743235363d60701b600a9290910191820152601c01949350505050565b634e487b7160e01b600052601160045260246000fd5b6000816000190483118215151615620008dc57620008dc620008a9565b500290565b60008219821115620008f757620008f7620008a9565b500190565b634e487b7160e01b600052603260045260246000fd5b600081620009245762000924620008a9565b506000190190565b600060018201620009415762000941620008a9565b5060010190565b634e487b7160e01b600052601260045260246000fd5b60008262000970576200097062000948565b500490565b6000828210156200098a576200098a620008a9565b500390565b600082620009a157620009a162000948565b500690565b60805160a05160c05160e05161010051610120516128be62000a19600039600081816103ae0152818161044601528181610b900152610cb20152600081816101e001526103880152600081816102f501526103d4015260006109bf015260006109960152600061096d01526128be6000f3fe608060405234801561001057600080fd5b50600436106101ae5760003560e01c80637d1d0c5b116100ee578063c87b56dd11610097578063e78cea9211610071578063e78cea92146103ac578063e9518196146103d2578063e985e9c5146103f8578063ee9a31a21461044157600080fd5b8063c87b56dd1461036b578063d547cfb71461037e578063d6c0b2c41461038657600080fd5b8063a1448194116100c8578063a144819414610332578063a22cb46514610345578063b88d4fde1461035857600080fd5b80637d1d0c5b146102f057806395d89b41146103175780639dc29fac1461031f57600080fd5b806323b872dd1161015b5780634f6ccce7116101355780634f6ccce7146102af57806354fd4d50146102c25780636352211e146102ca57806370a08231146102dd57600080fd5b806323b872dd146102765780632f745c591461028957806342842e0e1461029c57600080fd5b8063081812fc1161018c578063081812fc1461023c578063095ea7b31461024f57806318160ddd1461026457600080fd5b806301ffc9a7146101b3578063033964be146101db57806306fdde0314610227575b600080fd5b6101c66101c1366004612295565b610468565b60405190151581526020015b60405180910390f35b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60405173ffffffffffffffffffffffffffffffffffffffff90911681526020016101d2565b61022f6104c6565b6040516101d29190612328565b61020261024a36600461233b565b610558565b61026261025d36600461237d565b61058c565b005b6008545b6040519081526020016101d2565b6102626102843660046123a7565b61071d565b61026861029736600461237d565b6107be565b6102626102aa3660046123a7565b61088d565b6102686102bd36600461233b565b6108a8565b61022f610966565b6102026102d836600461233b565b610a09565b6102686102eb3660046123e3565b610a9b565b6102687f000000000000000000000000000000000000000000000000000000000000000081565b61022f610b69565b61026261032d36600461237d565b610b78565b61026261034036600461237d565b610c9a565b6102626103533660046123fe565b610db1565b610262610366366004612469565b610dc0565b61022f61037936600461233b565b610e68565b61022f610ece565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f0000000000000000000000000000000000000000000000000000000000000000610202565b7f0000000000000000000000000000000000000000000000000000000000000000610268565b6101c6610406366004612563565b73ffffffffffffffffffffffffffffffffffffffff918216600090815260056020908152604080832093909416825291909152205460ff1690565b6102027f000000000000000000000000000000000000000000000000000000000000000081565b60007f74259ebf000000000000000000000000000000000000000000000000000000007fffffffff0000000000000000000000000000000000000000000000000000000083168114806104bf57506104bf836112dc565b9392505050565b6060600080546104d590612596565b80601f016020809104026020016040519081016040528092919081815260200182805461050190612596565b801561054e5780601f106105235761010080835404028352916020019161054e565b820191906000526020600020905b81548152906001019060200180831161053157829003601f168201915b5050505050905090565b600061056382611332565b5060009081526004602052604090205473ffffffffffffffffffffffffffffffffffffffff1690565b600061059782610a09565b90508073ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603610659576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602160248201527f4552433732313a20617070726f76616c20746f2063757272656e74206f776e6560448201527f720000000000000000000000000000000000000000000000000000000000000060648201526084015b60405180910390fd5b3373ffffffffffffffffffffffffffffffffffffffff8216148061068257506106828133610406565b61070e576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603e60248201527f4552433732313a20617070726f76652063616c6c6572206973206e6f7420746f60448201527f6b656e206f776e6572206e6f7220617070726f76656420666f7220616c6c00006064820152608401610650565b61071883836113c0565b505050565b6107273382611460565b6107b3576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610650565b61071883838361151f565b60006107c983610a9b565b8210610857576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602b60248201527f455243373231456e756d657261626c653a206f776e657220696e646578206f7560448201527f74206f6620626f756e64730000000000000000000000000000000000000000006064820152608401610650565b5073ffffffffffffffffffffffffffffffffffffffff919091166000908152600660209081526040808320938352929052205490565b61071883838360405180602001604052806000815250610dc0565b60006108b360085490565b8210610941576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602c60248201527f455243373231456e756d657261626c653a20676c6f62616c20696e646578206f60448201527f7574206f6620626f756e647300000000000000000000000000000000000000006064820152608401610650565b60088281548110610954576109546125e9565b90600052602060002001549050919050565b60606109917f000000000000000000000000000000000000000000000000000000000000000061119f565b6109ba7f000000000000000000000000000000000000000000000000000000000000000061119f565b6109e37f000000000000000000000000000000000000000000000000000000000000000061119f565b6040516020016109f593929190612618565b604051602081830303815290604052905090565b60008181526002602052604081205473ffffffffffffffffffffffffffffffffffffffff1680610a95576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610650565b92915050565b600073ffffffffffffffffffffffffffffffffffffffff8216610b40576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602960248201527f4552433732313a2061646472657373207a65726f206973206e6f74206120766160448201527f6c6964206f776e657200000000000000000000000000000000000000000000006064820152608401610650565b5073ffffffffffffffffffffffffffffffffffffffff1660009081526003602052604090205490565b6060600180546104d590612596565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610c3d576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e0000000000006064820152608401610650565b610c4681611791565b8173ffffffffffffffffffffffffffffffffffffffff167fcc16f5dbb4873280815c1ee09dbd06736cffcc184412cf7a71a0fdb75d397ca582604051610c8e91815260200190565b60405180910390a25050565b3373ffffffffffffffffffffffffffffffffffffffff7f00000000000000000000000000000000000000000000000000000000000000001614610d5f576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603a60248201527f4f7074696d69736d4d696e7461626c654552433732313a206f6e6c792062726960448201527f6467652063616e2063616c6c20746869732066756e6374696f6e0000000000006064820152608401610650565b610d69828261186a565b8173ffffffffffffffffffffffffffffffffffffffff167f0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d412139688582604051610c8e91815260200190565b610dbc338383611884565b5050565b610dca3383611460565b610e56576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602e60248201527f4552433732313a2063616c6c6572206973206e6f7420746f6b656e206f776e6560448201527f72206e6f7220617070726f7665640000000000000000000000000000000000006064820152608401610650565b610e62848484846119b1565b50505050565b6060610e7382611332565b6000610e7d611a54565b90506000815111610e9d57604051806020016040528060008152506104bf565b80610ea78461119f565b604051602001610eb892919061268e565b6040516020818303038152906040529392505050565b600a8054610edb90612596565b80601f0160208091040260200160405190810160405280929190818152602001828054610f0790612596565b8015610f545780601f10610f2957610100808354040283529160200191610f54565b820191906000526020600020905b815481529060010190602001808311610f3757829003601f168201915b505050505081565b60606000610f6b8360026126ec565b610f76906002612729565b67ffffffffffffffff811115610f8e57610f8e61243a565b6040519080825280601f01601f191660200182016040528015610fb8576020820181803683370190505b5090507f300000000000000000000000000000000000000000000000000000000000000081600081518110610fef57610fef6125e9565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053507f780000000000000000000000000000000000000000000000000000000000000081600181518110611052576110526125e9565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a905350600061108e8460026126ec565b611099906001612729565b90505b6001811115611136577f303132333435363738396162636465660000000000000000000000000000000085600f16601081106110da576110da6125e9565b1a60f81b8282815181106110f0576110f06125e9565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a90535060049490941c9361112f81612741565b905061109c565b5083156104bf576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f537472696e67733a20686578206c656e67746820696e73756666696369656e746044820152606401610650565b6060816000036111e257505060408051808201909152600181527f3000000000000000000000000000000000000000000000000000000000000000602082015290565b8160005b811561120c57806111f681612776565b91506112059050600a836127dd565b91506111e6565b60008167ffffffffffffffff8111156112275761122761243a565b6040519080825280601f01601f191660200182016040528015611251576020820181803683370190505b5090505b84156112d4576112666001836127f1565b9150611273600a86612808565b61127e906030612729565b60f81b818381518110611293576112936125e9565b60200101907effffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff1916908160001a9053506112cd600a866127dd565b9450611255565b949350505050565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f780e9d63000000000000000000000000000000000000000000000000000000001480610a955750610a9582611a63565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff166113bd576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601860248201527f4552433732313a20696e76616c696420746f6b656e20494400000000000000006044820152606401610650565b50565b600081815260046020526040902080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff8416908117909155819061141a82610a09565b73ffffffffffffffffffffffffffffffffffffffff167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b60008061146c83610a09565b90508073ffffffffffffffffffffffffffffffffffffffff168473ffffffffffffffffffffffffffffffffffffffff1614806114da575073ffffffffffffffffffffffffffffffffffffffff80821660009081526005602090815260408083209388168352929052205460ff165b806112d457508373ffffffffffffffffffffffffffffffffffffffff1661150084610558565b73ffffffffffffffffffffffffffffffffffffffff1614949350505050565b8273ffffffffffffffffffffffffffffffffffffffff1661153f82610a09565b73ffffffffffffffffffffffffffffffffffffffff16146115e2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152602560248201527f4552433732313a207472616e736665722066726f6d20696e636f72726563742060448201527f6f776e65720000000000000000000000000000000000000000000000000000006064820152608401610650565b73ffffffffffffffffffffffffffffffffffffffff8216611684576040517f08c379a0000000000000000000000000000000000000000000000000000000008152602060048201526024808201527f4552433732313a207472616e7366657220746f20746865207a65726f2061646460448201527f72657373000000000000000000000000000000000000000000000000000000006064820152608401610650565b61168f838383611b46565b61169a6000826113c0565b73ffffffffffffffffffffffffffffffffffffffff831660009081526003602052604081208054600192906116d09084906127f1565b909155505073ffffffffffffffffffffffffffffffffffffffff8216600090815260036020526040812080546001929061170b908490612729565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff86811691821790925591518493918716917fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef91a4505050565b600061179c82610a09565b90506117aa81600084611b46565b6117b56000836113c0565b73ffffffffffffffffffffffffffffffffffffffff811660009081526003602052604081208054600192906117eb9084906127f1565b909155505060008281526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001690555183919073ffffffffffffffffffffffffffffffffffffffff8416907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908390a45050565b610dbc828260405180602001604052806000815250611c4c565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1603611919576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601960248201527f4552433732313a20617070726f766520746f2063616c6c6572000000000000006044820152606401610650565b73ffffffffffffffffffffffffffffffffffffffff83811660008181526005602090815260408083209487168084529482529182902080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff001686151590811790915591519182527f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a3505050565b6119bc84848461151f565b6119c884848484611cef565b610e62576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610650565b6060600a80546104d590612596565b60007fffffffff0000000000000000000000000000000000000000000000000000000082167f80ac58cd000000000000000000000000000000000000000000000000000000001480611af657507fffffffff0000000000000000000000000000000000000000000000000000000082167f5b5e139f00000000000000000000000000000000000000000000000000000000145b80610a9557507f01ffc9a7000000000000000000000000000000000000000000000000000000007fffffffff00000000000000000000000000000000000000000000000000000000831614610a95565b73ffffffffffffffffffffffffffffffffffffffff8316611bae57611ba981600880546000838152600960205260408120829055600182018355919091527ff3f7a9fe364faab93b216da50a3214154f22a0a2b415b23a84c8169e8b636ee30155565b611beb565b8173ffffffffffffffffffffffffffffffffffffffff168373ffffffffffffffffffffffffffffffffffffffff1614611beb57611beb8382611ee2565b73ffffffffffffffffffffffffffffffffffffffff8216611c0f5761071881611f99565b8273ffffffffffffffffffffffffffffffffffffffff168273ffffffffffffffffffffffffffffffffffffffff1614610718576107188282612048565b611c568383612099565b611c636000848484611cef565b610718576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610650565b600073ffffffffffffffffffffffffffffffffffffffff84163b15611ed7576040517f150b7a0200000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff85169063150b7a0290611d6690339089908890889060040161281c565b6020604051808303816000875af1925050508015611dbf575060408051601f3d9081017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0168201909252611dbc91810190612865565b60015b611e8c573d808015611ded576040519150601f19603f3d011682016040523d82523d6000602084013e611df2565b606091505b508051600003611e84576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152603260248201527f4552433732313a207472616e7366657220746f206e6f6e20455243373231526560448201527f63656976657220696d706c656d656e74657200000000000000000000000000006064820152608401610650565b805181602001fd5b7fffffffff00000000000000000000000000000000000000000000000000000000167f150b7a02000000000000000000000000000000000000000000000000000000001490506112d4565b506001949350505050565b60006001611eef84610a9b565b611ef991906127f1565b600083815260076020526040902054909150808214611f595773ffffffffffffffffffffffffffffffffffffffff841660009081526006602090815260408083208584528252808320548484528184208190558352600790915290208190555b50600091825260076020908152604080842084905573ffffffffffffffffffffffffffffffffffffffff9094168352600681528383209183525290812055565b600854600090611fab906001906127f1565b60008381526009602052604081205460088054939450909284908110611fd357611fd36125e9565b906000526020600020015490508060088381548110611ff457611ff46125e9565b600091825260208083209091019290925582815260099091526040808220849055858252812055600880548061202c5761202c612882565b6001900381819060005260206000200160009055905550505050565b600061205383610a9b565b73ffffffffffffffffffffffffffffffffffffffff9093166000908152600660209081526040808320868452825280832085905593825260079052919091209190915550565b73ffffffffffffffffffffffffffffffffffffffff8216612116576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820181905260248201527f4552433732313a206d696e7420746f20746865207a65726f20616464726573736044820152606401610650565b60008181526002602052604090205473ffffffffffffffffffffffffffffffffffffffff16156121a2576040517f08c379a000000000000000000000000000000000000000000000000000000000815260206004820152601c60248201527f4552433732313a20746f6b656e20616c7265616479206d696e746564000000006044820152606401610650565b6121ae60008383611b46565b73ffffffffffffffffffffffffffffffffffffffff821660009081526003602052604081208054600192906121e4908490612729565b909155505060008181526002602052604080822080547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff861690811790915590518392907fddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef908290a45050565b7fffffffff00000000000000000000000000000000000000000000000000000000811681146113bd57600080fd5b6000602082840312156122a757600080fd5b81356104bf81612267565b60005b838110156122cd5781810151838201526020016122b5565b83811115610e625750506000910152565b600081518084526122f68160208601602086016122b2565b601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169290920160200192915050565b6020815260006104bf60208301846122de565b60006020828403121561234d57600080fd5b5035919050565b803573ffffffffffffffffffffffffffffffffffffffff8116811461237857600080fd5b919050565b6000806040838503121561239057600080fd5b61239983612354565b946020939093013593505050565b6000806000606084860312156123bc57600080fd5b6123c584612354565b92506123d360208501612354565b9150604084013590509250925092565b6000602082840312156123f557600080fd5b6104bf82612354565b6000806040838503121561241157600080fd5b61241a83612354565b91506020830135801515811461242f57600080fd5b809150509250929050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000806000806080858703121561247f57600080fd5b61248885612354565b935061249660208601612354565b925060408501359150606085013567ffffffffffffffff808211156124ba57600080fd5b818701915087601f8301126124ce57600080fd5b8135818111156124e0576124e061243a565b604051601f82017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0908116603f011681019083821181831017156125265761252661243a565b816040528281528a602084870101111561253f57600080fd5b82602086016020830137600060208483010152809550505050505092959194509250565b6000806040838503121561257657600080fd5b61257f83612354565b915061258d60208401612354565b90509250929050565b600181811c908216806125aa57607f821691505b6020821081036125e3577f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b50919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603260045260246000fd5b6000845161262a8184602089016122b2565b80830190507f2e000000000000000000000000000000000000000000000000000000000000008082528551612666816001850160208a016122b2565b600192019182015283516126818160028401602088016122b2565b0160020195945050505050565b600083516126a08184602088016122b2565b8351908301906126b48183602088016122b2565b01949350505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b6000817fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0483118215151615612724576127246126bd565b500290565b6000821982111561273c5761273c6126bd565b500190565b600081612750576127506126bd565b507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff0190565b60007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff82036127a7576127a76126bd565b5060010190565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601260045260246000fd5b6000826127ec576127ec6127ae565b500490565b600082821015612803576128036126bd565b500390565b600082612817576128176127ae565b500690565b600073ffffffffffffffffffffffffffffffffffffffff80871683528086166020840152508360408301526080606083015261285b60808301846122de565b9695505050505050565b60006020828403121561287757600080fd5b81516104bf81612267565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052603160045260246000fdfea164736f6c634300080f000aa164736f6c634300080f000a"

func init() {
	if err := json.Unmarshal([]byte(OptimismMintableERC721FactoryStorageLayoutJSON), OptimismMintableERC721FactoryStorageLayout); err != nil {
		panic(err)
	}

	layouts["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryStorageLayout
	deployedBytecodes["OptimismMintableERC721Factory"] = OptimismMintableERC721FactoryDeployedBin
}
