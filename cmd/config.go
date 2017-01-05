package cmd

import(
	"github.com/hashicorp/vault/api"
	"strings"
	"os"
)

type Config struct {
	DataType string
	Address string
	File string
}

func NewConfig(datatype, address, file *string) *Config {
	// This is checked before creating a vault client and reading from the env vars
	// for VAULT_ADDR, so check that before failing.
	if empty(address) {
		v := os.Getenv(api.EnvVaultCACert)
		address = &v
	}
	if empty(datatype, address, file) {
		return nil
	}
	return &Config{
		DataType: *datatype,
		Address: *address,
		File: *file,
	}
}

// Do we have any empty strings?
func empty(s ...*string) bool {
	for _,v := range s {
		if v == nil || len(strings.TrimSpace(*v)) == 0 {
			return true
		}
	}
	return false
}
