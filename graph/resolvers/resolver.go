package resolvers

import (
	"gql-comments/graph/generated"
	"gql-comments/storage/inmemory"
)

const (
	ConstOffset = 0
	ConstLimit  = 10
)

type Resolver struct {
	Storage         *inmemory.InMemoryStoragePost
	StorageComments *inmemory.InMemoryStorageCommenst
}

func NewResolver(storage *inmemory.InMemoryStoragePost, storageComments *inmemory.InMemoryStorageCommenst) *Resolver {
	return &Resolver{
		Storage:         storage,
		StorageComments: storageComments,
	}
}

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }

func (r *Resolver) Mutation() generated.MutationResolver {
	return &mutationResolver{r}
}

func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}
