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
	"encoding/json"
	"fmt"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/nyulibraries/dlts/enm/db/models"
	"github.com/nyulibraries/dlts/enm/tct"
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
	// Topics table: just need the info from topics All Topics endpoint

	// Scopes table: tricky -- the only way to get the scope IDs is by
	//     retrieving the TopicDetail for every topic.

	//var err error
	//scopeExists := make(map[int64]bool)
	//tctTopics := tct.GetTopicsAll()
	//
	//for _, tctTopic := range tctTopics {
	//	enmTopic := models.Topic{
	//		TctID: int(tctTopic.ID),
	//		DisplayNameDoNotUse: tctTopic.DisplayName,
	//	}
	//
	//	// Insert into Topics table
	//	err = enmTopic.Insert(DB)
	//	if err != nil {
	//		fmt.Println(tctTopic)
	//		panic(err)
	//	}
	//
	//	// Get topic details
	//	tctTopicDetail := tct.GetTopicDetail(int(tctTopic.ID))
	//	tctTopicHits := tctTopicDetail.Basket.TopicHits
	//	for _, tctTopicHit := range tctTopicHits {
	//		if ! scopeExists[tctTopicHit.Scope.ID] {
	//			enmScope := models.Scope{
	//				TctID: int(tctTopicHit.Scope.ID),
	//				Scope: tctTopicHit.Scope.Scope,
	//			}
	//			err = enmScope.Insert(DB)
	//			if err != nil {
	//				fmt.Println(enmScope)
	//				panic(err)
	//			}
	//			scopeExists[tctTopicHit.Scope.ID] = true
	//		}
	//	}
	//}

	tctIndexPatterns := tct.GetIndexPatternsAll()
	for _, tctIndexPattern := range tctIndexPatterns {
		//enmIndexPattern := new models.Indexpattern{
		//	TctID: 0,
		//	Name: tctIndexPattern.Name,
		//	Description: tctIndexPattern.Description,
		//	IndicatorsOfOccurrenceRange: tctIndexPattern.IndicatorsOfOccurrenceRange,
		//	InlineSeeAlsoEnd: tctIndexPattern.InlineSeeAlsoEnd,
		//	PagenumberPreStrings: tctIndexPattern.PagenumberPreStrings
		//}
		models.XOLog("")
		fmt.Println(tctIndexPattern)
	}

	//tctEpubs := tct.GetEpubsAll()
	//for _, tctEpub := range tctEpubs {
	//	tctEpubDetail := tct.GetEpubDetail(int(tctEpub.ID))
	//	fmt.Println(tctEpubDetail)
	//
	//	tctLocations := tctEpubDetail.Locations
	//	for _, tctLocation := range tctLocations {
	//		tctLocationDetail := tct.GetLocation(int(tctLocation.ID))
	//		fmt.Println(tctLocationDetail)
	//		models.XOLog("")
	//	}
	//}
}

func serialize(stringArray []string) string {
	serializedBytes, err := json.Marshal(stringArray)
	if err != nil {
		panic(err)
	}

	return string(serializedBytes)
}
