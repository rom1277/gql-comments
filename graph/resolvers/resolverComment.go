package resolvers

import (
	"context"
	"errors"
	"gql-comments/graph/model"
	"gql-comments/structures"
)

func (r *mutationResolver) CreateComment(ctx context.Context, input model.NewComment) (*structures.Comment, error) {
	if input.PostID == 0 || input.User == "" || input.Text == "" {
		return nil, errors.New("invalid input: postID, user, and text must not be empty")
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
	createdComment, err := r.StorageComments.CreateComment(ctx, comment)
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

	topLevelComments, err := r.StorageComments.GetCommentsByPost(postID, limitVal, offsetVal)
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

// Вспомогательный метод для получения ответов на комментарии
func (r *queryResolver) getReplies(commentID int) ([]*structures.Comment, error) {
	replies, err := r.StorageComments.GetResponsesByCommentID(commentID, -1, 0)
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
	replies, err := r.StorageComments.GetResponsesByCommentID(commentID, limitVal, offsetVal)
	if err != nil {
		return nil, err
	}

	return replies, nil
}
