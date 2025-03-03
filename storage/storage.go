package storage

import (
	"context"
	"gql-comments/structures"
)

type PostStorage interface {
	CreatePost(ctx context.Context, post *structures.Post) (*structures.Post, error)
	GetAllPosts() []*structures.Post
	GetPostByID(ctx context.Context, id int) (*structures.Post, error)
	CloseComments(ctx context.Context, post *structures.Post) error
}
type CommentStorage interface {
	CreateComment(ctx context.Context, comment *structures.Comment) (*structures.Comment, error)
	GetCommentsByPost(postID, limit, offset int) ([]*structures.Comment, error)
	GetResponsesByCommentID(commentID, limit, offset int) ([]*structures.Comment, error)
}
type Notifier interface {
	Subscribe(postID int, ch chan *structures.Comment)
	Unsubscribe(postID int, ch chan *structures.Comment)
	Notify(postID int, comment *structures.Comment)
}
