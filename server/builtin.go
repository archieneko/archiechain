package server

import (
	"github.com/archieneko/archiechain/consensus"
	consensusDev "github.com/archieneko/archiechain/consensus/dev"
	consensusDummy "github.com/archieneko/archiechain/consensus/dummy"
	consensusIBFT "github.com/archieneko/archiechain/consensus/ibft"
	"github.com/archieneko/archiechain/secrets"
	"github.com/archieneko/archiechain/secrets/awsssm"
	"github.com/archieneko/archiechain/secrets/gcpssm"
	"github.com/archieneko/archiechain/secrets/hashicorpvault"
	"github.com/archieneko/archiechain/secrets/local"
)

type ConsensusType string

const (
	DevConsensus   ConsensusType = "dev"
	IBFTConsensus  ConsensusType = "ibft"
	DummyConsensus ConsensusType = "dummy"
)

var consensusBackends = map[ConsensusType]consensus.Factory{
	DevConsensus:   consensusDev.Factory,
	IBFTConsensus:  consensusIBFT.Factory,
	DummyConsensus: consensusDummy.Factory,
}

// secretsManagerBackends defines the SecretManager factories for different
// secret management solutions
var secretsManagerBackends = map[secrets.SecretsManagerType]secrets.SecretsManagerFactory{
	secrets.Local:          local.SecretsManagerFactory,
	secrets.HashicorpVault: hashicorpvault.SecretsManagerFactory,
	secrets.AWSSSM:         awsssm.SecretsManagerFactory,
	secrets.GCPSSM:         gcpssm.SecretsManagerFactory,
}

func ConsensusSupported(value string) bool {
	_, ok := consensusBackends[ConsensusType(value)]

	return ok
}
