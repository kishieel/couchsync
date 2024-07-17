package couchsync

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type FetchDocumentArgs struct {
	databaseName string
	documentName string
	config       Config
}

func FetchDocument(args FetchDocumentArgs) map[string]any {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", fmt.Sprintf("%s/%s/%s", args.config.CouchdbAddress, args.databaseName, args.documentName), nil)
	req.SetBasicAuth(args.config.CouchdbUsername, args.config.CouchdbPassword)

	res, _ := client.Do(req)
	defer res.Body.Close()

	var currentDocument map[string]any
	_ = json.NewDecoder(res.Body).Decode(&currentDocument)

	if res.StatusCode != http.StatusOK {
		return nil
	}

	return currentDocument
}
