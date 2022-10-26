package init

import (
	"errors"

	"github.com/archieneko/archiechain/command"
	"github.com/archieneko/archiechain/secrets"
	"github.com/archieneko/archiechain/secrets/helper"
)

const (
	dataDirFlag = "data-dir"
	configFlag  = "config"
	ecdsaFlag   = "ecdsa"
	blsFlag     = "bls"
	networkFlag = "network"
	numFlag     = "num"
)

var (
	errInvalidConfig   = errors.New("invalid secrets configuration")
	errInvalidParams   = errors.New("no config file or data directory passed in")
	errUnsupportedType = errors.New("unsupported secrets manager")
)

type initParams struct {
	dataDir          string
	configPath       string
	generatesECDSA   bool
	generatesBLS     bool
	generatesNetwork bool

	secretsManager secrets.SecretsManager
	secretsConfig  *secrets.SecretsManagerConfig
}

func (ip *initParams) validateFlags() error {
	if ip.dataDir == "" && ip.configPath == "" {
		return errInvalidParams
	}

	return nil
}

func (ip *initParams) initSecrets() error {
	if err := ip.initSecretsManager(); err != nil {
		return err
	}

	if err := ip.initValidatorKey(); err != nil {
		return err
	}

	return ip.initNetworkingKey()
}

func (ip *initParams) initSecretsManager() error {
	var err error
	if ip.hasConfigPath() {
		if err = ip.parseConfig(); err != nil {
			return err
		}

		ip.secretsManager, err = helper.InitCloudSecretsManager(ip.secretsConfig)

		return err
	}

	return ip.initLocalSecretsManager()
}

func (ip *initParams) hasConfigPath() bool {
	return ip.configPath != ""
}

func (ip *initParams) parseConfig() error {
	secretsConfig, readErr := secrets.ReadConfig(ip.configPath)
	if readErr != nil {
		return errInvalidConfig
	}

	if !secrets.SupportedServiceManager(secretsConfig.Type) {
		return errUnsupportedType
	}

	ip.secretsConfig = secretsConfig

	return nil
}

func (ip *initParams) initLocalSecretsManager() error {
	local, err := helper.SetupLocalSecretsManager(ip.dataDir)
	if err != nil {
		return err
	}

	ip.secretsManager = local

	return nil
}

func (ip *initParams) initValidatorKey() error {
	var err error

	if ip.generatesECDSA {
		if _, err = helper.InitECDSAValidatorKey(ip.secretsManager); err != nil {
			return err
		}
	}

	if ip.generatesBLS {
		if _, err = helper.InitBLSValidatorKey(ip.secretsManager); err != nil {
			return err
		}
	}

	return nil
}

func (ip *initParams) initNetworkingKey() error {
	if ip.generatesNetwork {
		if _, err := helper.InitNetworkingPrivateKey(ip.secretsManager); err != nil {
			return err
		}
	}

	return nil
}

// getResult gets keys from secret manager and return result to display
func (ip *initParams) getResult() (command.CommandResult, error) {
	var (
		res = &SecretsInitResult{}
		err error
	)

	if res.Address, err = helper.LoadValidatorAddress(ip.secretsManager); err != nil {
		return nil, err
	}

	if res.BLSPubkey, err = helper.LoadBLSPublicKey(ip.secretsManager); err != nil {
		return nil, err
	}

	if res.NodeID, err = helper.LoadNodeID(ip.secretsManager); err != nil {
		return nil, err
	}

	return res, nil
}