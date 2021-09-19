package services

import (
	"database/sql"
	"errors"
	"net/http"
	"weekend.side/SocialMedia/internal/daos"
	"weekend.side/SocialMedia/internal/infra/db"
)

func AuthenticateRequest(r *http.Request) (string, *daos.Error) {
	accountID := r.Header.Get("X-Account-Id")
	var id string
	err := db.DbDriver.QueryRow(`select username from accounts where id = ?`, accountID).Scan(&id)
	if err == sql.ErrNoRows {
		return accountID, &daos.Error{errors.New("No such account").Error()}
	}
	return accountID, nil
}
