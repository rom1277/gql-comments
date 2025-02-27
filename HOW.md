1) создал greph/schema.graphqls, написал схему 
2) добавил зависимости
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/graphql/handler
go get github.com/99designs/gqlgen/graphql/playground
3) определил структуру и методы storage/in_memory.go
4) написал cmd/main.go 
5) создаем файл gqlgen.yml - настройка процесса генерации кода
6) генерируем код go run github.com/99designs/gqlgen generate
7) прописываем graph/resolver.go




реализуй