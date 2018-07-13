package solrmysql

import (
	"encoding/json"
	"fmt"
	"log"
	"strings"

	db "github.com/nyulibraries/dlts-enm/db/mysql"
	"github.com/nyulibraries/dlts-enm/db/mysql/models"
)

var epubsNumberOfPages map[string]int

func Load() error {
	epubsNumberOfPages = make(map[string]int)
	for _, epubNumberOfPage := range db.GetEpubsNumberOfPages() {
		epubsNumberOfPages[ epubNumberOfPage.Isbn ] = int(epubNumberOfPage.NumberOfPages)
	}

	pages := db.GetPagesAll()
	for _, page := range pages {
		err := AddPage(page)
		if err != nil {
			panic(fmt.Sprintf("ERROR: couldn't add page %d to Solr index: \"%s\"\n", page.ID, err.Error()))
		}
	}

	return nil
}

func AddPage(page *models.Page) error {
	var topicNames []string
	// Using underscore in variable name to match Solr field name
	var topicNames_facet []string
	// This is what actually gets marshaled to JSON, it will be an ordered slice
	// of slices:
	// [
	//     [displayName1, alternateName1, alternateName2, ...],
	//     [displayName2, alternateName1, alternateName2, ...],
	//     [displayName3, alternateName1, alternateName2, ...],
	//     ...
	// ]
	// ...where displayNames are ordered using NYUP-376 rules.
	var topicNamesForDisplayArray [][]string = [][]string{}
	// This is an intermediate data structure to make it easier to create
	// topicNamesForDisplayArray
	var topicNamesForDisplayMap map[string][]string = make(map[string][]string)

	var sortedTopicDisplayNames []string

	pageTopicNames := db.GetPageTopicNamesByPageId(page.ID)

	// We are assuming that the pageTopicNames are sorted by TopicDisplayNameSortKey
	// and sub-sorted by TopicNameSortKey, which themselves are meant to implement
	// NYUP-376 sorting rules.
	for _, pageTopic := range pageTopicNames {
		if _, ok := topicNamesForDisplayMap[pageTopic.TopicDisplayName]; ! ok {
			topicNamesForDisplayMap[pageTopic.TopicDisplayName] = []string{}
			sortedTopicDisplayNames = append(sortedTopicDisplayNames, pageTopic.TopicDisplayName)
		}

		if pageTopic.TopicName != pageTopic.TopicDisplayName {
			topicNamesForDisplayMap[pageTopic.TopicDisplayName] =
				append(topicNamesForDisplayMap[pageTopic.TopicDisplayName], pageTopic.TopicName)
			topicNames = append(topicNames, pageTopic.TopicName)
		} else {
			topicNames = append(topicNames, pageTopic.TopicDisplayName)
			topicNames_facet = append(topicNames_facet, pageTopic.TopicDisplayName)
		}
	}

	// Add alternate names to topicNamesForDisplayArray
	for _, topicDisplayName := range sortedTopicDisplayNames {
		alternateNames := topicNamesForDisplayMap[topicDisplayName]
		// "Prepend" topicDisplayName to alternateNames and add to topicNamesForDisplayArray
		topicNamesForDisplayArray = append(topicNamesForDisplayArray,
			append([]string{topicDisplayName}, alternateNames...))
	}

	// Map of topic display names to alternate names as marshalled JSON
	topicNamesForDisplayBytes, err := json.Marshal(topicNamesForDisplayArray)
	if err != nil {
		panic(fmt.Sprintf("ERROR: couldn't marshal topicNamesForDisplay for %d", page.ID))
	}

	doc := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id": page.ID,
				"isbn": page.Isbn,
				"authors": page.Authors,
				"epubNumberOfPages": epubsNumberOfPages[ page.Isbn ],
				"publisher": page.Publisher,
				"title": page.Title,
				"yearOfPublication": 0,
				"pageLocalId": page.PageLocalid,
				"pageNumberForDisplay": pageNumberForDisplay(page.PageLocalid),
				"pageSequenceNumber": page.PageSequence,
				"pageText": page.PageText,
				"topicNames": topicNames,
				"topicNames_facet": topicNames_facet,
				"topicNamesForDisplay": string(topicNamesForDisplayBytes),
			},
		},
	}

	_, err = conn.Update(doc, true)
	if err != nil {
		return err
	}

	log.Printf("Added page %d\n", page.ID)

	return nil
}

func pageNumberForDisplay(pageLocalId string) (pageNumber string) {
	return strings.Replace(pageLocalId, "page_", "", 1)
}