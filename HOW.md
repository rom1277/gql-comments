1) создал greph/schema.graphqls, определил схему данных
2) добавил зависимости
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/graphql/handler
go get github.com/99designs/gqlgen/graphql/playground
3) определил структуру и методы storage/in_memory.go
4) написал cmd/main.go 
5) создаем файл gqlgen.yml - настройка процесса генерации кода
6) генерируем код go run github.com/99designs/gqlgen generate
7) прописываем graph/resolver.go



примеры запросов:


mutation {
  createPost(title: "My First Post", content: "This is the content of my first post.", allowComments: true) {
    id
    title
    content
    allowComments
  }
}





query {
  posts {
    id
    title
    content
    allowComments
  }
}



реализуй