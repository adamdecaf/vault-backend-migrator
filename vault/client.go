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

// List the keys at at given vault path. This has only been tested on the generic backend.
// It will return nil if something goes wrong.
func (v *Vault) List(path string) *[]string {
	secret, err := v.c.Logical().List(path)
	if secret == nil || err != nil {
		if err == nil {
			fmt.Println("Unable to read path, does it exist?")
		}
		fmt.Printf("Error reading secrets, err=%v\n", err)
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

// Read accepts a vault path to read the data out of. It will return a pointer to
// a base64 encoded string representing the secret's data.
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

// Write takes in a vault path and base64 encoded data to be written at that path.
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
