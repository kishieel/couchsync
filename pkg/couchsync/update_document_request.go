package couchsync

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type UpdateDocumentArgs struct {
	databaseName     string
	documentName     string
	documentRevision string
	documentContent  map[string]any
	config           Config
}

func updateDocument(args UpdateDocumentArgs) error {
	client := &http.Client{}
	req, _ := http.NewRequest("PUT", fmt.Sprintf("%s/%s/%s", args.config.CouchdbAddress, args.databaseName, args.documentName), nil)
	req.SetBasicAuth(args.config.CouchdbUsername, args.config.CouchdbPassword)
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("If-Match", args.documentRevision)

	documentContent, _ := json.Marshal(args.documentContent)
	req.Body = io.NopCloser(bytes.NewReader(documentContent))

	res, _ := client.Do(req)
	defer res.Body.Close()

	if res.StatusCode == http.StatusCreated || res.StatusCode == http.StatusOK {
		return nil
	}

	return fmt.Errorf("unexpected status code: %d", res.StatusCode)
}
