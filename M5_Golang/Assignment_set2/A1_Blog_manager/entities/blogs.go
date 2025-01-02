package entities

type BlogPost struct {
	PostID    int    `json:"post_id"`
	Heading   string `json:"heading"`
	Body      string `json:"body"`
	Writer    string `json:"writer"`
	CreatedAt string `json:"created_at"`
}
