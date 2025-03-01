package resolvers

import (
	"context"
	"gql-comments/graph/generated"
	"gql-comments/structures"
)

type subscriptionResolver struct{ *Resolver }

func (r *Resolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID int) (<-chan *structures.Comment, error) {

	commentChannel := make(chan *structures.Comment, 1)
	r.Notifier.Subscribe(postID, commentChannel)
	// Отписываемся при завершении запроса
	go func() {
		<-ctx.Done()
		r.Notifier.Unsubscribe(postID, commentChannel)
		close(commentChannel)
	}()

	return commentChannel, nil
}
