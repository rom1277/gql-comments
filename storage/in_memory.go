package storage

import "sync"

type Post struct {
	ID            string
	Title         string
	Content       string
	AllowComments bool
}

type InMemoryStorage struct {
	posts map[string]Post
	mu    sync.Mutex
}

func NewInMemoryStorage() *InMemoryStorage {
	return &InMemoryStorage{
		posts: make(map[string]Post),
	}
}

func (s *InMemoryStorage) CreatePost(post Post) {
	s.mu.Lock()
	defer s.mu.Unlock()
	s.posts[post.ID] = post
}

func (s *InMemoryStorage) GetAllPosts() []Post {
	s.mu.Lock()
	defer s.mu.Unlock()

	var posts []Post
	for _, post := range s.posts {
		posts = append(posts, post)
	}
	return posts
}
