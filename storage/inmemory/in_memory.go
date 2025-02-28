package inmemory

import (
	"context"
	"errors"
	"gql-comments/structures"
	"sync"
	"time"
)

type InMemoryStorage struct {
	posts map[int]structures.Post
	mu    sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		posts: make(map[int]structures.Post),
	}
}

func (s *InMemoryStorage) CreatePost(ctx context.Context, post *structures.Post) (*structures.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	post.CreatedAt = time.Now()
	post.ID = len(s.posts) + 1
	s.posts[post.ID] = *post
	return post, nil
}

func (s *InMemoryStorage) GetAllPosts() []structures.Post {
	s.mu.Lock()
	defer s.mu.Unlock()
	var posts []structures.Post
	for _, post := range s.posts {
		posts = append(posts, post)
	}
	return posts
}

func (s *InMemoryStorage) GetPostbyId(ctx context.Context, id int) (structures.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	post, ok := s.posts[id]
	if !ok {
		return post, errors.New("there is no such id")
	}
	return post, nil
}
