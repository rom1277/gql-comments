package resolvers

import (
	"github.com/rom1277/gql-comments/graph/generated"
	"github.com/rom1277/gql-comments/storage"
)

const (
	ConstOffset = 0
	ConstLimit  = 10
)

type Resolver struct {
	PostStorage    storage.PostStorage
	CommentStorage storage.CommentStorage
	Notifier       storage.Notifier
}

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func NewResolver(postStorage storage.PostStorage, commentStorage storage.CommentStorage, notifier storage.Notifier) *Resolver {
	return &Resolver{
		PostStorage:    postStorage,
		CommentStorage: commentStorage,
		Notifier:       notifier,
	}
}
