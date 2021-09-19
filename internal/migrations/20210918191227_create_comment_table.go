package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20210918191227, Down20210918191227)
}

func Up20210918191227(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table comments(
    id varchar(255),
    content varchar(255),
    post_id varchar(255) references posts(id),
    creator varchar(255),
    created_at datetime);`)
	return err
}

func Down20210918191227(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table comments;`)
	return err
}
