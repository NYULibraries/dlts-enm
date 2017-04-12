package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var username string
var password string

var DB *sql.DB
var Database string

func init() {
	Database = os.Getenv( "ENM_DATABASE")
	username = os.Getenv("ENM_DATABASE_USERNAME")
	password = os.Getenv("ENM_DATABASE_PASSWORD")

	if Database == "" {
		panic("db: ENM_DATABASE not set")

	}

	if username == "" {
		panic("db: ENM_DATABASE_USERNAME not set")

	}

	if password == "" {
		panic("db: ENM_DATABASE_PASSWORD not set")
	}

	var err error
	DB, err = sql.Open("mysql", username + ":" + password + "@/" + Database)
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err == nil {
		fmt.Println( "Database " + Database + " responding to ping")
	} else {
		panic(err.Error())
	}
}

func ClearTables() {
	tx, err := DB.Begin()
	if err != nil {
		panic( "db.ClearTables: " + err.Error())
	}

	tables := []string{
		"topics",
		"scopes",
		"relation_type",
		"relation_direction",
		"relations",
		"indexpatterns",
		"epubs",
		"locations",
		"names",
		"occurrences",
		"relations",
	}

	for _, table := range tables {
		_, err = tx.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			panic("db.ClearTables: " + err.Error())
		}
	}

	tx.Commit()
}
