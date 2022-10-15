package command

import "github.com/archieneko/archiechain/server"

const (
	DefaultGenesisFileName = "genesis.json"
	DefaultChainName       = "archie"
	DefaultChainID         = 1243
	DefaultPremineBalance  = "0x204FCE5E3E25026110000000" // 10 billion units of native network currency
	DefaultConsensus       = server.IBFTConsensus
	DefaultGenesisGasUsed  = 458752  // 0x70000
	DefaultGenesisGasLimit = 5242880 // 0x500000
)

const (
	JSONOutputFlag  = "json"
	GRPCAddressFlag = "grpc-address"
	JSONRPCFlag     = "jsonrpc"
)

// GRPCAddressFlagLEGACY Legacy flag that needs to be present to preserve backwards
// compatibility with running clients
const (
	GRPCAddressFlagLEGACY = "grpc"
)
