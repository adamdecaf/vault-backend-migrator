package cmd

import (
	"errors"
	"fmt"
	"github.com/adamdecaf/vault-backend-migrator/vault"
	"path"
	"strings"
)

func Export(path, file string) error {
	v, err := vault.NewClient()
	if v == nil || err != nil {
		if err != nil {
			return err
		}
		return errors.New("Unable to create vault client")
	}

	// Make sure path has a trailing slash
	if !strings.HasSuffix(path, "/") {
		path = path + "/"
	}

	// Get all nested keys
	var all []string
	accumulate(&all, *v, path)

	// Read each key's value

	return nil
}

func accumulate(acc *[]string, v vault.Vault, p string) {
	if strings.HasSuffix(p, "/") {
		// Another level exists!
		res := v.List(p)
		if res == nil {
			fmt.Printf("nil, p=%v\n", p)
			return
		}
		for _,k := range *res {
			accumulate(acc, v, path.Join(p, k))
		}
	} else {
		*acc = append(*acc, p)
	}
}
