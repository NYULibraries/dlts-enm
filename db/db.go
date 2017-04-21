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
	"sort"
	"strings"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nyulibraries/dlts/enm/util"
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
}

func init() {
	Database = os.Getenv("ENM_DATABASE")
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

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}
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

func GenerateDbCode() {
	var result string = `/* DO NOT EDIT: this file was generated by db.GenerateDbCode() */

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

%s
func prepareInsertStmts() {
	var err error
%s
}
`
	var insertStmtDeclarations string
	var prepareCode string

	for _, table := range tables {
		insertStmtDeclarations += getInsertStmtDeclaration(table)
		prepareCode += getPrepareStmtCode(table)
	}

	fmt.Printf(result, insertStmtDeclarations, prepareCode)
}

func getInsertStmtDeclaration(table string) (string) {
	return fmt.Sprintf( "var insertStmt_%s *sql.Stmt\n", util.SnakeToCamelCase(table))
}

func getPrepareStmtCode(table string) (prepareCode string) {
	var insertSql, cols, vals string

	var columns= getColumnNames(table)
	var numColumns= len(columns)
	cols = strings.Join(columns, ", ")

	placeholders := []string{}
	for j := 0; j < numColumns; j++ {
		placeholders = append(placeholders, "?")
	}

	vals = strings.Join(placeholders, ", ")
	insertSql = "INSERT INTO " + table + " (" + cols + ")" +
		" VALUES (" + vals + ")"

	prepareCode += fmt.Sprintf(
		`
	insertStmt_%s, err = DB.Prepare("%s")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		`, util.SnakeToCamelCase(table), insertSql)

	return
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

	sort.Strings(columns)

	return
}
