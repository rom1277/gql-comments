package graph

// This file will not be regenerated automatically.
// //
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct{}

// package graph

// import (
// 	"context"
// 	"errors"
// 	"fmt"
// 	"gql-comments/storage"
// 	"math/rand"
// )

// type Resolver struct {
// 	Storage *storage.InMemoryStorage
// }

// func (r *Resolver) Mutation() MutationResolver {
// 	return &mutationResolver{r}
// }

// func (r *Resolver) Query() QueryResolver {
// 	return &queryResolver{r}
// }

// type mutationResolver struct{ *Resolver }

// func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*storage.Post, error) {
// 	if title == "" || content == "" {
// 		return nil, errors.New("title and content must not be empty")
// 	}

// 	id := fmt.Sprintf("post-%d", rand.Intn(1000000))
// 	post := storage.Post{
// 		ID:            id,
// 		Title:         title,
// 		Content:       content,
// 		AllowComments: allowComments,
// 	}

// 	r.Storage.CreatePost(post)
// 	return &post, nil
// }

// type queryResolver struct{ *Resolver }

// func (r *queryResolver) Posts(ctx context.Context) ([]*storage.Post, error) {
// 	posts := r.Storage.GetAllPosts()
// 	var result []*storage.Post
// 	for i := range posts {
// 		result = append(result, &posts[i])
// 	}
// 	return result, nil
// }
