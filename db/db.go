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
	"names",
	"relations",
	"topics",
	"scopes",
	"relation_type",
	"relation_direction",
	"locations",
	"epubs",
	"indexpatterns",
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
	var err error
	scopeExists := make(map[int64]bool)
	tctTopics := tct.GetTopicsAll()

	for _, tctTopic := range tctTopics {
		enmTopic := models.Topic{
			TctID: int(tctTopic.ID),
			DisplayNameDoNotUse: tctTopic.DisplayName,
		}

		// Insert into Topics table
		err = enmTopic.Insert(DB)
		if err != nil {
			fmt.Println(tctTopic)
			panic(err)
		}

		// Get topic details
		tctTopicDetail := tct.GetTopicDetail(int(tctTopic.ID))
		tctTopicHits := tctTopicDetail.Basket.TopicHits
		for _, tctTopicHit := range tctTopicHits {
			if ! scopeExists[tctTopicHit.Scope.ID] {
				enmScope := models.Scope{
					TctID: int(tctTopicHit.Scope.ID),
					Scope: tctTopicHit.Scope.Scope,
				}
				err = enmScope.Insert(DB)
				if err != nil {
					fmt.Println(enmScope)
					panic(err)
				}
				scopeExists[tctTopicHit.Scope.ID] = true
			}

			enmName := models.Name{
				TctID: int(tctTopicHit.ID),
				TopicID: int(tctTopic.ID),
				Name: tctTopicHit.Name,
				ScopeID: int(tctTopicHit.Scope.ID),
				Bypass: tctTopicHit.Bypass,
				Hidden: tctTopicHit.Hidden,
				Preferred: tctTopicHit.Preferred,
			}
			err := enmName.Insert(DB)
			if err != nil {fmt.Println(enmName)
				panic(err)
			}
		}
	}

	tctIdTemp := 0
	tctIndexPatterns := tct.GetIndexPatternsAll()
	epubIndexPatternMap := make(map[int]int)
	for _, tctIndexPattern := range tctIndexPatterns {
		tctIdTemp++
		enmIndexPattern := models.Indexpattern{
			TctID: tctIdTemp,
			Name: tctIndexPattern.Name,
			Description: tctIndexPattern.Description,
			PagenumberPreStrings: serialize(tctIndexPattern.PagenumberPreStrings),
			PagenumberCSSSelectorPattern: tctIndexPattern.PagenumberCSSSelectorPattern,
			PagenumberXpathPattern: tctIndexPattern.PagenumberXpathPattern,
			XpathEntry: tctIndexPattern.XpathEntry,
			SeeSplitStrings: serialize(tctIndexPattern.SeeSplitStrings),
			SeeAlsoSplitStrings: serialize(tctIndexPattern.SeeAlsoSplitStrings),
			XpathSee: tctIndexPattern.XpathSee,
			XpathSeeAlso: tctIndexPattern.XpathSeealso,
			SeparatorBetweenSees: tctIndexPattern.SeparatorBetweenSees,
			SeparatorBetweenSeealsos: tctIndexPattern.SeparatorBetweenSeealsos,
			SeparatorSeeSubentry: tctIndexPattern.SeparatorSeeSubentry,
			InlineSeeStart: tctIndexPattern.InlineSeeStart,
			InlineSeeAlsoStart: tctIndexPattern.InlineSeeAlsoStart,
			InlineSeeEnd: tctIndexPattern.InlineSeeEnd,
			InlineSeeAlsoEnd: tctIndexPattern.InlineSeeAlsoEnd,
			SubentryClasses: serialize(tctIndexPattern.SubentryClasses),
			SeparatorBetweenSubentries: tctIndexPattern.SeparatorBetweenSubentries,
			SeparatorBetweenEntryAndOccurrences: tctIndexPattern.SeparatorBetweenEntryAndOccurrences,
			SeparatorBeforeFirstSubentry: tctIndexPattern.SeparatorBeforeFirstSubentry,
			XpathOccurrenceLink: tctIndexPattern.XpathOccurrenceLink,
			IndicatorsOfOccurrenceRange: serialize(tctIndexPattern.IndicatorsOfOccurrenceRange),
		}
		err := enmIndexPattern.Insert(DB)
		if err != nil {
			fmt.Println(enmIndexPattern)
			panic(err)
		}

		epubsUsingIndexPattern := tctIndexPattern.Documents
		for _, epubUsingIndexPattern := range epubsUsingIndexPattern {
			epubIndexPatternMap[int(epubUsingIndexPattern.ID)] = tctIdTemp
		}
	}

	tctEpubs := tct.GetEpubsAll()
	for _, tctEpub := range tctEpubs {
		tctEpubDetail := tct.GetEpubDetail(int(tctEpub.ID))

		enmEpub := models.Epub{
			TctID: int(tctEpub.ID),
			Title: tctEpubDetail.Title,
			Author: tctEpubDetail.Author,
			Publisher: tctEpubDetail.Publisher,
			Isbn: tctEpubDetail.Isbn,
			IndexpatternID: epubIndexPatternMap[int(tctEpub.ID)],
		}
		err := enmEpub.Insert(DB)
		if err != nil {
			fmt.Println(enmEpub)
			panic(err )
		}

		tctLocations := tctEpubDetail.Locations
		for _, tctLocation := range tctLocations {
			tctLocationDetail := tct.GetLocation(int(tctLocation.ID))

			nextId := sql.NullInt64{}
			if tctLocationDetail.PreviousLocationID != nil {
				nextId.Int64 = *tctLocationDetail.PreviousLocationID
				nextId.Valid = true
			} else {
				nextId.Valid = false
			}

			prevId := sql.NullInt64{}
			if tctLocationDetail.PreviousLocationID != nil {
				prevId.Int64 = *tctLocationDetail.PreviousLocationID
				prevId.Valid = true
			} else {
				prevId.Valid = false
			}
			enmLocation := models.Location{
				TctID: int(tctLocationDetail.ID),
				EpubID: int(tctEpub.ID),
				Localid: tctLocationDetail.Localid,
				SequenceNumber: int(tctLocation.SequenceNumber),
				ContentUniqueIndicator: tctLocationDetail.Content.ContentUniqueIndicator,
				ContentDescriptor: tctLocationDetail.Content.ContentDescriptor,
				ContentText: tctLocationDetail.Content.Text,

				// TODO: Create TCT occurrences definitions

				PagenumberFilepath: tctLocationDetail.Pagenumber.Filepath,
				PagenumberTag: tctLocationDetail.Pagenumber.PagenumberTag,
				PagenumberCSSSelector: tctLocationDetail.Pagenumber.CSSSelector,
				PagenumberXpath: tctLocationDetail.Pagenumber.Xpath,

				// TODO: Re-create FKs:
				// CONSTRAINT `fk__locations__next_location_id__locations__tct_id` FOREIGN KEY (`next_location_id`) REFERENCES `locations` (`tct_id`),
				// CONSTRAINT `fk__locations__previous_location_id__locations__tct_id` FOREIGN KEY (`previous_location_id`) REFERENCES `locations` (`tct_id`)
				// ...and have Reload command delete them before loading locations table, then re-create them after load is finished.
				NextLocationID: nextId,
				PreviousLocationID: prevId,
			}
			err := enmLocation.Insert(DB)
			if err != nil {
				fmt.Println(enmLocation)
				panic(err)
			}
		}
	}
}

func serialize(stringArray []string) string {
	serializedBytes, err := json.Marshal(stringArray)
	if err != nil {
		panic(err)
	}

	return string(serializedBytes)
}
