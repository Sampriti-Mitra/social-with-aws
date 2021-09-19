package main

import (
	"database/sql"
	"github.com/pressly/goose"
)

func init() {
	goose.AddMigration(Up20210918191206, Down20210918191206)
}

func Up20210918191206(tx *sql.Tx) error {
	// This code is executed when the migration is applied.
	_, err := tx.Exec(`create table accounts(
    id varchar(255),
    username varchar(255));`)
	return err
}

func Down20210918191206(tx *sql.Tx) error {
	// This code is executed when the migration is rolled back.
	_, err := tx.Exec(`drop table accounts;`)
	return err
}
