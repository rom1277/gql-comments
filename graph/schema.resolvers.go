package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.
// Code generated by github.com/99designs/gqlgen version v0.17.66

import (
	"context"
	"fmt"
	// "gql-comments/graph"
	"gql-comments/storage"
)

// CreatePost is the resolver for the createPost field.
func (r *mutationResolver) CreatePost(ctx context.Context, title string, content string, allowComments bool) (*storage.Post, error) {
	panic(fmt.Errorf("not implemented: CreatePost - createPost"))
}

// ApproveComments is the resolver for the approveComments field.
func (r *postResolver) ApproveComments(ctx context.Context, obj *storage.Post) (bool, error) {
	panic(fmt.Errorf("not implemented: ApproveComments - approveComments"))
}

// Posts is the resolver for the posts field.
func (r *queryResolver) Posts(ctx context.Context) ([]*storage.Post, error) {
	panic(fmt.Errorf("not implemented: Posts - posts"))
}

// Mutation returns MutationResolver implementation.
func (r *Resolver) Mutation() MutationResolver { return &mutationResolver{r} }

// Post returns PostResolver implementation.
func (r *Resolver) Post() PostResolver { return &postResolver{r} }

// Query returns QueryResolver implementation.
func (r *Resolver) Query() QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type postResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
