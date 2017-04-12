// Copyright © 2017 NYU
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"database/sql"
	"fmt"
	"strings"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var username string
var password string

var DB *sql.DB
var Database string

// Tables in order that they would need to be in when deleting all data table-by-table.
var tables = []string{
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

var insertStmts map[string]*sql.Stmt

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

	insertStmts = make(map[string]*sql.Stmt)
	prepareInsertStmts()
}

func ClearTables() {
	tx, err := DB.Begin()
	if err != nil {
		panic( "db.ClearTables: " + err.Error())
	}

	for _, table := range tables {
		_, err = tx.Exec(fmt.Sprintf("DELETE FROM %s", table))
		if err != nil {
			panic("db.ClearTables: " + err.Error())
		}
	}

	tx.Commit()
}

func getColumnNames(table string) (columns []string) {
	rows, err := DB.Query( "SELECT * FROM " + table)
	if err != nil {
		panic("db.GetColumnNames: " + err.Error())
	}

	columns, err = rows.Columns()
	if err != nil {
		panic("db.GetColumnNames: " + err.Error())
	}

	return
}

func prepareInsertStmts() {
	var insertSql, cols, vals string
	for _, table := range tables {
		var columns = getColumnNames(table)
		var numColumns = len(columns)
		cols = strings.Join(columns, ", ")

		placeholders := []string{}
		for j := 0; j < numColumns; j++ {
			placeholders = append(placeholders, "?")
		}

		vals = strings.Join(placeholders, ", ")
		insertSql = "INSERT INTO " + table + " (" + cols + ")" +
			" VALUES (" + vals + ")"

		var err error
		insertStmts[table], err = DB.Prepare(insertSql)
		if err != nil {
			panic("db.prepareInsertStatements: " + err.Error())
		}

		fmt.Printf("Prepared: \"%s\"\n", insertSql)
	}
}
