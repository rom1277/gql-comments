package resolvers

import (
	"context"
	"errors"
	"gql-comments/graph/model"
	"gql-comments/structures"
)

// Посты:
func (r *mutationResolver) CreatePost(ctx context.Context, input *model.NewPost) (*structures.Post, error) {
	if input.Title == "" || input.Content == "" {
		return nil, errors.New("title and content must not be empty")
	}
	post := &structures.Post{
		User:          input.User,
		Title:         input.Title,
		Content:       input.Content,
		AllowComments: input.AllowComments,
	}
	createdPost, err := r.Storage.CreatePost(ctx, post)
	if err != nil {
		return nil, err
	}
	return createdPost, nil
}

func (r *queryResolver) Posts(ctx context.Context) ([]*structures.Post, error) {
	posts := r.Storage.GetAllPosts()
	var result []*structures.Post
	for i := range posts {
		result = append(result, &posts[i])
	}
	if len(result) == 0 {
		return nil, errors.New("no added posts")
	}
	return result, nil
}

func (r *queryResolver) Post(ctx context.Context, id int) (*structures.Post, error) {
	post, err := r.Storage.GetPostbyId(ctx, id)
	if err != nil {
		return post, err
	}
	return post, nil
}

func (r *mutationResolver) CloseCommentsPost(ctx context.Context, user string, postID int, commentsAllowed bool) (*structures.Post, error) {
	post, err := r.Storage.GetPostbyId(ctx, postID)
	if err != nil {
		return nil, errors.New("post not found")
	}
	if post.User != user {
		return nil, errors.New("only the author can modify this post")
	}
	post.AllowComments = commentsAllowed

	err = r.Storage.CloseComments(ctx, post)
	if err != nil {
		return nil, errors.New("failed to update post")
	}

	return post, nil
}
