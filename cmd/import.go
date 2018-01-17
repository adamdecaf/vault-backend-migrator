package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adamdecaf/vault-backend-migrator/vault"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Import(path, file string) error {
	abs, err := filepath.Abs(file)
	if err != nil {
		return err
	}

	// Check the input file exists
	if _, err := os.Stat(abs); err != nil {
		f, err := os.Create(abs)
		defer f.Close()
		if err != nil {
			return err
		}
	}

	// Read input file
	b, err := ioutil.ReadFile(abs)
	if err != nil {
		return err
	}

	// Parse data
	var wrap Wrap
	err = json.Unmarshal(b, &wrap)
	if err != nil {
		return err
	}

	// Setup vault client
	v, err := vault.NewClient()
	if v == nil || err != nil {
		if err != nil {
			return err
		}
		return errors.New("Unable to create vault client")
	}

	// Write each keypair to vault
	for _, item := range wrap.Data {
		data := make(map[string]string)
		for _, kv := range item.Pairs {
			data[kv.Key] = kv.Value
		}
		fmt.Printf("Writing %s\n", item.Path)
		v.Write(item.Path, data)
	}

	return nil
}
