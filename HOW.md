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



примеры запросов:

# создать пост:
mutation {
  createPost(input: { user: "JohnDoe", title: "My First Post", content: "This is the content of my first post.", allowComments: true }) {
    id
    title
    content
    allowComments
    createdAt
  }
}


# Создать коммент
mutation {
  createComment(input: { postID: 1, user: "PIsa popa", text: "This is a great post!" }) {
    id
    user
    postID
    text
    createdAt  
   }
}

# вывести комментарии по postID
query {
  comments(postID: 1) {
    id
    user
    text
    createdAt
    postID
    parentID
  }
}



# вывести все посты
query {
  posts {
    id
    title
    content
    allowComments
    createdAt
  }
}

# запрос поста по id
query {
  post(id: 88) {
    id
    title
    content
    user
    allowComments
    createdAt
    comments {
      id
      user
      text
      createdAt
      replies {
        id
        user
        text
        createdAt
      }
    }
  }
}



#

#

#

#