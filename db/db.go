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
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nyulibraries/dlts/enm/db/models"
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

func Reload() {
	models.XOLog("Hello XOLog")
}