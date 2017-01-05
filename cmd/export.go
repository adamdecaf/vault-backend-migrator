package cmd

import (
	"errors"
	"fmt"
)

func Export(config Config) error {
	vault, err := newVault(config)
	if vault == nil || err != nil {
		if err != nil {
			return err
		}
		return errors.New("Unable to create vault client")
	}

	// fmt.Printf("%v\n", *vault.Logical())

	// fmt.Println(vault.Address())
	// fmt.Println(vault.Token())

	// vault.Auth()

	// list
	secret, err := vault.Logical().List("secret/banno/")
	if err != nil {
		return fmt.Errorf("Error reading secrets, err=%v", err)
	}
	// if secret == nil {
	// 	fmt.Println("bad")
	// }

	fmt.Printf("A - %v\n", secret)
	// fmt.Printf("AA - %v\n", secret.WrapInfo)

	// read
	secret, err = vault.Logical().Read("secret/banno/config/small-deployable-web-server")
	if err != nil {
		return fmt.Errorf("Error reading secrets, err=%v", err)
	}
	fmt.Printf("B - %v\n", secret)
	return nil
}
