package generator

import (
	"crypto/ecdsa"
	"encoding/json"
	"math/big"
	"os"

	"github.com/archieneko/archiechain/types"
	"github.com/umbracle/ethgo"
	"github.com/umbracle/ethgo/abi"
)

type TransactionGenerator interface {
	GenerateTransaction() (*types.Transaction, error)
	GetExampleTransaction() (*types.Transaction, error)
	GetTransactionErrors() []*FailedTxnInfo
	SetGasEstimate(gasEstimate uint64)
	MarkFailedTxn(failedTxn *FailedTxnInfo)
}

type ContractTxnGenerator interface {
	TransactionGenerator
	MarkFailedContractTxn(failedContractTxn *FailedContractTxnInfo)
	SetContractAddress(types.Address)
}

type TxnErrorType string

const (
	ReceiptErrorType   TxnErrorType = "ReceiptErrorType"
	AddErrorType       TxnErrorType = "AddErrorType"
	ContractDeployType TxnErrorType = "ContractDeployErrorType"
)

const (
	//nolint:lll
	DefaultContractBytecode = "60806040523480156200001157600080fd5b506040516200107a3803806200107a833981810160405281019062000037919062000278565b620000676040518060600160405280602281526020016200105860229139826200008760201b620003781760201c565b80600090805190602001906200007f92919062000156565b5050620004c5565b620001298282604051602401620000a0929190620002fe565b6040516020818303038152906040527f4b5c4277000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506200012d60201b60201c565b5050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b8280546200016490620003ea565b90600052602060002090601f016020900481019282620001885760008555620001d4565b82601f10620001a357805160ff1916838001178555620001d4565b82800160010185558215620001d4579182015b82811115620001d3578251825591602001919060010190620001b6565b5b509050620001e39190620001e7565b5090565b5b8082111562000202576000816000905550600101620001e8565b5090565b60006200021d620002178462000362565b62000339565b9050828152602081018484840111156200023657600080fd5b62000243848285620003b4565b509392505050565b600082601f8301126200025d57600080fd5b81516200026f84826020860162000206565b91505092915050565b6000602082840312156200028b57600080fd5b600082015167ffffffffffffffff811115620002a657600080fd5b620002b4848285016200024b565b91505092915050565b6000620002ca8262000398565b620002d68185620003a3565b9350620002e8818560208601620003b4565b620002f381620004b4565b840191505092915050565b600060408201905081810360008301526200031a8185620002bd565b90508181036020830152620003308184620002bd565b90509392505050565b60006200034562000358565b905062000353828262000420565b919050565b6000604051905090565b600067ffffffffffffffff82111562000380576200037f62000485565b5b6200038b82620004b4565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60005b83811015620003d4578082015181840152602081019050620003b7565b83811115620003e4576000848401525b50505050565b600060028204905060018216806200040357607f821691505b602082108114156200041a576200041962000456565b5b50919050565b6200042b82620004b4565b810181811067ffffffffffffffff821117156200044d576200044c62000485565b5b80604052505050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b610b8380620004d56000396000f3fe608060405234801561001057600080fd5b50600436106100625760003560e01c80632802348a146100675780633fd50a271461008357806361bc221a1461008d57806363b3105d146100ab578063a4136862146100b5578063cfae3217146100d1575b600080fd5b610081600480360381019061007c9190610611565b6100ef565b005b61008b610137565b005b6100956101a8565b6040516100a29190610809565b60405180910390f35b6100b36101ae565b005b6100cf60048036038101906100ca919061064d565b61021f565b005b6100d96102e6565b6040516100e69190610764565b60405180910390f35b80600260008473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff168152602001908152602001600020819055505050565b60005b602881101561016a57600180546101519190610896565b600181905550808061016290610a15565b91505061013a565b507f68db6a931af790fce1a961d4f878434db37a222949cde6f246cf3d7e9e09a6916001600260405161019e929190610712565b60405180910390a1565b60015481565b60005b605081101561021c57600180546101c89190610896565b6001819055507f68db6a931af790fce1a961d4f878434db37a222949cde6f246cf3d7e9e09a6916003600460405161020192919061073b565b60405180910390a1808061021490610a15565b9150506101b1565b50565b6102cc604051806060016040528060238152602001610b2b6023913960008054610248906109b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610274906109b2565b80156102c15780601f10610296576101008083540402835291602001916102c1565b820191906000526020600020905b8154815290600101906020018083116102a457829003601f168201915b505050505083610414565b80600090805190602001906102e29291906104dc565b5050565b6060600080546102f5906109b2565b80601f0160208091040260200160405190810160405280929190818152602001828054610321906109b2565b801561036e5780601f106103435761010080835404028352916020019161036e565b820191906000526020600020905b81548152906001019060200180831161035157829003601f168201915b5050505050905090565b610410828260405160240161038e929190610786565b6040516020818303038152906040527f4b5c4277000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506104b3565b5050565b6104ae83838360405160240161042c939291906107bd565b6040516020818303038152906040527f2ced7cef000000000000000000000000000000000000000000000000000000007bffffffffffffffffffffffffffffffffffffffffffffffffffffffff19166020820180517bffffffffffffffffffffffffffffffffffffffffffffffffffffffff83818316178352505050506104b3565b505050565b60008151905060006a636f6e736f6c652e6c6f679050602083016000808483855afa5050505050565b8280546104e8906109b2565b90600052602060002090601f01602090048101928261050a5760008555610551565b82601f1061052357805160ff1916838001178555610551565b82800160010185558215610551579182015b82811115610550578251825591602001919060010190610535565b5b50905061055e9190610562565b5090565b5b8082111561057b576000816000905550600101610563565b5090565b600061059261058d84610849565b610824565b9050828152602081018484840111156105aa57600080fd5b6105b5848285610970565b509392505050565b6000813590506105cc81610afc565b92915050565b600082601f8301126105e357600080fd5b81356105f384826020860161057f565b91505092915050565b60008135905061060b81610b13565b92915050565b6000806040838503121561062457600080fd5b6000610632858286016105bd565b9250506020610643858286016105fc565b9150509250929050565b60006020828403121561065f57600080fd5b600082013567ffffffffffffffff81111561067957600080fd5b610685848285016105d2565b91505092915050565b61069781610928565b82525050565b6106a68161093a565b82525050565b6106b58161094c565b82525050565b6106c48161095e565b82525050565b60006106d58261087a565b6106df8185610885565b93506106ef81856020860161097f565b6106f881610aeb565b840191505092915050565b61070c8161091e565b82525050565b6000604082019050610727600083018561068e565b610734602083018461069d565b9392505050565b600060408201905061075060008301856106ac565b61075d60208301846106bb565b9392505050565b6000602082019050818103600083015261077e81846106ca565b905092915050565b600060408201905081810360008301526107a081856106ca565b905081810360208301526107b481846106ca565b90509392505050565b600060608201905081810360008301526107d781866106ca565b905081810360208301526107eb81856106ca565b905081810360408301526107ff81846106ca565b9050949350505050565b600060208201905061081e6000830184610703565b92915050565b600061082e61083f565b905061083a82826109e4565b919050565b6000604051905090565b600067ffffffffffffffff82111561086457610863610abc565b5b61086d82610aeb565b9050602081019050919050565b600081519050919050565b600082825260208201905092915050565b60006108a18261091e565b91506108ac8361091e565b9250827fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff038211156108e1576108e0610a5e565b5b828201905092915050565b60006108f7826108fe565b9050919050565b600073ffffffffffffffffffffffffffffffffffffffff82169050919050565b6000819050919050565b60006109338261091e565b9050919050565b60006109458261091e565b9050919050565b60006109578261091e565b9050919050565b60006109698261091e565b9050919050565b82818337600083830152505050565b60005b8381101561099d578082015181840152602081019050610982565b838111156109ac576000848401525b50505050565b600060028204905060018216806109ca57607f821691505b602082108114156109de576109dd610a8d565b5b50919050565b6109ed82610aeb565b810181811067ffffffffffffffff82111715610a0c57610a0b610abc565b5b80604052505050565b6000610a208261091e565b91507fffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffffff821415610a5357610a52610a5e565b5b600182019050919050565b7f4e487b7100000000000000000000000000000000000000000000000000000000600052601160045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052602260045260246000fd5b7f4e487b7100000000000000000000000000000000000000000000000000000000600052604160045260246000fd5b6000601f19601f8301169050919050565b610b05816108ec565b8114610b1057600080fd5b50565b610b1c8161091e565b8114610b2757600080fd5b5056fe4368616e67696e67206772656574696e672066726f6d202725732720746f2027257327a2646970667358221220a4f650513c5038cc3c919651c8ca206cda3ae09ac470aebc8d96da7b43318d5e64736f6c634300080400334465706c6f79696e67206120477265657465722077697468206772656574696e673a"
)

type ContractArtifact struct {
	Bytecode string   `json:"bytecode"`
	ABI      *abi.ABI `json:"abi"`
}

type TxnError struct {
	Error     error
	ErrorType TxnErrorType
}

type FailedTxnInfo struct {
	Index  uint64
	TxHash string
	Error  *TxnError
}

type FailedContractTxnInfo struct {
	TxHash string
	Error  *TxnError
}

type GeneratorParams struct {
	Nonce            uint64
	ChainID          uint64
	SenderAddress    types.Address
	RecieverAddress  types.Address
	SenderKey        *ecdsa.PrivateKey
	Value            *big.Int
	GasPrice         *big.Int
	ContractArtifact *ContractArtifact
	ConstructorArgs  []byte // smart contract constructor arguments
	ContractAddress  ethgo.Address
}

// ReadContractArtifact reads the contract bytecode from the specified path
func ReadContractArtifact(configPath string) (*ContractArtifact, error) {
	rawData, readErr := os.ReadFile(configPath)
	if readErr != nil {
		return nil, readErr
	}

	var artifact ContractArtifact
	if jsonErr := json.Unmarshal(rawData, &artifact); jsonErr != nil {
		return nil, jsonErr
	}

	return &artifact, nil
}
