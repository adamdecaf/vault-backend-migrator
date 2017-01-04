package cmd

import(
	"strings"
)

type Config struct {
	DataType string
	Address string
	File string
}

func NewConfig(datatype, address, file *string) *Config {
	if empty(datatype, address, file) {
		return nil
	}
	return &Config{
		DataType: *datatype,
		Address: *address,
		File: *file,
	}
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
