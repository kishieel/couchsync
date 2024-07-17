package couchsync

import (
	"encoding/json"
	"fmt"
)

type CreateOrUpdateDocumentArgs struct {
	databaseName    string
	documentName    string
	documentContent map[string]any
	config          Config
}

func CreateOrUpdateDocument(args CreateOrUpdateDocumentArgs) {
	currentDocument := FetchDocument(FetchDocumentArgs{args.databaseName, args.documentName, args.config})

	if currentDocument == nil {
		err := CreateDocument(CreateDocumentArgs{args.databaseName, args.documentName, args.documentContent, args.config})

		if err != nil {
			fmt.Printf("[✗] Document %s could not be created for %s.\n", args.documentName, args.databaseName)
		} else {
			fmt.Printf("[✓] Document %s created for %s.\n", args.documentName, args.databaseName)
		}

		return
	}

	currentDocumentRevision := currentDocument["_rev"].(string)
	delete(currentDocument, "_rev")

	currentDocumentBytes, _ := json.Marshal(currentDocument)
	newestDocumentBytes, _ := json.Marshal(args.documentContent)

	documentChanged := string(currentDocumentBytes) != string(newestDocumentBytes)

	if documentChanged {
		err := UpdateDocument(UpdateDocumentArgs{args.databaseName, args.documentName, currentDocumentRevision, args.documentContent, args.config})

		if err != nil {
			fmt.Printf("[✗] Document %s could not be updated for %s.\n", args.documentName, args.databaseName)
		} else {
			fmt.Printf("[✓] Document %s updated for %s.\n", args.documentName, args.databaseName)
		}

		return
	}

	fmt.Printf("[✓] Document %s for %s is up to date.\n", args.documentName, args.databaseName)
}
