package model

type Mutation struct {
}

type NewComment struct {
	PostID   int    `json:"postID"`
	ParentID *int   `json:"parentID,omitempty"`
	User     string `json:"user"`
	Text     string `json:"text"`
}

type NewPost struct {
	User          string `json:"user"`
	Title         string `json:"title"`
	Content       string `json:"content"`
	AllowComments bool   `json:"allowComments"`
}

type Query struct {
}

type Subscription struct {
}
