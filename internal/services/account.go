package services

import (
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
