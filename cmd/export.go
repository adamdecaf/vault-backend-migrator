package cmd

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/adamdecaf/vault-backend-migrator/vault"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"strings"
)

const (
	OutputFileMode = 0644
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
	fmt.Printf("Reading all keys under %s\n", path)
	var all []string
	accumulate(&all, *v, path)

	// Read each key's value
	fmt.Println("Reading all secrets")
	var items []Item
	for _, p := range all {
		kvs := v.Read(p)
		if kvs == nil {
			fmt.Printf("invalid read on %s\n", p)
			continue
		}

		var pairs []Pair
		for k, v := range kvs {
			pairs = append(pairs, Pair{Key: k, Value: v})
		}
		items = append(items, Item{Path: p, Pairs: pairs})
	}

	// Convert to json and write to a file
	export := Wrap{Data: items}
	out, err := json.Marshal(&export)
	if err != nil {
		fmt.Println(err)
	}
	abs, err := filepath.Abs(file)
	if err != nil {
		fmt.Println(err)
	}

	// Create the output file if it's not there
	if _, err := os.Stat(abs); err != nil {
		f, err := os.Create(abs)
		defer f.Close()
		if err != nil {
			fmt.Println(err)
		}
	}

	// Write json to the file
	err = ioutil.WriteFile(abs, out, OutputFileMode)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("Wrote %s\n", abs)

	return nil
}

func accumulate(acc *[]string, v vault.Vault, p string) {
	res := v.List(p)
	if res == nil { // We ran into a leaf
		*acc = append(*acc, p)
		return
	}
	for _, k := range res {
		accumulate(acc, v, path.Join(p, k))
	}
}
