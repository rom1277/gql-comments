package models

import (
	"time"
)

// ИИ
type Post struct {
	ID            string    `json:"id"`
	Title         string    `json:"title"`
	Content       string    `json:"content"`
	AuthorID      string    `json:"author_id"`
	AllowComments bool      `json:"allow_comments"`
	CreatedAt     time.Time `json:"created_at"`
}

// type Post struct {
// 	ID              int       `json:"id" db:"id"`
// 	CreatedAt       time.Time `json:"createdAt" db:"created_at"`
// 	Name            string    `json:"name" db:"name"`
// 	Author          string    `json:"author" db:"author"`
// 	Content         string    `json:"content" db:"content"`
// 	CommentsAllowed bool      `json:"commentsAllowed" db:"comments_allowed"`
// 	//Comments        []*Comment `json:"comments,omitempty"`
// }
