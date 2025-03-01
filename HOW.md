1) создал greph/schema.graphqls, определил схему данных
2) добавил зависимости
go get github.com/99designs/gqlgen
go get github.com/99designs/gqlgen/graphql/handler
go get github.com/99designs/gqlgen/graphql/playground
3) определил структуру и методы storage/in_memory.go
4) написал cmd/main.go 
5) создаем файл gqlgen.yml - настройка процесса генерации кода
6) генерируем код 
go run github.com/99designs/gqlgen generate
7) прописываем graph/resolver.go

можно задать порт, на котором запустится программа  PORT=9090 go run cmd/main.go


примеры запросов:
# # создаём пост
# mutation {
#   createPost(input: { user: "JohnDoe", title: "My First Post", content: "This is the content of my first post.", allowComments: true }) {
#     id
#     title
#     content
#     allowComments
#     createdAt
#   }
# }
# # создаем коммент под постом
# mutation {
#   createComment(input: { postID: 1, user: "zzzzzzzzzz", text: "yuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuuu" }) {
#     id
#     user
#     postID
#     text
#     createdAt  
#    }
# }

# # отвечаем на коммент
# mutation {
#   createComment(input: { postID: 1, parentID: 101, user: "XXXX", text: "dadasdasdas comment 102" }) {
#     id
#     user
#     text
#     createdAt
#     parentID
#   }
# }

    
# # #  вывод поста с комментариями
# query {
#   comments(postID: 1, limit: 20, offset: 0) {
#     id
#     user
#     text
#     parentID
#     createdAt
#     replies {
#       id
#       user
#       text
#       createdAt
#       replies {
#         id
#         user
#         text
#         createdAt
#       }
#     }
#   }
# }
    
# query {
#   comments(postID: 1) {
#     id
#     user
#     text
#     createdAt
#     postID
#     parentID

#   }
# }


# закрыть коммент
# mutation {
#   closeCommentsPost(user: "JohnDoe", postID: 1, commentsAllowed: false){
#         id
#     title
#     content
#     allowComments
#     createdAt 
#   }
# }

# subscription {
#   commentAdded(postID: 1) {
#     id
#     user
#     text
#     createdAt
#   }
# }