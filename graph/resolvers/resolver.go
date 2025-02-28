package resolvers

// This file will not be regenerated automatically.

// It serves as dependency injection for your app, add any dependencies you require here.

import (
	"context"
	"errors"
	"fmt"
	"gql-comments/graph/generated"
	"gql-comments/graph/model"
	"gql-comments/storage/inmemory"
	"gql-comments/structures"
)

type Resolver struct {
	Storage *inmemory.InMemoryStorage
}
type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*structures.Post, error) {
	if input.Title == "" || input.Content == "" {
		return nil, errors.New("title and content must not be empty")
	}
	post := &structures.Post{
		Title:   input.Title,
		User:    input.User,
		Content: input.Content,
	}
	createdPost, err := r.Storage.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*structures.Post, error) {
	posts := r.Storage.GetAllPosts()
	var result []*structures.Post
	for i := range posts {
		result = append(result, &posts[i])
	}
	if len(result) == 0 {
		return nil, errors.New("no added posts")
	}
	return result, nil
}

// CreateComment is the resolver for the createComment field.
func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*structures.Comment, error) {
	panic(fmt.Errorf("not implemented: CreateComment - createComment"))
}

// Post is the resolver for the post field.
func (r *queryResolver) Post(ctx context.Context, id int) (*structures.Post, error) {
	panic(fmt.Errorf("not implemented: Post - post"))
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID int) ([]*structures.Comment, error) {
	panic(fmt.Errorf("not implemented: Comments - comments"))
}

// Mutation returns generated.MutationResolver implementation.
