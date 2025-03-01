package inmemory

import (
	"gql-comments/structures"
	"sync"
)

type InMemoryStoragePost struct {
	posts map[int]structures.Post
	mu    sync.Mutex
}

type InMemoryStorageCommenst struct {
	comments     map[int]structures.Comment
	postComments map[int][]int
	replies      map[int][]int
	mu           sync.Mutex
}

type Notifier struct {
	subscribers map[int][]chan *structures.Comment
	mu          sync.Mutex
}

func NewInMemoryStoragePost() *InMemoryStoragePost {
	return &InMemoryStoragePost{
		posts: make(map[int]structures.Post),
	}
}

func NewInMemoryStorageCommenst() *InMemoryStorageCommenst {
	return &InMemoryStorageCommenst{
		comments:     make(map[int]structures.Comment),
		postComments: make(map[int][]int),
		replies:      make(map[int][]int),
	}
}

func NewNotifier() *Notifier {
	return &Notifier{
		subscribers: make(map[int][]chan *structures.Comment),
	}
}
