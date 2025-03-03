package resolvers

import (
	"context"
	"github.com/rom1277/gql-comments/graph/generated"
	"github.com/rom1277/gql-comments/structures"
)

type subscriptionResolver struct{ *Resolver }

func (r *Resolver) Subscription() generated.SubscriptionResolver {
	return &subscriptionResolver{r}
}

func (r *subscriptionResolver) CommentAdded(ctx context.Context, postID int) (<-chan *structures.Comment, error) {
	commentChannel := make(chan *structures.Comment, 1)
	r.Notifier.Subscribe(postID, commentChannel)
	go func() {
		<-ctx.Done()
		r.Notifier.Unsubscribe(postID, commentChannel)
		close(commentChannel)
	}()

	return commentChannel, nil
}
