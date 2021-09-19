package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20210918191221, Down20210918191221)
}

func Up20210918191221(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table posts(
    id varchar(255),
    caption varchar(255),
    image_url varchar(255),
    creator varchar(255) references accounts(id),
    created_at datetime);`)
	return err
}

func Down20210918191221(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table posts;`)
	return err
}
