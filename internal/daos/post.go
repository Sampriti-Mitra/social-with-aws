package daos

type Post struct {
	Id        string          `json:"id" sql:"id"`
	Caption   string          `json:"caption"`
	ImageUrl  string          `json:"image_url"`
	Creator   string          `json:"creator"`
	CreatedAt string          `json:"created_at"`
	Comments  []CommmentFetch `json:"comments"`
}

type Commment struct {
	Id        string `json:"id"`
	PostId    string `json:"post_id"`
	Content   string `json:"content"`
	User      string `json:"user"`
	CreatedAt string `json:"created_at"`
}

type CommmentFetch struct {
	CommentId      *string `json:"comment_id"`
	CommentContent *string `json:"comment_content"`
	CommentUser    *string `json:"comment_user"`
	CommentDate    *string `json:"comment_created_at"`
	PostId         *string `json:"post_id"`
}
