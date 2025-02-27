package main

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"gql-comments/graph"
	"gql-comments/storage"
	"log"
	"net/http"
	"os"
)

func main() {
	// можно задать порт, на котором запустится программа  PORT=9090 go run cmd/main.go
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	// Инициализация in-memory хранилища
	inMemoryStorage := storage.NewInMemoryStorage()

	// Создание GraphQL сервера
	// srv := handler.GraphQL(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Storage: inMemoryStorage}})) //! новая версия, надо разобраться

	// graph.NewExecutableSchema создаёт исполняемую схему GraphQL.
	// handler.NewDefaultServer оборачивает схему в HTTP-обработчик.
	srv := handler.NewDefaultServer(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{Storage: inMemoryStorage}}))
	//http.Handle — связывает определенный путь ("/") с обработчиком запросов.
	// playground.Handler("GraphQL playground", "/query") — это обработчик, который предоставляет интерактивный веб-интерфейс (GraphQL Playground) для тестирования GraphQL-запросов.
	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	//Здесь мы связываем путь "/query" с обработчиком srv, который отвечает за выполнение фактических GraphQL-запросов.
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)

	log.Fatal(http.ListenAndServe(":"+port, nil))
	// http.ListenAndServe — это функция, которая запускает HTTP-сервер.
	// Аргумент ":" + port указывает адрес и порт, на котором сервер будет слушать входящие соединения.
	// (nil) указывает, что будут использованы глобальные маршруты, зарегистрированные через http.Handle.
}
