package services

import (
	"fmt"
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

	respCommentMap := make(map[string][]daos.CommmentFetch)

	rows, err := db.DbDriver.Query(`SELECT
	posts.*
FROM
	posts
	LEFT JOIN comments AS c ON posts.id = c.post_id
GROUP BY
	posts.id,
	posts.caption,
	posts.image_url,
	posts.creator,
	posts.created_at
ORDER BY
	count(c.id)
	DESC;`)

	if err != nil {
		return nil, err
	}

	commentRows, err := db.DbDriver.Query(`SELECT * from comments order by created_at desc`)
	if err != nil {
		return nil, err
	}

	// for every post, group the comments by post id

	for i := 0; commentRows.Next(); i++ {
		comment := daos.CommmentFetch{}

		err := commentRows.Scan(&comment.CommentId, &comment.CommentContent, &comment.PostId, &comment.CommentUser, &comment.CommentDate)
		if err != nil {
			return nil, err
		}

		if comment.CommentId != nil && len(respCommentMap[*comment.PostId]) < 2 {
			respCommentMap[*comment.PostId] = append(respCommentMap[*comment.PostId], daos.CommmentFetch{
				comment.CommentId,
				comment.CommentContent,
				comment.CommentUser,
				comment.CommentDate,
				comment.PostId,
			})
		}
	}

	for i := 0; rows.Next(); i++ {
		var reqPost daos.Post

		err := rows.Scan(&reqPost.Id, &reqPost.Caption, &reqPost.ImageUrl, &reqPost.Creator, &reqPost.CreatedAt)
		if err != nil {
			return nil, err
		}

		reqPost.Comments = respCommentMap[reqPost.Id]

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

	res, err := db.DbDriver.Exec(`insert into comments (id, content, post_id, commented_user, created_at) values (?, ?, ?, ?, ?);`, comment.Id, comment.Content, comment.PostId, comment.User, comment.CreatedAt)

	if err != nil {
		return nil, &daos.Error{err.Error()}
	}

	fmt.Print(res)

	return comment, nil
}
