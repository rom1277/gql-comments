postgr:
	docker compose up --build
postgrStop:
	docker compose down
rebuild:
	docker compose up --build --force-recreate
db-up:
	dockercompose up -d db
clean-db:
	rm -rf ./pgdata
test:
	go test graph/resolvers/testResolvers/*.go
	go test storage/inmemory/inmemory_test/*.go
inmem:
	go run cmd/main.go -storage inmemory

db-logs:
	docker logs -f postgres-container