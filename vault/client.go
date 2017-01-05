package vault

import (
	"github.com/hashicorp/vault/api"
)

func NewClient(address string) (*api.Client, error) {
	cfg := api.DefaultConfig()

	// Read vault env variables
	cfg.ReadEnvironment()

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}
	if err = client.SetAddress(address); err != nil {
		return nil, err
	}
	return client, nil
}
