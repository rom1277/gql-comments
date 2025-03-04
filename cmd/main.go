package main

import (
	"database/sql"
	"flag"
	"fmt"
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/rom1277/gql-comments/graph/generated"
	"github.com/rom1277/gql-comments/graph/resolvers"
	"github.com/rom1277/gql-comments/storage"
	"github.com/rom1277/gql-comments/storage/inmemory"
	"github.com/rom1277/gql-comments/storage/postgres"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	storageType := flag.String("storage", "inmemory", "Storage type: 'inmemory' or 'postgres'")
	flag.Parse()

	var postStorage storage.PostStorage
	var commentStorage storage.CommentStorage
	var notifier storage.Notifier

	switch *storageType {
	case "inmemory":
		postStorage = inmemory.NewInMemoryStoragePost()
		commentStorage = inmemory.NewInMemoryStorageCommenst()
		notifier = inmemory.NewNotifier()
	case "postgres":
		connStr := "postgres://postgres:postgres@db:5432/postgres?sslmode=disable"
		waitForPostgres(connStr)

		var err error

		postStorage, err = postgres.NewPostgresPostStorage(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL for posts: %v", err)
		}
		commentStorage, err = postgres.NewPostgresCommentStorage(connStr)
		if err != nil {
			log.Fatalf("Failed to connect to PostgreSQL for comments: %v", err)
		}
		notifier = inmemory.NewNotifier()
	default:
		log.Fatalf("Unknown storage type: %s", *storageType)
	}

	resolver := resolvers.NewResolver(postStorage, commentStorage, notifier)
	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: resolver}))
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

func waitForPostgres(dsn string) {
	for {
		db, err := sql.Open("postgres", dsn)
		if err == nil {
			err = db.Ping()
			if err == nil {
				fmt.Println("Successfully connected to PostgreSQL!")
				return
			}
		}
		log.Println("Waiting for PostgreSQL to be ready...")
		time.Sleep(5 * time.Second)
	}
}
