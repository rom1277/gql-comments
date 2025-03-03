package resolvers_test

import (
	"context"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gql-comments/graph/model"
	"gql-comments/graph/resolvers"
	"gql-comments/graph/resolvers/testResolvers/mocks"
	"gql-comments/structures"
)

func TestCreatePost(t *testing.T) {
	mockPostStorage := new(mocks.MockPostStorage)

	mockPostStorage.On("CreatePost", mock.Anything, mock.Anything).Return(&structures.Post{
		ID:            1,
		User:          "JohnDoe",
		Title:         "Test Post",
		Content:       "This is a test post",
		AllowComments: true,
	}, nil)

	resolver := &resolvers.Resolver{
		PostStorage: mockPostStorage,
	}

	input := &model.NewPost{
		User:          "JohnDoe",
		Title:         "Test Post",
		Content:       "This is a test post",
		AllowComments: true,
	}
	post, err := resolver.Mutation().CreatePost(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Post", post.Title)

	mockPostStorage.AssertCalled(t, "CreatePost", mock.Anything, mock.Anything)
}

func TestPosts(t *testing.T) {

	mockPostStorage := new(mocks.MockPostStorage)
	mockCommentStorage := new(mocks.MockCommentStorage)

	mockPostStorage.On("GetAllPosts").Return([]*structures.Post{
		{
			ID:            1,
			User:          "JohnDoe",
			Title:         "Test Post",
			Content:       "This is a test post",
			AllowComments: true,
		},
	})
	mockCommentStorage.On("GetCommentsByPost", 1, mock.Anything, mock.Anything).Return([]*structures.Comment{}, nil)

	resolver := &resolvers.Resolver{
		PostStorage:    mockPostStorage,
		CommentStorage: mockCommentStorage,
	}

	posts, err := resolver.Query().Posts(context.Background())

	assert.NoError(t, err)
	assert.Len(t, posts, 1)
	assert.Equal(t, "Test Post", posts[0].Title)

	mockPostStorage.AssertCalled(t, "GetAllPosts")
	mockCommentStorage.AssertCalled(t, "GetCommentsByPost", 1, mock.Anything, mock.Anything)
}

func TestPost(t *testing.T) {
	mockPostStorage := new(mocks.MockPostStorage)
	mockCommentStorage := new(mocks.MockCommentStorage)

	mockPostStorage.On("GetPostByID", mock.Anything, 1).Return(&structures.Post{
		ID:            1,
		User:          "JohnDoe",
		Title:         "Test Post",
		Content:       "This is a test post",
		AllowComments: true,
	}, nil)
	mockCommentStorage.On("GetCommentsByPost", 1, mock.Anything, mock.Anything).Return([]*structures.Comment{}, nil)

	resolver := &resolvers.Resolver{
		PostStorage:    mockPostStorage,
		CommentStorage: mockCommentStorage,
	}

	post, err := resolver.Query().Post(context.Background(), 1)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.Equal(t, "Test Post", post.Title)

	mockPostStorage.AssertCalled(t, "GetPostByID", mock.Anything, 1)
	mockCommentStorage.AssertCalled(t, "GetCommentsByPost", 1, mock.Anything, mock.Anything)
}

func TestCloseCommentsPost(t *testing.T) {
	mockPostStorage := new(mocks.MockPostStorage)
	mockPostStorage.On("GetPostByID", mock.Anything, 1).Return(&structures.Post{
		ID:            1,
		User:          "JohnDoe",
		Title:         "Test Post",
		Content:       "This is a test post",
		AllowComments: true,
	}, nil)
	mockPostStorage.On("CloseComments", mock.Anything, mock.Anything).Return(nil)

	resolver := &resolvers.Resolver{
		PostStorage: mockPostStorage,
	}

	post, err := resolver.Mutation().CloseCommentsPost(context.Background(), "JohnDoe", 1, false)

	assert.NoError(t, err)
	assert.NotNil(t, post)
	assert.False(t, post.AllowComments)

	mockPostStorage.AssertCalled(t, "GetPostByID", mock.Anything, 1)
	mockPostStorage.AssertCalled(t, "CloseComments", mock.Anything, mock.Anything)

	_, err = resolver.Mutation().CloseCommentsPost(context.Background(), "AnotherUser", 1, false)
	assert.Error(t, err)
	assert.EqualError(t, err, "only the author can modify this post")
}
