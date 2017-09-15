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
	"strconv"

	_ "github.com/go-sql-driver/mysql"

	"github.com/nyulibraries/dlts-enm/db/models"
	"github.com/nyulibraries/dlts-enm/tct"
)

var username string
var password string

var DB *sql.DB
var Database string

// Tables in order that they would need to be in when deleting all data table-by-table.
var tables = []string{
	"names",
	"relations",
	"occurrences",
	"topics",
	"scopes",
	"relation_type",
	"relation_direction",
	"locations",
	"epubs",
	"indexpatterns",
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

func GetPagesAll() (pages []*models.Page) {
	pages, err := models.GetPages(DB)
	if err != nil {
		panic("db.GetPagesAll: " + err.Error())
	}

	return pages
}

func GetPageTopicNamesByPageId(pageId int) (pageTopicNames []models.PageTopicName) {
	// sql query
	var sqlstr = `SELECT ` +
	`page_id, topic_id, preferred_topic_name, topic_name ` +
	`FROM enm.page_topic_names ` +
	`WHERE page_id = ` + strconv.Itoa(pageId)

	var (
		topicId int
		preferredTopicName string
		topicName string
	)

	rows, err := DB.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pageId, &topicId, &preferredTopicName, &topicName)
		if err != nil {
			panic(err)
		}
		pageTopicNames = append(pageTopicNames, models.PageTopicName{
			PageID: pageId,
			TopicID: topicId,
			PreferredTopicName: preferredTopicName,
			TopicName: topicName,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}

func GetTopicsAll() (topics []models.Topic) {
	const sqlstr = `SELECT tct_id, display_name_do_not_use FROM topics ORDER BY display_name_do_not_use`

	var (
		id int
		name string
	)
	rows, err := DB.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		topics = append(topics, models.Topic{
			TctID: id,
			DisplayNameDoNotUse: name,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}

func GetTopicSubEntries(topicId int) (subentryTopics []models.Topic){
	const sqlstr = `SELECT t2.tct_id, t2.display_name_do_not_use
FROM topics t1
	INNER JOIN relations     r  on t1.tct_id          = r.role_from_topic_id
	INNER JOIN topics        t2 on t2.tct_id          = r.role_to_topic_id
	INNER JOIN relation_type rt on r.relation_type_id = rt.tct_id
WHERE rt.rtype = 'Subentry'
AND   t1.tct_id = ?
ORDER BY t2.display_name_do_not_use`

	var (
		id int
		name string
	)
	rows, err := DB.Query(sqlstr, topicId)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		subentryTopics = append(subentryTopics, models.Topic{
			TctID: id,
			DisplayNameDoNotUse: name,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}

func Reload() {
	var err error
	relationDirectionIds := make(map[string]int)
	relationDirectionId := 0
	relationExists := make(map[int64]bool)
	relationTypeExists := make(map[int64]bool)
	scopeExists := make(map[int64]bool)
	tctTopics := tct.GetTopicsAll()

	// All Topics must be loaded first before attempting to load everything
	// that comes in from TopicDetails, the reason being topics table is
	// target of FKs in Relations.
	for _, tctTopic := range tctTopics {
		// TODO: remove display_name from table after figure out how to
		// fix or workaround xo bug that prevents creation of full model
		// code if tct_id is the only column.
		enmTopic := models.Topic{
			TctID:               int(tctTopic.ID),
			DisplayNameDoNotUse: tctTopic.DisplayName,
		}

		// Insert into Topics table
		err = enmTopic.Insert(DB)
		if err != nil {
			fmt.Println(tctTopic)
			panic(err)
		}
	}

	// Second pass: get topic details
	for _, tctTopic := range tctTopics {
		tctTopicDetail := tct.GetTopicDetail(int(tctTopic.ID))

		tctTopicRelations := tctTopicDetail.Relations
		for _, tctTopicRelation := range tctTopicRelations {
			if ! relationExists[tctTopicRelation.ID] {
				tctRelationType := tctTopicRelation.Relationtype
				if ! relationTypeExists[tctRelationType.ID] {
					enmRelationType := models.RelationType{
						TctID: int(tctRelationType.ID),
						Rtype: tctRelationType.Rtype,
						RoleFrom: tctRelationType.RoleFrom,
						RoleTo: tctRelationType.RoleTo,
						Symmetrical: tctRelationType.Symmetrical,
					}
					err = enmRelationType.Insert(DB)
					if err != nil {
						fmt.Println(enmRelationType)
						panic(err)
					}
					relationTypeExists[tctRelationType.ID] = true
				}

				tctRelationDirection := tctTopicRelation.Direction
				if relationDirectionIds[tctRelationDirection] == 0 {
					relationDirectionId++
					enmRelationDirection := models.RelationDirection{
						ID:        relationDirectionId,
						Direction: tctRelationDirection,
					}
					err = enmRelationDirection.Insert(DB)
					if err != nil {
						fmt.Println(enmRelationDirection)
						panic(err)
					}
					relationDirectionIds[tctRelationDirection] = relationDirectionId
				}

				var roleFromTopicId int
				var roleToTopicId int
				if tctTopicRelation.Direction == "source" {
					roleFromTopicId = int(tctTopicRelation.Basket.ID)
					roleToTopicId = int(tctTopic.ID)
				} else if tctTopicRelation.Direction == "destination" {
					roleFromTopicId = int(tctTopic.ID)
					roleToTopicId = int(tctTopicRelation.Basket.ID)
				} else {
					panic( "Unknown relation direction: " + tctTopicRelation.Direction)
				}
				enmRelation := models.Relation{
					TctID: int(tctTopicRelation.ID),
					RelationTypeID: int(tctRelationType.ID),
					RelationDirectionID: relationDirectionIds[tctRelationDirection],
					RoleFromTopicID: roleFromTopicId,
					RoleToTopicID: roleToTopicId,
				}
				err = enmRelation.Insert(DB)
				if err != nil {
					fmt.Println(err)
					panic(err)
				}

				relationExists[tctTopicRelation.ID] = true
			}
		}

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

	// TODO: Find out if there is an internal TCT ID for IndexPatterns, and if
	// so, maybe ask Infoloom to include it in the API responses instead of
	// making our own ids.
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

		// All Locations must be loaded first before attempting to load Occurrences,
		// the reason being Occurrences target Locations for RingNext and RingPrevious FKs
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

				// Occurrences are processed in the second pass

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

		// Second pass: get Occurrences
		for _, tctLocation := range tctLocations {
			tctLocationDetail := tct.GetLocation(int(tctLocation.ID))

			for _, tctOccurrence := range tctLocationDetail.Occurrences {
				ringNextId := sql.NullInt64{}
				if tctOccurrence.RingNext != nil {
					ringNextId.Int64 = *tctOccurrence.RingNext
					ringNextId.Valid = true
				} else {
					ringNextId.Valid = false
				}

				ringPrevId := sql.NullInt64{}
				if tctOccurrence.RingPrevious != nil {
					ringPrevId.Int64 = *tctOccurrence.RingPrevious
					ringPrevId.Valid = true
				} else {
					ringPrevId.Valid = false
				}

				// TODO: Re-create FKs after figuring why inserting
				// NULL into them is not working:
				// * fk__occurrences__ring_next__locations__tct_id
				// * fk__occurrences__ring_prev__locations__tct_id
				enmOccurrence := models.Occurrence{
					TctID: int(tctOccurrence.ID),
					LocationID: int(tctLocation.ID),
					TopicID: int(tctOccurrence.Basket.ID),
					RingNext: ringNextId,
					RingPrev: ringPrevId,
				}
				err = enmOccurrence.Insert(DB)
				if err != nil {
					fmt.Println(enmOccurrence)
					panic(err)
				}
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
