package mocks

import (
	"context"
	// "errors"
	// "github.com/stretchr/testify/assert"
	"github.com/rom1277/gql-comments/structures"
	"github.com/stretchr/testify/mock"
)

type MockPostStorage struct {
	mock.Mock
}

type MockCommentStorage struct {
	mock.Mock
}

type MockNotifier struct {
	mock.Mock
}

func (m *MockPostStorage) CreatePost(ctx context.Context, post *structures.Post) (*structures.Post, error) {
	args := m.Called(ctx, post)
	return args.Get(0).(*structures.Post), args.Error(1)
}

func (m *MockPostStorage) GetAllPosts() []*structures.Post {
	args := m.Called()
	return args.Get(0).([]*structures.Post)
}

func (m *MockPostStorage) GetPostByID(ctx context.Context, id int) (*structures.Post, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*structures.Post), args.Error(1)
}

func (m *MockPostStorage) CloseComments(ctx context.Context, post *structures.Post) error {
	args := m.Called(ctx, post)
	return args.Error(0)
}

func (m *MockCommentStorage) CreateComment(ctx context.Context, comment *structures.Comment) (*structures.Comment, error) {
	args := m.Called(ctx, comment)
	return args.Get(0).(*structures.Comment), args.Error(1)
}

func (m *MockCommentStorage) GetCommentsByPost(postID, limit, offset int) ([]*structures.Comment, error) {
	args := m.Called(postID, limit, offset)
	return args.Get(0).([]*structures.Comment), args.Error(1)
}

func (m *MockCommentStorage) GetResponsesByCommentID(commentID, limit, offset int) ([]*structures.Comment, error) {
	args := m.Called(commentID, limit, offset)
	return args.Get(0).([]*structures.Comment), args.Error(1)
}

func (m *MockNotifier) Subscribe(postID int, ch chan *structures.Comment) {
	m.Called(postID, ch)
}

func (m *MockNotifier) Unsubscribe(postID int, ch chan *structures.Comment) {
	m.Called(postID, ch)
}

func (m *MockNotifier) Notify(postID int, comment *structures.Comment) {
	m.Called(postID, comment)
}
