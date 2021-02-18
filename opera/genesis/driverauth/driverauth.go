package driverauth

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

// GetContractBin is NodeDriverAuth contract genesis implementation bin code
// Has to be compiled with flag bin-runtime
func GetContractBin() []byte {
	return hexutil.MustDecode("0x608060405234801561001057600080fd5b50600436106101365760003560e01c80638da5cb5b116100b2578063c0c53b8b11610081578063e08d7e6611610066578063e08d7e66146104f5578063ebdf104c14610565578063f2fde38b146106cb57610136565b8063c0c53b8b14610475578063d6a0c7af146104ba57610136565b80638da5cb5b146103955780638f32d59b146103c6578063a4066fbe146103e2578063b9cc6b1c1461040557610136565b8063267ab446116101095780634feb92f3116100ee5780634feb92f3146102a957806366e7ea0f14610354578063715018a61461038d57610136565b8063267ab446146102595780634ddaf8f21461027657610136565b80630aeeca001461013b57806318f628d41461015a5780631e702f83146101bf578063242a6e3f146101e2575b600080fd5b6101586004803603602081101561015157600080fd5b50356106fe565b005b610158600480360361012081101561017157600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060208101359060408101359060608101359060808101359060a08101359060c08101359060e08101359061010001356107e5565b610158600480360360408110156101d557600080fd5b508035906020013561090c565b610158600480360360408110156101f857600080fd5b8135919081019060408101602082013564010000000081111561021a57600080fd5b82018360208201111561022c57600080fd5b8035906020019184600183028401116401000000008311171561024e57600080fd5b5090925090506109f8565b6101586004803603602081101561026f57600080fd5b5035610b27565b6101586004803603602081101561028c57600080fd5b503573ffffffffffffffffffffffffffffffffffffffff16610bf3565b61015860048036036101008110156102c057600080fd5b73ffffffffffffffffffffffffffffffffffffffff823516916020810135918101906060810160408201356401000000008111156102fd57600080fd5b82018360208201111561030f57600080fd5b8035906020019184600183028401116401000000008311171561033157600080fd5b919350915080359060208101359060408101359060608101359060800135610cc0565b6101586004803603604081101561036a57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff8135169060200135610e1b565b610158610f80565b61039d611048565b6040805173ffffffffffffffffffffffffffffffffffffffff9092168252519081900360200190f35b6103ce611064565b604080519115158252519081900360200190f35b610158600480360360408110156103f857600080fd5b5080359060200135611082565b6101586004803603602081101561041b57600080fd5b81019060208101813564010000000081111561043657600080fd5b82018360208201111561044857600080fd5b8035906020019184600183028401116401000000008311171561046a57600080fd5b509092509050611168565b6101586004803603606081101561048b57600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160208101358216916040909101351661125e565b610158600480360360408110156104d057600080fd5b5073ffffffffffffffffffffffffffffffffffffffff813581169160200135166113b9565b6101586004803603602081101561050b57600080fd5b81019060208101813564010000000081111561052657600080fd5b82018360208201111561053857600080fd5b8035906020019184602083028401116401000000008311171561055a57600080fd5b50909250905061151d565b6101586004803603608081101561057b57600080fd5b81019060208101813564010000000081111561059657600080fd5b8201836020820111156105a857600080fd5b803590602001918460208302840111640100000000831117156105ca57600080fd5b9193909290916020810190356401000000008111156105e857600080fd5b8201836020820111156105fa57600080fd5b8035906020019184602083028401116401000000008311171561061c57600080fd5b91939092909160208101903564010000000081111561063a57600080fd5b82018360208201111561064c57600080fd5b8035906020019184602083028401116401000000008311171561066e57600080fd5b91939092909160208101903564010000000081111561068c57600080fd5b82018360208201111561069e57600080fd5b803590602001918460208302840111640100000000831117156106c057600080fd5b509092509050611616565b610158600480360360208110156106e157600080fd5b503573ffffffffffffffffffffffffffffffffffffffff1661181c565b610706611064565b610757576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f0aeeca0000000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff90921691630aeeca009160248082019260009290919082900301818387803b1580156107ca57600080fd5b505af11580156107de573d6000803e3d6000fd5b5050505050565b60675473ffffffffffffffffffffffffffffffffffffffff16331461083b5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f18f628d400000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff8c81166004830152602482018c9052604482018b9052606482018a90526084820189905260a4820188905260c4820187905260e482018690526101048201859052915191909216916318f628d49161012480830192600092919082900301818387803b1580156108e957600080fd5b505af11580156108fd573d6000803e3d6000fd5b50505050505050505050505050565b60675473ffffffffffffffffffffffffffffffffffffffff1633146109625760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606654604080517f1e702f830000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff90921691631e702f839160448082019260009290919082900301818387803b1580156109dc57600080fd5b505af11580156109f0573d6000803e3d6000fd5b505050505050565b60665473ffffffffffffffffffffffffffffffffffffffff163314610a64576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517f242a6e3f00000000000000000000000000000000000000000000000000000000815260048101868152602482019283526044820185905273ffffffffffffffffffffffffffffffffffffffff9093169263242a6e3f928792879287929091606401848480828437600081840152601f19601f820116905080830192505050945050505050600060405180830381600087803b158015610b0a57600080fd5b505af1158015610b1e573d6000803e3d6000fd5b50505050505050565b610b2f611064565b610b80576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517f267ab44600000000000000000000000000000000000000000000000000000000815260048101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163267ab4469160248082019260009290919082900301818387803b1580156107ca57600080fd5b610bfb611064565b610c4c576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b606754604080517fda7fc24f00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff84811660048301529151919092169163da7fc24f91602480830192600092919082900301818387803b1580156107ca57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff163314610d165760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16634feb92f38a8a8a8a8a8a8a8a8a6040518a63ffffffff1660e01b8152600401808a73ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001898152602001806020018781526020018681526020018581526020018481526020018381526020018281038252898982818152602001925080828437600081840152601f19601f8201169050808301925050509a5050505050505050505050600060405180830381600087803b1580156108e957600080fd5b60665473ffffffffffffffffffffffffffffffffffffffff163314610e87576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff838116911614610ee05760405162461bcd60e51b8152600401808060200182810382526021815260200180611b8c6021913960400191505060405180910390fd5b60675473ffffffffffffffffffffffffffffffffffffffff9081169063e30443bc908490610f17908216318563ffffffff61188116565b6040518363ffffffff1660e01b8152600401808373ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200182815260200192505050600060405180830381600087803b1580156109dc57600080fd5b610f88611064565b610fd9576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60335460405160009173ffffffffffffffffffffffffffffffffffffffff16907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908390a3603380547fffffffffffffffffffffffff0000000000000000000000000000000000000000169055565b60335473ffffffffffffffffffffffffffffffffffffffff1690565b60335473ffffffffffffffffffffffffffffffffffffffff16331490565b60665473ffffffffffffffffffffffffffffffffffffffff1633146110ee576040805162461bcd60e51b815260206004820152601e60248201527f63616c6c6572206973206e6f74207468652053464320636f6e74726163740000604482015290519081900360640190fd5b606754604080517fa4066fbe0000000000000000000000000000000000000000000000000000000081526004810185905260248101849052905173ffffffffffffffffffffffffffffffffffffffff9092169163a4066fbe9160448082019260009290919082900301818387803b1580156109dc57600080fd5b611170611064565b6111c1576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b6067546040517fb9cc6b1c0000000000000000000000000000000000000000000000000000000081526020600482019081526024820184905273ffffffffffffffffffffffffffffffffffffffff9092169163b9cc6b1c91859185918190604401848480828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b600054610100900460ff168061127757506112776118e2565b80611285575060005460ff16155b6112c05760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff1615801561132657600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b61132f826118e8565b6067805473ffffffffffffffffffffffffffffffffffffffff8086167fffffffffffffffffffffffff000000000000000000000000000000000000000092831617909255606680549287169290911691909117905580156113b357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b50505050565b6113c1611064565b611412576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b60665473ffffffffffffffffffffffffffffffffffffffff83811691161480611450575073ffffffffffffffffffffffffffffffffffffffff821630145b6114a1576040805162461bcd60e51b815260206004820152601760248201527f6e6f7420534643206f722073656c662061646472657373000000000000000000604482015290519081900360640190fd5b606754604080517fd6a0c7af00000000000000000000000000000000000000000000000000000000815273ffffffffffffffffffffffffffffffffffffffff858116600483015284811660248301529151919092169163d6a0c7af91604480830192600092919082900301818387803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff1633146115735760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b6066546040517fe08d7e660000000000000000000000000000000000000000000000000000000081526020600482018181526024830185905273ffffffffffffffffffffffffffffffffffffffff9093169263e08d7e6692869286929182916044909101908590850280828437600081840152601f19601f8201169050808301925050509350505050600060405180830381600087803b1580156109dc57600080fd5b60675473ffffffffffffffffffffffffffffffffffffffff16331461166c5760405162461bcd60e51b8152600401808060200182810382526025815260200180611bad6025913960400191505060405180910390fd5b606660009054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff1663ebdf104c89898989898989896040518963ffffffff1660e01b8152600401808060200180602001806020018060200185810385528d8d82818152602001925060200280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe01690910186810385528b8152602090810191508c908c0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038452898152602090810191508a908a0280828437600083820152601f017fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffe0169091018681038352878152602090810191508890880280828437600081840152601f19601f8201169050808301925050509c50505050505050505050505050600060405180830381600087803b1580156117fa57600080fd5b505af115801561180e573d6000803e3d6000fd5b505050505050505050505050565b611824611064565b611875576040805162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604482015290519081900360640190fd5b61187e81611a57565b50565b6000828201838110156118db576040805162461bcd60e51b815260206004820152601b60248201527f536166654d6174683a206164646974696f6e206f766572666c6f770000000000604482015290519081900360640190fd5b9392505050565b303b1590565b600054610100900460ff168061190157506119016118e2565b8061190f575060005460ff16155b61194a5760405162461bcd60e51b815260040180806020018281038252602e815260200180611b5e602e913960400191505060405180910390fd5b600054610100900460ff161580156119b057600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff007fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff909116610100171660011790555b603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff84811691909117918290556040519116906000907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a38015611a5357600080547fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff00ff1690555b5050565b73ffffffffffffffffffffffffffffffffffffffff8116611aa95760405162461bcd60e51b8152600401808060200182810382526026815260200180611b386026913960400191505060405180910390fd5b60335460405173ffffffffffffffffffffffffffffffffffffffff8084169216907f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e090600090a3603380547fffffffffffffffffffffffff00000000000000000000000000000000000000001673ffffffffffffffffffffffffffffffffffffffff9290921691909117905556fe4f776e61626c653a206e6577206f776e657220697320746865207a65726f2061646472657373436f6e747261637420696e7374616e63652068617320616c7265616479206265656e20696e697469616c697a6564726563697069656e74206973206e6f74207468652053464320636f6e747261637463616c6c6572206973206e6f7420746865204e6f646544726976657220636f6e7472616374a265627a7a723158207237d87a37d8963872e44c451691e26d2b3ebb173dc9eb1a280b145916c36e7464736f6c634300050c0032")
}

// ContractAddress is the NodeDriverAuth contract address
var ContractAddress = common.HexToAddress("0xd100ae0000000000000000000000000000000000")
