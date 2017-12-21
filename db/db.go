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
	"regexp"
	"strconv"
	"strings"

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

func GetEpubsForTopicWithNumberOfMatchedPages(topicID int) []*models.EpubsForTopicWithNumberOfMatchedPages{
	epubsForTopicWithNumberOfMatchedPages, err := models.EpubsForTopicWithNumberOfMatchedPagesByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetEpubsForTopicWithNumberOfMatchedPages: " + err.Error())
	}

	return epubsForTopicWithNumberOfMatchedPages
}

func GetEpubsNumberOfPages() (epubsNumberOfPages []*models.EpubsNumberOfPage) {
	epubsNumberOfPages, err := models.GetEpubsNumberOfPages(DB)
	if err != nil {
		panic("db.GetEpubsNumberOfPages: " + err.Error())
	}

	return epubsNumberOfPages
}

func GetExternalRelationsForTopics(topicID int) []*models.ExternalRelationsForTopic{
	externalRelationsForTopic, err := models.ExternalRelationsForTopicsByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetExternalRelationsForTopic: " + err.Error())
	}

	return externalRelationsForTopic
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
	var sqlstr = `SELECT` +
	` page_id, topic_id, topic_display_name, topic_name` +
	` FROM enm.page_topic_names` +
	` WHERE page_id = ` + strconv.Itoa(pageId) +
	` ORDER BY page_id, topic_id, topic_display_name, topic_name`

	var (
		topicId int
		topicDisplayName string
		topicName string
	)

	rows, err := DB.Query(sqlstr)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&pageId, &topicId, &topicDisplayName, &topicName)
		if err != nil {
			panic(err)
		}
		pageTopicNames = append(pageTopicNames, models.PageTopicName{
			PageID: pageId,
			TopicID: topicId,
			TopicDisplayName: topicDisplayName,
			TopicName: topicName,
		})
	}
	err = rows.Err()
	if err != nil {
		panic(err)
	}

	return
}

func GetRelatedTopicNamesForTopicWithNumberOfOccurrences(topicID int) (relatedTopicsWithNumberOfOccurrences []*models.RelatedTopicNamesForTopicWithNumberOfOccurrences) {
	relatedTopicsWithNumberOfOccurrences, err := models.RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetRelatedTopicNamesForTopicWithNumberOfOccurrences: " + err.Error())
	}

	return relatedTopicsWithNumberOfOccurrences
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

func GetTopicNumberOfOccurrencesByTopicId(topicID int) int {
	topicNumberOfOccurrences, err := models.TopicNumberOfOccurrencesByTopic_id(DB, topicID)
	if err != nil {
		panic("db.GetTopicNumberOfOccurrencesByTopicId: " + err.Error())
	}

	return int(topicNumberOfOccurrences[0].NumberOfOccurrences)
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

func GetTopicsWithAlternateNamesAll() (topicsWithAlternateNames []*models.TopicAlternateName) {
	topicsWithAlternateNames, err := models.GetTopicAlternateNames(DB)
	if err != nil {
		panic("db.GetTopicsWithAlternateNamesAll: " + err.Error())
	}

	return topicsWithAlternateNames
}

func Reload() {
	loadEditorialReviewStatusStates()

	loadRelationTypes()

	loadTopicTypes()

	tctTopics := loadTopicsFirstPass()

	loadTopicsSecondPass(tctTopics)

	epubIndexPatternMap := loadIndexPatterns()

	loadEpubs(epubIndexPatternMap)
}

func loadEditorialReviewStatusStates() {
	for _, row := range editorialReviewStatusStatesMap {
		enmEditorialReviewStatusState := models.EditorialReviewStatusState{
			ID: row.Id,
			State: row.State,
		}

		err := enmEditorialReviewStatusState.Insert(DB)
		if err != nil {
			fmt.Println(row.State)
			panic(err)
		}
	}
}

func loadRelationTypes() {
	var tctRelationTypes = tct.GetRelationTypesAll()

	for _, tctRelationType := range tctRelationTypes {
		enmRelationType := models.RelationType{
			TctID:       int(tctRelationType.ID),
			Rtype:       tctRelationType.Rtype,
			RoleFrom:    tctRelationType.RoleFrom,
			RoleTo:      tctRelationType.RoleTo,
			Symmetrical: tctRelationType.Symmetrical,
		}

		// Insert into relation types table
		err := enmRelationType.Insert(DB)
		if err != nil {
			fmt.Println(tctRelationType)
			panic(err)
		}
	}
}

func loadTopicTypes() {
	var tctTopicTypes = tct.GetTopicTypesAll()

	for _, tctTopicType := range tctTopicTypes {
		enmTopicType := models.TopicType{
			TctID:       int(tctTopicType.ID),
			Ttype:       tctTopicType.Ttype,
		}

		// Insert into topic types table
		err := enmTopicType.Insert(DB)
		if err != nil {
			fmt.Println(tctTopicType)
			panic(err)
		}
	}
}

func loadTopicsFirstPass() []tct.Topic {
	tctTopics := tct.GetTopicsAll()

	var err error

	// All Topics must be loaded first before attempting to load everything
	// that comes in from TopicDetails, the reason being topics table is
	// target of FKs in Relations.  Note that editorial review status columns
	// are not filled in until the second pass when topic details are fetched,
	// with the exception of FK column editorial_review_status_state_id, which
	// we pre-fill with temporary value defaultEditorialReviewStatusStateId.
	for _, tctTopic := range tctTopics {
		// TODO: originally had display_name_do_not_use column in topics table
		// due to xo bug that prevents creation of full model code if tct_id was
		// the only column.  Now there are several more FK columns, so
		// this display name colume could be dropped.  It's proven to be convenient
		// though, and there might be some advantage to keeping it because it
		// represents an end calculation done by TCT.  Originally we were thinking
		// we'd be making the determination of which name should be the display
		// name post-TCT, but do we actually need to do that?  And even if we do,
		// we could still keep this TCT choice.
		enmTopic := models.Topic{
			TctID:               int(tctTopic.ID),
			DisplayNameDoNotUse: tctTopic.DisplayName,
			EditorialReviewStatusStateID: defaultEditorialReviewStatusStateId,
		}

		// Insert into Topics table
		err = enmTopic.Insert(DB)
		if err != nil {
			fmt.Println(tctTopic)
			panic(err)
		}
	}

	return tctTopics
}

func loadTopicsSecondPass(tctTopics []tct.Topic) {
	relationDirectionIds := make(map[string]int)
	relationExists := make(map[int64]bool)

	scopeExists := make(map[int64]bool)

	weblinkRelationshipIds := make(map[string]int)
	weblinkVocabularyIds := make(map[string]int)
	weblinkExists := make(map[int64]bool)

	for _, tctTopic := range tctTopics {
		tctTopicDetail := tct.GetTopicDetail(int(tctTopic.ID))

		loadRelations(tctTopicDetail, relationExists, relationDirectionIds)

		setEditorialReviewStatus(tctTopicDetail)

		loadNames(tctTopicDetail, scopeExists)

		loadWeblinks(tctTopicDetail,
			weblinkExists,
			weblinkRelationshipIds,
			weblinkVocabularyIds)
	}
}

func loadRelations(
	tctTopicDetail tct.TopicDetail,
	relationExists map[int64]bool,
	relationDirectionIds map[string]int) {

	tctTopicRelations := tctTopicDetail.Relations

	for _, tctTopicRelation := range tctTopicRelations {
		if ! relationExists[tctTopicRelation.ID] {
			tctRelationType := tctTopicRelation.Relationtype
			tctRelationDirection := tctTopicRelation.Direction
			if relationDirectionIds[tctRelationDirection] == 0 {
				relationDirectionId++
				enmRelationDirection := models.RelationDirection{
					ID:        relationDirectionId,
					Direction: tctRelationDirection,
				}
				err := enmRelationDirection.Insert(DB)
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
				roleToTopicId = int(tctTopicDetail.Basket.ID)
			} else if tctTopicRelation.Direction == "destination" {
				roleFromTopicId = int(tctTopicDetail.Basket.ID)
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
			err := enmRelation.Insert(DB)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			relationExists[tctTopicRelation.ID] = true
		}
	}
}

func setEditorialReviewStatus(tctTopicDetail tct.TopicDetail) {
	tctReview := tctTopicDetail.Basket.Review
	tctEditorialReviewStatusState := tct.EditorialReviewStatusState{
		ReviewerIsNull: tctReview.Reviewer == "",
		TimeIsNull:     tctReview.Time == "",
		Changed:        tctReview.Changed,
		Reviewed:       tctReview.Reviewed,
	}
	editorialReviewStatusStateId :=
		editorialReviewStatusStatesMap[tctEditorialReviewStatusState].Id

	// Update topic
	enmTopic, err := models.TopicByTctID(DB, int(tctTopicDetail.Basket.ID))
	if err != nil {
		panic(fmt.Sprintf("ERROR: could not fetch topic %d", int(tctTopicDetail.Basket.ID)))
	}

	enmTopic.EditorialReviewStatusReviewer = sql.NullString{
		String: tctReview.Reviewer,
		Valid: true,
	}
	enmTopic.EditorialReviewStatusTime = sql.NullString{
		String: tctReview.Time,
		Valid: true,
	}
	enmTopic.EditorialReviewStatusStateID = editorialReviewStatusStateId

	enmTopic.Update(DB)
}

func loadNames(tctTopicDetail tct.TopicDetail, scopeExists map[int64]bool) {
	tctTopicHits := tctTopicDetail.Basket.TopicHits
	for _, tctTopicHit := range tctTopicHits {
		if ! scopeExists[tctTopicHit.Scope.ID] {
			enmScope := models.Scope{
				TctID: int(tctTopicHit.Scope.ID),
				Scope: tctTopicHit.Scope.Scope,
			}
			err := enmScope.Insert(DB)
			if err != nil {
				fmt.Println(enmScope)
				panic(err)
			}
			scopeExists[tctTopicHit.Scope.ID] = true
		}

		enmName := models.Name{
			TctID: int(tctTopicHit.ID),
			TopicID: int(tctTopicDetail.Basket.ID),
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

func loadWeblinks(tctTopicDetail tct.TopicDetail,
	weblinkExists map[int64]bool,
	weblinkRelationshipIds map[string]int,
	weblinkVocabularyIds map[string]int) {

	tctWeblinks := tctTopicDetail.Basket.Weblinks

	for _, tctWeblink := range tctWeblinks {
		if ! weblinkExists[tctWeblink.ID] {
			// Example tctWeblinkContent: "Library of Congress (exactMatch)"
			// vocabulary = "Library of Congress"
			// relationship = "exactMatch"
			re := regexp.MustCompile(`^([^(]+)\(([^)]+)\)$`)
			matches := re.FindStringSubmatch(tctWeblink.Content)
			tctWeblinkVocabulary := strings.TrimSpace(matches[1])
			tctWeblinkRelationship := strings.TrimSpace(matches[2])

			if weblinkRelationshipIds[tctWeblinkRelationship] == 0 {
				weblinkRelationshipId++
				enmWeblinkRelationship := models.WeblinksRelationship{
					ID:        weblinkRelationshipId,
					Relationship: tctWeblinkRelationship,
				}
				err := enmWeblinkRelationship.Insert(DB)
				if err != nil {
					fmt.Println(enmWeblinkRelationship)
					panic(err)
				}
				weblinkRelationshipIds[tctWeblinkRelationship] = weblinkRelationshipId
			}

			if weblinkVocabularyIds[tctWeblinkVocabulary] == 0 {
				weblinkVocabularyId++
				enmWeblinkVocabulary := models.WeblinksVocabulary{
					ID:        weblinkVocabularyId,
					Vocabulary: tctWeblinkVocabulary,
				}
				err := enmWeblinkVocabulary.Insert(DB)
				if err != nil {
					fmt.Println(enmWeblinkVocabulary)
					panic(err)
				}
				weblinkVocabularyIds[tctWeblinkVocabulary] = weblinkVocabularyId
			}

			enmWeblink := models.Weblink{
				TctID:                  int(tctWeblink.ID),
				URL:                    tctWeblink.URL,
				WeblinksRelationshipID: weblinkRelationshipId,
				WeblinksVocabularyID:   weblinkVocabularyId,
			}
			err := enmWeblink.Insert(DB)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}

			weblinkExists[tctWeblink.ID] = true

			// Create topic to weblink mapping
			enmTopicsWeblinks := models.TopicsWeblink{
				TopicID: int(tctTopicDetail.Basket.ID),
				WeblinkID: int(tctWeblink.ID),
			}
			enmTopicsWeblinks.Insert(DB)
			if err != nil {
				fmt.Println(err)
				panic(err)
			}
		}
	}
}


func loadIndexPatterns() (epubIndexPatternMap map[int]int) {
	// TODO: Find out if there is an internal TCT ID for IndexPatterns, and if
	// so, maybe ask Infoloom to include it in the API responses instead of
	// making our own ids.
	tctIdTemp := 0
	tctIndexPatterns := tct.GetIndexPatternsAll()
	epubIndexPatternMap = make(map[int]int)
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

	return
}

func loadEpubs(epubIndexPatternMap map[int]int) {
	tctEpubs := tct.GetEpubsAll()
	for _, tctEpub := range tctEpubs {
		epubId := int(tctEpub.ID)

		tctEpubDetail := tct.GetEpubDetail(epubId)

		insertEpub(tctEpubDetail, epubId, epubIndexPatternMap[epubId])

		// All Locations must be loaded first before attempting to load Occurrences,
		// the reason being Occurrences target Locations for RingNext and RingPrevious FKs
		locationIds := loadLocations(tctEpubDetail, epubId)

		loadOccurrences(locationIds)
	}
}

func insertEpub(tctEpubDetail tct.EpubDetail, epubId int, indexPatternId int) {
	enmEpub := models.Epub{
		TctID: epubId,
		Title: tctEpubDetail.Title,
		Author: tctEpubDetail.Author,
		Publisher: tctEpubDetail.Publisher,
		Isbn: tctEpubDetail.Isbn,
		IndexpatternID: indexPatternId,
	}
	err := enmEpub.Insert(DB)
	if err != nil {
		fmt.Println(enmEpub)
		panic(err )
	}
}

func loadLocations(tctEpubDetail tct.EpubDetail, epubId int) (locationIds []int) {
	tctLocations := tctEpubDetail.Locations

	for _, tctLocation := range tctLocations {
		tctLocationDetail := tct.GetLocation(int(tctLocation.ID))

		nextId := sql.NullInt64{}
		if tctLocationDetail.NextLocationID != nil {
			nextId.Int64 = *tctLocationDetail.NextLocationID
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
			EpubID: int(epubId),
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

		locationIds = append(locationIds, int(tctLocationDetail.ID))
	}

	return
}

func loadOccurrences(locationIds []int) {
	for _, locationId := range locationIds {
		tctLocationDetail := tct.GetLocation(locationId)

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
				LocationID: int(locationId),
				TopicID: int(tctOccurrence.Basket.ID),
				RingNext: ringNextId,
				RingPrev: ringPrevId,
			}
			err := enmOccurrence.Insert(DB)
			if err != nil {
				fmt.Println(enmOccurrence)
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
