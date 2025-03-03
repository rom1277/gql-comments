package resolvers

import (
	"context"
	"errors"
	"github.com/rom1277/gql-comments/graph/model"
	"github.com/rom1277/gql-comments/structures"
)

func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*structures.Post, error) {
	post := &structures.Post{
		User:          input.User,
		Title:         input.Title,
		Content:       input.Content,
		AllowComments: input.AllowComments,
	}
	createdPost, err := r.PostStorage.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*structures.Post, error) {
	posts := r.PostStorage.GetAllPosts()
	for _, post := range posts {
		comments, err := r.CommentStorage.GetCommentsByPost(post.ID, ConstLimit, ConstOffset)
		if err != nil {
			return nil, err
		}
		post.Comments = comments
	}
	return posts, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*structures.Post, error) {
	post, err := r.PostStorage.GetPostByID(ctx, id)
	if err != nil {
		return nil, errors.New("failed to fetch post")
	}
	comments, err := r.CommentStorage.GetCommentsByPost(id, ConstLimit, ConstOffset)
	if err != nil {
		return nil, errors.New("failed to fetch post")
	}
	post.Comments = comments

	return post, nil
}

func (r *mutationResolver) CloseCommentsPost(ctx context.Context, user string, postID int, commentsAllowed bool) (*structures.Post, error) {
	post, err := r.PostStorage.GetPostByID(ctx, postID)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if post.User != user {
		return nil, errors.New("only the author can modify this post")
	}

	post.AllowComments = commentsAllowed
	err = r.PostStorage.CloseComments(ctx, post)
	if err != nil {
		return nil, errors.New("failed to update post")
	}

	return post, nil
}
