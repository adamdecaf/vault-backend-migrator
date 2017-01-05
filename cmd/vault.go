package cmd

import (
	"github.com/hashicorp/vault/api"
)

func newVault(config Config) (*api.Client, error) {
	cfg := api.DefaultConfig()

	// Read vault env variables
	cfg.ReadEnvironment()

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	if err = client.SetAddress(config.Address); err != nil {
		return nil, err
	}
	return client, nil
}
