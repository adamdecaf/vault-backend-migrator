package cmd

import (
	"encoding/json"
	"fmt"
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

	// Write each keypair to vault
	for _, pair := range wrap.Data {
		fmt.Printf("%s - %d\n", pair.Key, len(pair.Value))
	}

	return nil
}
