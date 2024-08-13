package couchsync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type CreateDocumentArgs struct {
	databaseName    string
	documentName    string
	documentContent map[string]any
	config          Config
}

func createDocument(args CreateDocumentArgs) error {
	client := &http.Client{}
	req, _ := http.NewRequest("POST", fmt.Sprintf("%s/%s", args.config.CouchdbAddress, args.databaseName), nil)
	req.SetBasicAuth(args.config.CouchdbUsername, args.config.CouchdbPassword)
	req.Header.Set("Content-Type", "application/json")

	documentContent, _ := json.Marshal(args.documentContent)
	req.Body = io.NopCloser(bytes.NewReader(documentContent))

	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("unexpected status code: %d", res.StatusCode)
}
