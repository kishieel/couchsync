package couchsync

import (
	"fmt"
	"os"
)

func SynchronizeDocuments(config Config) {
	databases, _ := os.ReadDir(config.DocumentSource)

	for _, database := range databases {
		if !database.IsDir() {
			continue
		}

		databaseName := database.Name()
		documentNames := []string{}

		normalDocuments, _ := os.ReadDir(fmt.Sprintf("%s/%s", config.DocumentSource, databaseName))
		designDocuments, _ := os.ReadDir(fmt.Sprintf("%s/%s/_design", config.DocumentSource, databaseName))

		for _, document := range normalDocuments {
			if document.IsDir() && document.Name() != "_design" {
				documentNames = append(documentNames, document.Name())
			}
		}

		for _, document := range designDocuments {
			if document.IsDir() {
				documentNames = append(documentNames, "_design/"+document.Name())
			}
		}

		for _, documentName := range documentNames {
			documentPath := fmt.Sprintf("%s/%s/%s", config.DocumentSource, databaseName, documentName)

			documentContent := buildDocument(BuildDocumentArgs{documentPath})
			documentContent["_id"] = documentName

			createOrUpdateDocument(CreateOrUpdateDocumentArgs{databaseName, documentName, documentContent, config})
		}
	}
}
