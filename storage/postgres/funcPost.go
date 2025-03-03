package postgres

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rom1277/gql-comments/structures"
	"time"
)

func (s *PostgresPostStorage) CreatePost(ctx context.Context, post *structures.Post) (*structures.Post, error) {
	if post.CreatedAt.IsZero() {
		post.CreatedAt = time.Now()
	}

	query := `
        INSERT INTO posts (title, content, author, allow_comments, created_at)
        VALUES ($1, $2, $3, $4, $5)
        RETURNING id
    `
	err := s.db.QueryRow(
		query, post.Title, post.Content, post.User, post.AllowComments, post.CreatedAt,
	).Scan(&post.ID)
	if err != nil {
		return nil, fmt.Errorf("failed to create post: %w", err)
	}
	return post, nil
}

func (s *PostgresPostStorage) GetAllPosts() []*structures.Post {
	query := `
        SELECT id, author, title, content, allow_comments, created_at
        FROM posts
        ORDER BY created_at DESC
    `
	rows, err := s.db.Query(query)
	if err != nil {
		fmt.Println("Error fetching posts:", err)
		return nil
	}
	defer rows.Close()

	var posts []*structures.Post
	for rows.Next() {
		var post structures.Post
		if err := rows.Scan(&post.ID, &post.User, &post.Title, &post.Content, &post.AllowComments, &post.CreatedAt); err != nil {
			fmt.Println("Error scanning post:", err)
			continue
		}
		posts = append(posts, &post)
	}
	return posts
}

func (s *PostgresPostStorage) GetPostByID(ctx context.Context, id int) (*structures.Post, error) {
	query := `
        SELECT id, author, title, content, allow_comments, created_at
        FROM posts
        WHERE id = $1
    `
	var post structures.Post
	err := s.db.QueryRowContext(ctx, query, id).Scan(&post.ID, &post.User, &post.Title, &post.Content, &post.AllowComments, &post.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("post not found")
	} else if err != nil {
		return nil, fmt.Errorf("failed to fetch post: %w", err)
	}
	return &post, nil
}

func (s *PostgresPostStorage) CloseComments(ctx context.Context, post *structures.Post) error {
	query := `
        UPDATE posts
        SET allow_comments = $1
        WHERE id = $2
    `
	_, err := s.db.ExecContext(ctx, query, post.AllowComments, post.ID)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}
	return nil
}
