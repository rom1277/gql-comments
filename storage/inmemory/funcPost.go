package inmemory

import (
	"context"
	"errors"
	"gql-comments/structures"
	"time"
)

func (sp *InMemoryStoragePost) CreatePost(ctx context.Context, post *structures.Post) (*structures.Post, error) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	post.CreatedAt = time.Now()
	post.ID = len(sp.posts) + 1
	sp.posts[post.ID] = *post
	return post, nil
}

func (sp *InMemoryStoragePost) GetAllPosts() []structures.Post {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	var posts []structures.Post
	for _, post := range sp.posts {
		posts = append(posts, post)
	}
	return posts
}

func (sp *InMemoryStoragePost) GetPostbyId(ctx context.Context, id int) (*structures.Post, error) {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	post, ok := sp.posts[id]
	if !ok {
		return nil, errors.New("there is no such id")
	}
	return &post, nil
}

func (sp *InMemoryStoragePost) CloseComments(ctx context.Context, post *structures.Post) error {
	sp.mu.Lock()
	defer sp.mu.Unlock()
	if _, exists := sp.posts[post.ID]; !exists {
		return errors.New("post not found")
	}
	sp.posts[post.ID] = *post
	return nil
}
