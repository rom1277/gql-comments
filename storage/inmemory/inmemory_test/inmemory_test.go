package inmemory_test

import (
	"context"
	"gql-comments/storage/inmemory"
	"gql-comments/structures"
	"testing"
)

func TestCreateComment(t *testing.T) {
	storage := inmemory.NewInMemoryStorageCommenst()

	comment := &structures.Comment{
		PostID: 1,
		User:   "User1",
		Text:   "This is a comment",
	}

	createdComment, err := storage.CreateComment(context.Background(), comment)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if createdComment.ID != 101 {
		t.Errorf("expected ID to be 101, got %d", createdComment.ID)
	}
	if createdComment.User != "User1" {
		t.Errorf("expected User to be 'User1', got '%s'", createdComment.User)
	}
	if createdComment.Text != "This is a comment" {
		t.Errorf("expected Text to be 'This is a comment', got '%s'", createdComment.Text)
	}
	if createdComment.CreatedAt.IsZero() {
		t.Error("expected CreatedAt to be set")
	}

}

func TestGetCommentsByPost(t *testing.T) {
	storage := inmemory.NewInMemoryStorageCommenst()

	storage.CreateComment(context.Background(), &structures.Comment{
		PostID: 1,
		User:   "User1",
		Text:   "Comment 1",
	})
	storage.CreateComment(context.Background(), &structures.Comment{
		PostID: 1,
		User:   "User2",
		Text:   "Comment 2",
	})

	comments, err := storage.GetCommentsByPost(1, -1, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(comments) != 2 {
		t.Errorf("expected 2 comments, got %d", len(comments))
	}

	if comments[0].Text != "Comment 1" {
		t.Errorf("expected first comment text to be 'Comment 1', got '%s'", comments[0].Text)
	}
	if comments[1].Text != "Comment 2" {
		t.Errorf("expected second comment text to be 'Comment 2', got '%s'", comments[1].Text)
	}
}

func TestGetResponsesByCommentID(t *testing.T) {
	storage := inmemory.NewInMemoryStorageCommenst()

	parentComment, _ := storage.CreateComment(context.Background(), &structures.Comment{
		PostID: 1,
		User:   "ParentUser",
		Text:   "Parent Comment",
	})

	responses, err := storage.GetResponsesByCommentID(parentComment.ID, 0, 0)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if len(responses) != 0 {
		t.Errorf("expected 0 response, got %d", len(responses))
	}

	if len(responses) != 0 {
		t.Errorf("expected 0 response, got %d", len(responses))
		return
	}
}

func TestGetPostByID(t *testing.T) {
	storage := inmemory.NewInMemoryStoragePost()

	post := &structures.Post{
		Title:         "My First Post",
		User:          "JohnDoe",
		Content:       "This is the content",
		AllowComments: true,
	}
	createdPost, _ := storage.CreatePost(context.Background(), post)

	retrievedPost, err := storage.GetPostByID(context.Background(), createdPost.ID)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}

	if retrievedPost.Title != "My First Post" {
		t.Errorf("expected Title to be 'My First Post', got '%s'", retrievedPost.Title)
	}
}
