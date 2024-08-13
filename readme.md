# CouchSync

Managing CouchDB design documents can be a tedious task. JavaScript functions mixed with JSON can be hard to maintain and debug. For this reason, CouchSync was created. CouchSync is a command-line tool that allows you to manage CouchDB design documents in a directory structure. You can keep your functions in separate files and CouchSync will automagically parse and upload them to CouchDB as a JSON document.

### Installation

To install CouchSync you can use the following command.

```shell
wget -nv https://github.com/kishieel/couchsync/raw/master/install.sh -O - | sudo bash
```

### Usage

To use CouchSync you need to prepare a directory structure that represents documents that are supposed to be uploaded to
CouchDB. The example below shows a directory structure that represents a single design document with various functions
for specified database.

```text
docs/
└─ database-name/
   └─ _design/
      └─ document-name/
         ├─ filters/
         │  ├─ filter-1.js
         │  └─ filter-2.js
         ├─ updates/
         │  ├─ update-1.js   
         │  └─ update-2.js
         ├─ views/
         │  ├─ view_1/
         │  │  ├─ map.js
         │  │  └─ reduce.js
         │  └─ view_2/
         │     ├─ map.js
         │     └─ reduce.js
         └─ validate_doc_update.js
        
```

With this directory structure you can use the following command to upload the documents to CouchDB.

```shell
couchsync \
  --address http://localhost:5984 \
  --username username \
  --password password \
  --source docs
```

Consequently, the document with name `_design/document-name` will be created in database with name `database-name` with
following content.

```json
{
  "_id": "_design/document-name",
  "filters": {
    "filter-1": "function(doc, req) { ... }",
    "filter-2": "function(doc, req) { ... }"
  },
  "updates": {
    "update-1": "function(doc, req) { ... }",
    "update-2": "function(doc, req) { ... }"
  },
  "views": {
    "view_1": {
      "map": "function(doc) { ... }",
      "reduce": "function(keys, values, rereduce) { ... }"
    },
    "view_2": {
      "map": "function(doc) { ... }",
      "reduce": "function(keys, values, rereduce) { ... }"
    }
  },
  "validate_doc_update": "function(newDoc, oldDoc, userCtx, secObj) { ... }"
}
```

## Contributing

Please see the [contributing.md](contributing.md) file for details on how to contribute to this project.

## License

This project is licensed under the MIT License - see the [license.md](license.md) file for details.
