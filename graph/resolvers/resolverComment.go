package resolvers

import (
	"context"
	"errors"
	"github.com/rom1277/gql-comments/graph/model"
	"github.com/rom1277/gql-comments/structures"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*structures.Comment, error) {
	if input.PostID == 0 || input.User == "" || input.Text == "" {
		return nil, errors.New("invalid input: postID, user, and text must not be empty")
	}
	post, err := r.PostStorage.GetPostByID(ctx, input.PostID)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if !post.AllowComments {
		return nil, errors.New("comments are disabled for this post")
	}
	if len(input.Text) > 2000 {
		return nil, errors.New("comment exceeds 2000 characters")
	}
	comment := &structures.Comment{
		PostID:   input.PostID,
		ParentID: input.ParentID,
		User:     input.User,
		Text:     input.Text,
	}
	createdComment, err := r.CommentStorage.CreateComment(ctx, comment)
	if err != nil {
		return nil, err
	}
	r.Notifier.Notify(input.PostID, createdComment)
	return createdComment, nil
}

func (r *queryResolver) Comments(ctx context.Context, postID int, limit *int, offset *int) ([]*structures.Comment, error) {
	limitVal := ConstLimit
	if limit != nil {
		limitVal = *limit
	}

	offsetVal := ConstOffset
	if offset != nil {
		offsetVal = *offset
	}

	topLevelComments, err := r.CommentStorage.GetCommentsByPost(postID, limitVal, offsetVal)
	if err != nil {
		return nil, err
	}

	for _, comment := range topLevelComments {
		comment.Replies, err = r.getReplies(comment.ID)
		if err != nil {
			return nil, err
		}
	}

	return topLevelComments, nil
}

func (r *queryResolver) getReplies(commentID int) ([]*structures.Comment, error) {
	replies, err := r.CommentStorage.GetResponsesByCommentID(commentID, -1, 0)
	if err != nil {
		return nil, err
	}

	for _, reply := range replies {
		reply.Replies, err = r.getReplies(reply.ID)
		if err != nil {
			return nil, err
		}
	}

	return replies, nil
}

func (r *queryResolver) Replies(ctx context.Context, commentID int, limit *int, offset *int) ([]*structures.Comment, error) {
	limitVal := ConstLimit
	if limit != nil {
		limitVal = *limit
	}
	offsetVal := ConstOffset
	if offset != nil {
		offsetVal = *offset
	}
	replies, err := r.CommentStorage.GetResponsesByCommentID(commentID, limitVal, offsetVal)
	if err != nil {
		return nil, err
	}

	return replies, nil
}
