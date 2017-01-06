package vault

import (
	"encoding/base64"
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
	if ok {
		out := make([]string, len(r))
		for i := range r {
			out[i] = r[i].(string)
		}
		return &out
	}
	return nil
}

// todo: note that this returns base64 encoded strings
func (v *Vault) Read(path string) *string {
	s, err := v.c.Logical().Read(path)
	if err != nil {
		fmt.Printf("Error reading secrets, err=%v", err)
		return nil
	}
	r, ok := s.Data["value"].(string)
	if !ok {
		return nil
	}

	// Encode to base64
	e := base64.StdEncoding.EncodeToString([]byte(r))
	return &e
}

// todo: note that this expects base64 encoded data
func (v *Vault) Write(path, data string) error {
	b, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return err
	}

	d := make(map[string]interface{})
	d["value"] = string(b)

	secret, err := v.c.Logical().Write(path, d)
	if secret == nil {
		return fmt.Errorf("No secret returned when writing to %s", path)
	}
	if err != nil {
		return err
	}

	fmt.Println(secret)
	return nil
}
