.PHONY: clean build couchdb-start couchdb-stop couchdb-logs

clean:
	rm -rdf bin/couchdb-sync

build:
	go build -o ./bin/couchsync ./cmd/couchsync
	chmod +x ./bin/couchsync

couchdb-start:
	docker run -d -p 5984:5984 --rm --name couchdb -e COUCHDB_USER=admin -e COUCHDB_PASSWORD=admin bitnami/couchdb:3.3.3

couchdb-stop:
	docker stop couchdb

couchdb-logs:
	docker logs -f couchdb