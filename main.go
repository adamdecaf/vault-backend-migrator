package main

import (
	"flag"
	"fmt"
	"os"
	"strings"

	"github.com/adamdecaf/vault-backend-migrator/cmd"
)

var (
	// Actions
	ex = flag.String("export", "", "The vault path to export")
	im = flag.String("import", "", "The vault path to import data into")
	vr = flag.String("ver", "", "KV version")
	md = flag.String("metadata", "", "Metadata path")

	// Required during export or import
	file = flag.String("file", "", "The local file location to use")

	// Output the version
	version = flag.Bool("version", false, "Output the version number")

	// JSON formatting
	formatJson = flag.Bool("pretty", true, "Format JSON output")
)

const Version = "0.2.1-dev"

func main() {
	flag.Parse()

	// Import
	if im != nil && *im != "" {
		if empty(im, file) {
			exit()
		}
		err := cmd.Import(*im, *file, *vr)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
		return
	}

	// Export
	if ex != nil && *ex != "" {
		if empty(ex, file) {
			exit()
		}
		err := cmd.Export(*ex, *file, *md, *vr, *formatJson)
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
	for _, v := range s {
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
