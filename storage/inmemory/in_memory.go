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

func (s *InMemoryStorage) GetPostbyId(ctx context.Context, id int) (*structures.Post, error) {
	s.mu.Lock()
	defer s.mu.Unlock()
	post, ok := s.posts[id]
	if !ok {
		return nil, errors.New("there is no such id")
	}
	return &post, nil
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

func (c *InMemoryStorageCommenst) CreateComment(ctx context.Context, comment *structures.Comment) (*structures.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()
	comment.CreatedAt = time.Now()
	comment.ID = comment.PostID*100 + len(c.comments) + 1
	c.comments[comment.ID] = *comment
	if comment.PostID != 0 {
		if _, ok := c.postComments[comment.PostID]; !ok {
			c.postComments[comment.PostID] = []int{}
		}
		c.postComments[comment.PostID] = append(c.postComments[comment.PostID], comment.ID)
	}

	// Если комментарий является ответом на другой комментарий, добавляем его в replies
	if comment.ParentID != nil {
		parentID := *comment.ParentID
		if _, ok := c.replies[parentID]; !ok {
			c.replies[parentID] = []int{}
		}
		c.replies[parentID] = append(c.replies[parentID], comment.ID)
	} else {
		if _, ok := c.replies[comment.PostID]; !ok {
			c.replies[comment.PostID] = []int{}
		}
		c.replies[comment.PostID] = append(c.replies[comment.PostID], comment.ID)
	}

	return comment, nil
}

func (c *InMemoryStorageCommenst) GetCommentsByPost(postID, limit, offset int) ([]*structures.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	var result []*structures.Comment
	for _, comment := range c.comments {
		if comment.PostID == postID && comment.ParentID == nil {
			com := comment
			result = append(result, &com)
		}
	}

	if offset > len(result) {
		return nil, nil
	}
	if offset+limit > len(result) || limit == -1 {
		return result[offset:], nil
	}
	if offset < 0 || limit < 0 {
		return nil, errors.New("limit and offset should not be negative")
	}

	return result[offset : offset+limit], nil
}

func (c *InMemoryStorageCommenst) GetResponsesByCommentID(commentID, limit, offset int) ([]*structures.Comment, error) {
	c.mu.Lock()
	defer c.mu.Unlock()

	replies, ok := c.replies[commentID]
	if !ok {
		return nil, nil
	}

	var result []*structures.Comment
	for _, replyID := range replies {
		comment, ok := c.comments[replyID]
		if ok {
			com := comment
			result = append(result, &com)
		}
	}

	if offset > len(result) {
		return nil, nil
	}
	if offset+limit > len(result) || limit == -1 {
		return result[offset:], nil
	}
	if offset < 0 || limit < 0 {
		return nil, errors.New("limit and offset should not be negative")
	}
	return result[offset : offset+limit], nil
}

func (s *InMemoryStorage) UpdatePost(ctx context.Context, post *structures.Post) error {
	s.mu.Lock()
	defer s.mu.Unlock()
	if _, exists := s.posts[post.ID]; !exists {
		return errors.New("post not found")
	}
	s.posts[post.ID] = *post
	return nil
}
