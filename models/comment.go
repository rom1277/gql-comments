package models

import (
	"time"
)

// ИИ
type Comment struct {
	ID        string    `json:"id"`
	PostID    string    `json:"post_id"`
	ParentID  *string   `json:"parent_id"`
	AuthorID  string    `json:"author_id"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}

// type Comment struct {
// 	ID        int       `json:"id" db:"id"`
// 	CreatedAt time.Time `json:"createdAt" db:"created_at"`
// 	Author    string    `json:"author" db:"author"`
// 	Content   string    `json:"content" db:"content"`
// 	Post      int       `json:"post" db:"post"`
// 	// Replies   []*Comment `json:"replies,omitempty"`
// 	ReplyTo *int `json:"replyTo,omitempty" db:"reply_to"`
// }
