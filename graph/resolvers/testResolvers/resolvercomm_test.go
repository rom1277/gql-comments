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

func TestCreateComment(t *testing.T) {
	mockPostStorage := new(mocks.MockPostStorage)
	mockCommentStorage := new(mocks.MockCommentStorage)
	mockNotifier := new(mocks.MockNotifier)
	mockPostStorage.On("GetPostByID", mock.Anything, 1).Return(&structures.Post{
		ID:            1,
		User:          "JohnDoe",
		Title:         "Test Post",
		Content:       "This is a test post",
		AllowComments: true,
	}, nil)
	mockCommentStorage.On("CreateComment", mock.Anything, mock.Anything).Return(&structures.Comment{
		ID:     1,
		PostID: 1,
		User:   "User1",
		Text:   "This is a comment",
	}, nil)
	mockNotifier.On("Notify", 1, mock.Anything).Return()
	resolver := &resolvers.Resolver{
		PostStorage:    mockPostStorage,
		CommentStorage: mockCommentStorage,
		Notifier:       mockNotifier,
	}
	input := model.NewComment{
		PostID: 1,
		User:   "User1",
		Text:   "This is a comment",
	}
	comment, err := resolver.Mutation().CreateComment(context.Background(), input)

	assert.NoError(t, err)
	assert.NotNil(t, comment)
	assert.Equal(t, "This is a comment", comment.Text)

	mockPostStorage.AssertCalled(t, "GetPostByID", mock.Anything, 1)
	mockCommentStorage.AssertCalled(t, "CreateComment", mock.Anything, mock.Anything)
	mockNotifier.AssertCalled(t, "Notify", 1, mock.Anything)
}

func TestComments(t *testing.T) {
	mockCommentStorage := new(mocks.MockCommentStorage)
	mockCommentStorage.On("GetCommentsByPost", 1, mock.Anything, mock.Anything).Return([]*structures.Comment{
		{
			ID:     1,
			PostID: 1,
			User:   "User1",
			Text:   "Top-level comment",
		},
	}, nil)
	mockCommentStorage.On("GetResponsesByCommentID", 1, -1, 0).Return([]*structures.Comment{}, nil)
	resolver := &resolvers.Resolver{
		CommentStorage: mockCommentStorage,
	}

	topLevelComments, err := resolver.Query().Comments(context.Background(), 1, nil, nil)

	assert.NoError(t, err)
	assert.Len(t, topLevelComments, 1)
	assert.Equal(t, "Top-level comment", topLevelComments[0].Text)

	mockCommentStorage.AssertCalled(t, "GetCommentsByPost", 1, mock.Anything, mock.Anything)
	mockCommentStorage.AssertCalled(t, "GetResponsesByCommentID", 1, -1, 0)
}
