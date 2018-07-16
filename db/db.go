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
	"sort"
	"strconv"
	_ "github.com/lib/pq"

	models "github.com/nyulibraries/dlts-enm/db/postgres/models"
	"github.com/nyulibraries/dlts-enm/tct"
)

type Topic struct {
	ID          int
	DisplayName string
	DisplayNameSortKey string
}

type EpubsNumberOfPage = models.EpubsNumberOfPages
type EpubsForTopicWithNumberOfMatchedPages = models.EpubsForTopicWithNumberOfMatchedPages
type ExternalRelationsForTopic = models.ExternalRelationsForTopic
type Page = models.Page
type RelatedTopicNamesForTopicWithNumberOfOccurrences = models.RelatedTopicNamesForTopicWithNumberOfOccurrences
type TopicAlternateName = models.TopicAlternateName
type TopicNamesForPage = models.TopicNamesForPage
type TopicNumberOfOccurrences = models.TopicNumberOfOccurrences

var username string
var password string

var DB *sql.DB
var Database string

// Tables in order that they would need to be in when deleting all data table-by-table.
var tables = []string{
	"names",
	"relations",
	"occurrences",
	"topics_weblinks",
	"topics",
	"topic_type",
	"editorial_review_status_state",
	"weblinks",
	"weblinks_relationship",
	"weblinks_vocabulary",
	"scopes",
	"relation_type",
	"relation_direction",
	"locations",
	"epubs",
	"indexpatterns",
}

var editorialReviewStatusStatesMap map[tct.EditorialReviewStatusState]EditorialStatusStateTableRow

type EditorialStatusStateTableRow struct {
	Id int
	State string
}

// Incremented ids for relation_direction table
var relationDirectionId int = 0

// Incremented ids for weblinks_relationship table
var weblinkRelationshipId int = 0

// Incremented ids for weblinks_vocabulary table
var weblinkVocabularyId int = 0

// For initial load of topics table we don't have editorial review status data yet,
// but need to fill in FK column with a valid, non-null value
var defaultEditorialReviewStatusStateId int = 1

func init() {
	Database = os.Getenv("ENM_POSTGRES_DATABASE")
	username = os.Getenv("ENM_POSTGRES_DATABASE_USERNAME")
	password = os.Getenv("ENM_POSTGRES_DATABASE_PASSWORD")

	if Database == "" {
		panic("db: ENM_POSTGRES_DATABASE not set")

	}

	if username == "" {
		panic("db: ENM_POSTGRES_DATABASE_USERNAME not set")

	}

	if password == "" {
		panic("db: ENM_POSTGRES_DATABASE_PASSWORD not set")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@localhost/%s?sslmode=disable", username, password, Database)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}

	editorialReviewStatusStatesMap =
		make(map[tct.EditorialReviewStatusState]EditorialStatusStateTableRow)

	editorialReviewStatusStatesMap[tct.EditorialReviewStatusState{
		ReviewerIsNull: true,
		TimeIsNull: true,
		Changed: false,
		Reviewed: false,
	}] = EditorialStatusStateTableRow{
		Id: defaultEditorialReviewStatusStateId,
		State: "Not Reviewed",
	}
	editorialReviewStatusStatesMap[tct.EditorialReviewStatusState{
		ReviewerIsNull: false,
		TimeIsNull: false,
		Changed: false,
		Reviewed: true,
	}] = EditorialStatusStateTableRow{
		Id: 2,
		State: "Reviewed, Unchanged",
	}
	editorialReviewStatusStatesMap[tct.EditorialReviewStatusState{
		ReviewerIsNull: false,
		TimeIsNull: false,
		Changed: true,
		Reviewed: true,
	}] = EditorialStatusStateTableRow{
		Id: 3,
		State: "Reviewed, Changed",
	}
	editorialReviewStatusStatesMap[tct.EditorialReviewStatusState{
		ReviewerIsNull: false,
		TimeIsNull: false,
		Changed: false,
		Reviewed: false,
	}] = EditorialStatusStateTableRow{
		Id: 4,
		State: "Reviewed Status Revoked",
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

func GetEpubsForTopicWithNumberOfMatchedPages(topicID int) []*EpubsForTopicWithNumberOfMatchedPages{
	epubsForTopicWithNumberOfMatchedPages, err := models.EpubsForTopicWithNumberOfMatchedPagesByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetEpubsForTopicWithNumberOfMatchedPages: " + err.Error())
	}

	return epubsForTopicWithNumberOfMatchedPages
}

func GetEpubsNumberOfPages() (epubsNumberOfPages []*EpubsNumberOfPage) {
	epubsNumberOfPages, err := models.GetEpubsNumberOfPages(DB)
	if err != nil {
		panic("db.GetEpubsNumberOfPages: " + err.Error())
	}

	return epubsNumberOfPages
}

func GetExternalRelationsForTopics(topicID int) []*ExternalRelationsForTopic{
	externalRelationsForTopic, err := models.ExternalRelationsForTopicsByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetExternalRelationsForTopic: " + err.Error())
	}

	return externalRelationsForTopic
}

func GetPagesAll() (pages []*Page) {
	pages, err := models.GetPages(DB)
	if err != nil {
		panic("db.GetPagesAll: " + err.Error())
	}

	return pages
}

func GetPageTopicNamesByPageId(pageId int) (topicNamesForPage []*TopicNamesForPage) {
	topicNamesForPage, err := models.TopicNamesForPagesByPage_id(DB, pageId)
	if err != nil {
		panic("db.GetPageTopicNamesByPageId: " + err.Error())
	}

	return topicNamesForPage
}

func GetRelatedTopicNamesForTopicWithNumberOfOccurrences(topicID int) (relatedTopicsWithNumberOfOccurrences []*RelatedTopicNamesForTopicWithNumberOfOccurrences) {
	// xo-generated model requires 4 topic IDs because there are 4 placeholders
	// in the query for the 4 times the topic ID is needed.
	relatedTopicsWithNumberOfOccurrences, err :=
		models.RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id_1Topic_id_2Topic_id_3Topic_id_4(DB, topicID, topicID, topicID, topicID)
	if err != nil {
		panic("db.GetRelatedTopicNamesForTopicWithNumberOfOccurrences: " + err.Error())
	}

	return relatedTopicsWithNumberOfOccurrences
}

func GetTopicNumberOfOccurrencesByTopicId(topicID int) int {
	topicNumberOfOccurrences, err := models.TopicNumberOfOccurrencesByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetTopicNumberOfOccurrencesByTopicId: " + err.Error())
	}

	if (len(topicNumberOfOccurrences) == 0 ) {
		return 0
	} else {
		return int(topicNumberOfOccurrences[0].NumberOfOccurrences)
	}
}

func GetTopicsWithAlternateNamesAll() (topicsWithAlternateNames []*TopicAlternateName) {
	topicsWithAlternateNames, err := models.GetTopicAlternateNames(DB)
	if err != nil {
		panic("db.GetTopicsWithAlternateNamesAll: " + err.Error())
	}

	return topicsWithAlternateNames
}

func GetTopicsWithAlternateNamesByTopicIDs(topicIDs []string) (topicsWithAlternateNamesByTopicIDs []*TopicAlternateName) {
	sort.Strings(topicIDs)

	topicsWithAlternateNamesByTopicIDs = make([]*TopicAlternateName, 0)

	topicsWithAlternateNamesAll := GetTopicsWithAlternateNamesAll()

	for _, topicWithAlternateName := range topicsWithAlternateNamesAll {
		topicIDToTest := strconv.Itoa(topicWithAlternateName.TctID)
		i := sort.SearchStrings( topicIDs, topicIDToTest )
		if i < len(topicIDs) && topicIDs[i] == topicIDToTest {
			topicsWithAlternateNamesByTopicIDs =
				append(topicsWithAlternateNamesByTopicIDs, topicWithAlternateName)
		}
	}

	return
}

func GetTopicsWithSortKeysByFirstSortableCharacterRegexp(regexp string) (topics []Topic) {
	// sql query
	var sqlstr = `SELECT id, display_name,` +
		` CASE WHEN LEFT( display_name, 1 ) = '"'` +
		`      THEN LOWER( SUBSTR( display_name, 2 ) )` +
		`      ELSE LOWER( display_name )` +
		` END AS display_name_sort_key` +

	` FROM hit_basket` +

	` WHERE` +
	` CASE WHEN LEFT( display_name, 1 ) = '"'` +
	`      THEN LOWER( SUBSTR( display_name, 2, 1) )` +
	`      ELSE LOWER( LEFT( display_name, 1 ) )` +
    ` END SIMILAR TO '` + regexp + `'` +

    ` ORDER BY display_name_sort_key`;

	var (
		id int
		name string
		sortKey string
	)
	rows, err := DB.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&id, &name, &sortKey)
		if err != nil {
			panic(err)
		}
		topics = append(topics, Topic{
			ID: id,
			DisplayName: name,
			DisplayNameSortKey: sortKey,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}

func serialize(stringArray []string) string {
	serializedBytes, err := json.Marshal(stringArray)
	if err != nil {
		panic(err)
	}

	return string(serializedBytes)
}
