package inmemory

import (
	"github.com/rom1277/gql-comments/structures"
)

func (n *Notifier) Subscribe(postID int, ch chan *structures.Comment) {
	n.mu.Lock()
	defer n.mu.Unlock()

	n.subscribers[postID] = append(n.subscribers[postID], ch)
}

func (n *Notifier) Unsubscribe(postID int, ch chan *structures.Comment) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if channels, ok := n.subscribers[postID]; ok {
		for i, channel := range channels {
			if channel == ch {
				n.subscribers[postID] = append(channels[:i], channels[i+1:]...)
				break
			}
		}
	}
}

func (n *Notifier) Notify(postID int, comment *structures.Comment) {
	n.mu.Lock()
	defer n.mu.Unlock()

	if channels, ok := n.subscribers[postID]; ok {
		for _, ch := range channels {
			go func(ch chan *structures.Comment) {
				ch <- comment
			}(ch)
		}
	}
}
