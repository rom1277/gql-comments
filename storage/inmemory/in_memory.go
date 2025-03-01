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

type InMemoryStorageCommenst struct {
	comments     map[int]structures.Comment
	postComments map[int][]int
	replies      map[int][]int
	mu           sync.Mutex
}

func NewInMemoryStorageCommenst() *InMemoryStorageCommenst {
	return &InMemoryStorageCommenst{
		comments:     make(map[int]structures.Comment),
		postComments: make(map[int][]int),
		replies:      make(map[int][]int),
	}
}

func (r *InMemoryStorageCommenst) CreateComment(ctx context.Context, comment *structures.Comment) (*structures.Comment, error) {
	r.mu.Lock()
	defer r.mu.Unlock()
	comment.CreatedAt = time.Now()
	comment.ID = len(r.comments) + 1
	r.comments[comment.ID] = *comment
	r.replies[comment.PostID] = append(r.replies[comment.PostID], comment.ID)
	// if comment.ParentID != nil {
	// 	r.repliers[*comment.ParentID] = append(r.repliers[*comment.ParentID], comment.ID)
	// }
	return comment, nil
}
