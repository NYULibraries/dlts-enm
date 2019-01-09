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
	"sort"
	"strconv"
	_ "github.com/lib/pq"

	"github.com/nyulibraries/dlts-enm/db/postgres/models"
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

var DB *sql.DB

func init() {
	database := os.Getenv("ENM_POSTGRES_DATABASE")
	hostname := os.Getenv("ENM_POSTGRES_DATABASE_HOSTNAME")
	username := os.Getenv("ENM_POSTGRES_DATABASE_USERNAME")
	password := os.Getenv("ENM_POSTGRES_DATABASE_PASSWORD")

	if database == "" {
		panic("db: ENM_POSTGRES_DATABASE not set")
	}

	if hostname == "" {
			panic("db: ENM_POSTGRES_DATABASE_HOSTNAME not set")
	}

	if username == "" {
		panic("db: ENM_POSTGRES_DATABASE_USERNAME not set")
	}

	if password == "" {
		panic("db: ENM_POSTGRES_DATABASE_PASSWORD not set")
	}

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable", username, password, hostname, database)
	var err error
	DB, err = sql.Open("postgres", connStr)
	if err != nil {
		panic(err.Error())
	}

	if err = DB.Ping(); err != nil {
		panic(err.Error())
	}
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

	// Note that duplicate display_name_sort_key values should not be possible
	// because topics are not supposed to share display names, but we do in fact
	// have some duplicates in TCT for whatever reason, and there is no technical
	// safeguard against it happening again in the future even after the current
	// duplicates are fixed.  We need a sub-sort to keep order deterministic for
	// tests.
    ` ORDER BY display_name_sort_key, id`;

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
