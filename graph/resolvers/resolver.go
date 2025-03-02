package resolvers

import (
	"gql-comments/graph/generated"
	"gql-comments/storage/inmemory"
)

const (
	ConstOffset = 0
	ConstLimit  = 10
)

type mutationResolver struct{ *Resolver }

type queryResolver struct{ *Resolver }

type Resolver struct {
	StoragePost     *inmemory.InMemoryStoragePost
	StorageComments *inmemory.InMemoryStorageCommenst
	Notifier        *inmemory.Notifier
}

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

func NewResolver(storagePost *inmemory.InMemoryStoragePost, storageComments *inmemory.InMemoryStorageCommenst, notifier *inmemory.Notifier) *Resolver {
	return &Resolver{
		StoragePost:     storagePost,
		StorageComments: storageComments,
		Notifier:        notifier,
	}
}
