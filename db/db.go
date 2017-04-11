package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var database string
var username string
var password string

var Db *sql.DB

func init() {
	database = os.Getenv( "ENM_DATABASE")
	username = os.Getenv("ENM_DATABASE_USERNAME")
	password = os.Getenv("ENM_DATABASE_PASSWORD")

	if database == "" {
		panic("db: ENM_DATABASE not set")

	}

	if username == "" {
		panic("db: ENM_DATABASE_USERNAME not set")

	}

	if password == "" {
		panic("db: ENM_DATABASE_PASSWORD not set")
	}

	Db, err := sql.Open("mysql", username + ":" + password + "@/" + database)
	if err != nil {
		panic(err.Error())
	}

	if err = Db.Ping(); err == nil {
		fmt.Println( "Database " + database + " responding to ping")
	} else {
		panic(err.Error())
	}
}
