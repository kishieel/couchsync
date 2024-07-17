package couchsync

import (
	"fmt"
	"os"
)

func ValidateConfig(config Config) {
	if config.CouchdbAddress == "" {
		fmt.Println("Error: --address option is required")
		os.Exit(1)
	}

	if config.DocumentSource == "" {
		fmt.Println("Error: --source option is required")
		os.Exit(1)
	}
}
