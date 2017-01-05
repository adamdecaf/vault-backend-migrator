package vault

import (
	"fmt"
	"github.com/hashicorp/vault/api"
	"os"
)

type Vault struct {
	c *api.Client
}

func NewClient() (*Vault, error) {
	cfg := api.DefaultConfig()

	// Read vault env variables
	cfg.ReadEnvironment()

	client, err := api.NewClient(cfg)
	if err != nil {
		return nil, err
	}

	// Sanity checks
	if v := os.Getenv(api.EnvVaultAddress); v == "" {
		fmt.Println("Did you mean to use localhost vault? Try setting VAULT_ADDR")
	}

	return &Vault{
		c: client,
	}, nil
}

func (v *Vault) List(path string) *[]string {
	secret, err := v.c.Logical().List(path)
	if secret == nil || err != nil {
		if err == nil {
			fmt.Println("Unable to read path, does it exist?")
		}
		fmt.Println("Error reading secrets, err=%v", err)
		return nil
	}

	r, ok := secret.Data["keys"].([]interface{})
	fmt.Println(ok)
	if ok {
		out := make([]string, len(r))
		for i := range r {
			out[i] = r[i].(string)
		}
		return &out
	}
	return nil

}
