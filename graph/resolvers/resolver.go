package resolvers

import (
	"context"
	"errors"
	"fmt"
	"gql-comments/graph/generated"
	"gql-comments/graph/model"
	"gql-comments/storage/inmemory"
	"gql-comments/structures"
)

const (
	ConstOffset = 0
	ConstLimit  = 10
)

type Resolver struct {
	Storage         *inmemory.InMemoryStorage
	StorageComments *inmemory.InMemoryStorageCommenst
}

func NewResolver(storage *inmemory.InMemoryStorage, storageComments *inmemory.InMemoryStorageCommenst) *Resolver {
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

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver {
	return &queryResolver{r}
}

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

// Комментарии
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
	// r.Notifier.Notify(comment.PostID, createdComment)
	return createdComment, nil
}

// Comments is the resolver for the comments field.
func (r *queryResolver) Comments(ctx context.Context, postID int, limit *int, offset *int) ([]*structures.Comment, error) {
	// Устанавливаем значения по умолчанию для limit и offset
	limitVal := ConstLimit // Значение по умолчанию для limit
	if limit != nil {
		limitVal = *limit
	}

	offsetVal := ConstOffset // Значение по умолчанию для offset
	if offset != nil {
		offsetVal = *offset
	}

	// Вызываем метод хранилища для получения комментариев
	comments, err := r.StorageComments.GetCommentsByPost(postID, limitVal, offsetVal)
	if err != nil {
		return nil, err
	}

	return comments, nil
}

// Replies is the resolver for the replies field.
func (r *queryResolver) Replies(ctx context.Context, commentID int, limit *int, offset *int) ([]*structures.Comment, error) {
	panic(fmt.Errorf("not implemented: Replies - replies"))
}

// CloseCommentsPost is the resolver for the closeCommentsPost field.
func (r *mutationResolver) CloseCommentsPost(ctx context.Context, user string, postID int, commentsAllowed bool) (*structures.Post, error) {
	panic(fmt.Errorf("not implemented: CloseCommentsPost - closeCommentsPost"))
}
