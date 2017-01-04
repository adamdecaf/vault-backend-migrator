package cmd

import (
	"fmt"
)

func Export(config Config) error {
	fmt.Printf("Export - %s", config.DataType)
	return nil
}
