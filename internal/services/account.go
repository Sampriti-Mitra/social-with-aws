package services

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/hashicorp/go-uuid"
	"weekend.side/SocialMedia/internal/daos"
	"weekend.side/SocialMedia/internal/infra/db"
)

func CreateAccount(req daos.Account) (*daos.Account, *daos.Error) {

	// verify email in db not exist
	var id string
	err := db.DbDriver.QueryRow(`select id from accounts where username = ?`, req.Username).Scan(&id)
	if err != sql.ErrNoRows {
		return nil, &daos.Error{errors.New("Username is taken").Error()}
	}

	accountId, _ := uuid.GenerateUUID()
	req.Id = accountId

	res, err := db.DbDriver.Exec(`insert into accounts values (?, ?);`, req.Id, req.Username)

	if err != nil {
		return nil, &daos.Error{err.Error()}
	}
	// save email in db
	fmt.Print(res)

	return &req, nil

}

func DeleteAccount(username string) (resp interface{}, error2 *daos.Error) {

	ctx := context.Background()

	// begin db transaction
	tx, err := db.DbDriver.BeginTx(ctx, nil)
	if err != nil {
		return nil, &daos.Error{err}
	}

	// get all posts of the account
	rows, err := tx.Query(`select id from posts where creator = ?`, username)
	if err != nil {
		tx.Rollback()
		return nil, &daos.Error{err}
	}

	var postIds []string

	for i := 0; rows.Next(); i++ {
		var postId string
		err := rows.Scan(&postId)
		if err != nil {
			return nil, &daos.Error{err}
		}
		postIds = append(postIds, postId)
	}

	// delete the posts of the user account
	_, transactionErr := tx.ExecContext(ctx, `delete from posts where creator = ?`, username)
	if transactionErr != nil {
		tx.Rollback()
		return nil, &daos.Error{transactionErr}
	}

	// delete all the comments of usrr account
	_, transactionErr = tx.ExecContext(ctx, `delete from comments where commented_user = ?`, username)
	if transactionErr != nil {
		tx.Rollback()
		return nil, &daos.Error{transactionErr}
	}

	if len(postIds) != 0 {
		queryString := "delete from comments where post_id in ("

		for i := 0; i < len(postIds); i++ {
			queryString = queryString + "'" + postIds[i] + "',"
		}

		queryString = queryString[:len(queryString)-1]

		queryString += ");"

		// delete all the comments in the posts of account user
		_, transactionErr = tx.ExecContext(ctx, queryString)
		if transactionErr != nil {
			tx.Rollback()
			return nil, &daos.Error{transactionErr}
		}
	}

	// delete user account holder
	_, err = tx.ExecContext(ctx, `delete from accounts where username = ?;`, username)

	if transactionErr != nil {
		tx.Rollback()
		return nil, &daos.Error{transactionErr}
	}

	// commit to db
	commitErr := tx.Commit()
	if commitErr != nil {
		tx.Rollback()
		return nil, &daos.Error{transactionErr}
	}

	return map[string]interface{}{
		"success": true,
	}, nil
}
