package main

import (
	"flag"
	"fmt"
	"github.com/adamdecaf/vault-backend-migrator/cmd"
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

	// Import
	if im != nil && *im != "" {
		config := cmd.NewConfig(im, address, file)
		if config == nil {
			exit()
		}
		err := cmd.Import(*config)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Export
	if ex != nil && *ex != "" {
		config := cmd.NewConfig(ex, address, file)
		if config == nil {
			exit()
		}
		err := cmd.Export(*config)
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

func exit() {
	fmt.Println("There was an error reading your config flags, please fix")
	flag.Usage()
	os.Exit(1)
}
