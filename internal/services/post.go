package services

import (
	"github.com/hashicorp/go-uuid"
	"time"
	"weekend.side/SocialMedia/internal/daos"
	"weekend.side/SocialMedia/internal/infra/db"
)

func CreatePost(req daos.Post) (*daos.Post, *daos.Error) {

	// verify email in db not exist

	accountId, _ := uuid.GenerateUUID()
	req.Id = accountId
	req.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.DbDriver.Exec(`insert into posts values (?, ?, ?, ?, ?);`, req.Id, req.Caption, req.ImageUrl, req.Creator, req.CreatedAt)

	if err != nil {
		return nil, &daos.Error{err.Error()}
	}

	return &req, nil

}

func FetchPost() ([]daos.Post, error) {

	reqPostsMap := make(map[string]daos.Post)
	var respPosts []daos.Post

	rows, err := db.DbDriver.Query(`SELECT
	posts.*,
	c.id as comment_id, c.content as comment_content, c.user as comment_user, c.created_at as comment_date
FROM
	posts
	LEFT JOIN (
		SELECT
			comments.*
		FROM
			comments
		LIMIT 2) AS c ON posts.id = c.post_id;`)

	if err != nil {
		return nil, err
	}
	for i := 0; rows.Next(); i++ {
		var reqPost daos.Post
		commentFetch := daos.CommmentFetch{}
		err := rows.Scan(&reqPost.Id, &reqPost.Caption, &reqPost.ImageUrl, &reqPost.Creator, &reqPost.CreatedAt, &commentFetch.CommentId, &commentFetch.CommentContent, &commentFetch.CommentUser, &commentFetch.CommentDate)
		if err != nil {
			return nil, err
		}

		if commentFetch.CommentId != nil {
			reqPost.Comments = append(reqPost.Comments, commentFetch)
		}
		if _, ok := reqPostsMap[reqPost.Id]; !ok {
			reqPostsMap[reqPost.Id] = reqPost
		} else {
			prevPost := reqPostsMap[reqPost.Id]
			prevPost.Comments = append(prevPost.Comments, reqPost.Comments...)
			reqPostsMap[reqPost.Id] = prevPost
		}
	}

	for _, post := range reqPostsMap {
		respPosts = append(respPosts, post)
	}

	return respPosts, nil
}

func CommentOnPost(comment *daos.Commment) (resp *daos.Commment, respErr *daos.Error) {
	accountId, _ := uuid.GenerateUUID()
	comment.Id = accountId
	comment.CreatedAt = time.Now().Format("2006-01-02 15:04:05")

	_, err := db.DbDriver.Exec(`insert into comments values (?, ?, ?, ?, ?);`, comment.Id, comment.Content, comment.PostId, comment.User, comment.CreatedAt)

	if err != nil {
		return nil, &daos.Error{err.Error()}
	}

	return comment, nil
}
