package postgres

import (
	"context"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/rom1277/gql-comments/structures"
	"time"
)

func (s *PostgresCommentStorage) CreateComment(ctx context.Context, comment *structures.Comment) (*structures.Comment, error) {
	comment.CreatedAt = time.Now()
	var commentCount int
	countQuery := `
        SELECT COUNT(*) FROM comments WHERE post_id = $1
    `
	err := s.db.QueryRowContext(ctx, countQuery, comment.PostID).Scan(&commentCount)
	if err != nil {
		return nil, fmt.Errorf("failed to count comments: %w", err)
	}
	comment.ID = comment.PostID*100 + commentCount + 1

	query := `
        INSERT INTO comments (id, post_id, parent_id, author, text, created_at)
        VALUES ($1, $2, $3, $4, $5, $6)
    `
	_, err = s.db.ExecContext(ctx, query, comment.ID, comment.PostID, comment.ParentID, comment.User, comment.Text, comment.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("failed to create comment: %w", err)
	}
	if comment.PostID != 0 {
		updateQuery := `
            INSERT INTO post_comments (post_id, comment_id)
            VALUES ($1, $2)
            ON CONFLICT DO NOTHING
        `
		_, err := s.db.ExecContext(ctx, updateQuery, comment.PostID, comment.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to link comment to post: %w", err)
		}
	}

	if comment.ParentID != nil {
		parentID := *comment.ParentID
		updateQuery := `
            INSERT INTO comment_replies (parent_id, child_id)
            VALUES ($1, $2)
            ON CONFLICT DO NOTHING
        `
		_, err := s.db.ExecContext(ctx, updateQuery, parentID, comment.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to link comment to parent: %w", err)
		}
	} else {
		updateQuery := `
            INSERT INTO post_comments (post_id, comment_id)
            VALUES ($1, $2)
            ON CONFLICT DO NOTHING
        `
		_, err := s.db.ExecContext(ctx, updateQuery, comment.PostID, comment.ID)
		if err != nil {
			return nil, fmt.Errorf("failed to link root comment to post: %w", err)
		}
	}

	return comment, nil
}

func (s *PostgresCommentStorage) GetCommentsByPost(postID, limit, offset int) ([]*structures.Comment, error) {
	query := `
        SELECT id, post_id, parent_id, author, text, created_at
        FROM comments
        WHERE post_id = $1 AND parent_id IS NULL
        ORDER BY created_at ASC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.Query(query, postID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch comments: %w", err)
	}
	defer rows.Close()

	var comments []*structures.Comment
	for rows.Next() {
		var comment structures.Comment
		if err := rows.Scan(&comment.ID, &comment.PostID, &comment.ParentID, &comment.User, &comment.Text, &comment.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan comment: %w", err)
		}
		comments = append(comments, &comment)
	}
	return comments, nil
}

func (s *PostgresCommentStorage) GetResponsesByCommentID(commentID, limit, offset int) ([]*structures.Comment, error) {
	query := `
        SELECT id, post_id, parent_id, author, text, created_at
        FROM comments
        WHERE parent_id = $1
        ORDER BY created_at ASC
        LIMIT $2 OFFSET $3
    `
	rows, err := s.db.Query(query, commentID, limit, offset)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch replies: %w", err)
	}
	defer rows.Close()

	var responses []*structures.Comment
	for rows.Next() {
		var response structures.Comment
		if err := rows.Scan(&response.ID, &response.PostID, &response.ParentID, &response.User, &response.Text, &response.CreatedAt); err != nil {
			return nil, fmt.Errorf("failed to scan response: %w", err)
		}
		responses = append(responses, &response)
	}
	return responses, nil
}
