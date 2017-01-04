package cmd

import (
	"fmt"
)

func Import(config Config) error {
	fmt.Printf("Import - %s", config.DataType)
	return nil
}
