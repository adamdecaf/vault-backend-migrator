package main

import (
	"flag"
	"fmt"
	"github.com/adamdecaf/vault-backend-migrator/cmd"
	"github.com/hashicorp/vault/api"
	"strings"
	"os"
)

var (
	// Actions
	ex = flag.String("export", "", "The type of data to export")
	im = flag.String("import", "", "The type of data to import")

	// Required during export or import
	address = flag.String("address", "", "The address of vault")
	file = flag.String("file", "", "The local file location to use")

	// Output the version
	version = flag.Bool("version", false, "Output the version number")
)

const Version = "0.0.1-dev"

func main() {
	flag.Parse()

	// Read VAULT_ADDR if '-address' is empty since we're about to check it.
	if empty(address) {
		v := os.Getenv(api.EnvVaultCACert)
		address = &v
	}

	// Import
	if im != nil && *im != "" {
		if empty(address, file) {
			exit()
		}
		err := cmd.Import(*address, *file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Export
	if ex != nil && *ex != "" {
		if empty(address, file) {
			exit()
		}
		err := cmd.Export(*address, *file)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Version
	if version != nil && *version {
		fmt.Println(Version)
		os.Exit(1)
	}

	// No commands, print help.
	flag.Usage()
	os.Exit(1)
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

func exit() {
	fmt.Println("There was an error reading your config flags, please fix")
	flag.Usage()
	os.Exit(1)
}
