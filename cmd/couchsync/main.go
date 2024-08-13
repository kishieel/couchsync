package main

import (
	"fmt"
	"github.com/kishieel/couchdb-sync/pkg/couchsync"
	"github.com/spf13/cobra"
	"os"
)

func main() {
	var rootCmd = &cobra.Command{}
	var config = couchsync.Config{}

	rootCmd.Use = "couchsync"
	rootCmd.Short = "CouchSync is a CLI utility to manage CouchDB documents based on a file system directory structure."
	rootCmd.Version = "1.0.1"

	var examples = []string{
		"couchsync -a http://localhost:5984 -s /path/to/directory -u admin -p admin",
		//"@todo: couchsync -a http://localhost:5984 -s /path/to/directory -u admin --password-stdin",
		//"@todo: couchsync -a http://localhost:5984 -s /path/to/directory --access-token eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJzdWIiOiJhZG1pbiJ9.TGGTTHuuGpEU8WgobXxkrBtW3NiR3dgw5LR-1DEW3BQ",
	}
	rootCmd.Example = fmt.Sprintf("  %s\n", examples[0])

	rootCmd.Flags().StringVarP(&config.CouchdbAddress, "address", "a", "", "address of the CouchDB instance.")
	rootCmd.Flags().StringVarP(&config.CouchdbUsername, "username", "u", "", "username to use when connecting to the host.")
	rootCmd.Flags().StringVarP(&config.CouchdbPassword, "password", "p", "", "password to use when connecting to the host.")
	rootCmd.Flags().StringVarP(&config.DocumentSource, "source", "s", "", "path to the directory containing the documents structure.")

	rootCmd.PreRun = func(cmd *cobra.Command, args []string) {
		couchsync.ValidateConfig(config)
	}

	rootCmd.Run = func(cmd *cobra.Command, args []string) {
		couchsync.SynchronizeDocuments(config)
	}

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
