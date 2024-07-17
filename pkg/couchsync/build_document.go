package couchsync

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

type BuildDocumentArgs struct {
	source string
}

func buildDocument(args BuildDocumentArgs) map[string]any {
	documentContent := map[string]any{}
	filesOrDirectories, _ := os.ReadDir(args.source)

	for _, fileOrDirectory := range filesOrDirectories {
		if fileOrDirectory.Name() == "document.json" {
			continue
		}

		if fileOrDirectory.IsDir() {
			propertyContent := buildDocument(BuildDocumentArgs{source: fmt.Sprintf("%s/%s", args.source, fileOrDirectory.Name())})
			propertyName := fileOrDirectory.Name()
			documentContent[propertyName] = propertyContent
		} else {
			fileContent, _ := os.ReadFile(fmt.Sprintf("%s/%s", args.source, fileOrDirectory.Name()))
			propertyName := strings.TrimSuffix(fileOrDirectory.Name(), filepath.Ext(fileOrDirectory.Name()))
			documentContent[propertyName] = string(fileContent)
		}
	}

	jsonContent, _ := os.ReadFile(fmt.Sprintf("%s/document.json", args.source))
	_ = json.Unmarshal(jsonContent, &documentContent)

	return documentContent
}
