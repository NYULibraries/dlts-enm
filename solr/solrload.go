// Copyright © 2018 NYU
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

package solr

import (
	"encoding/json"
	"fmt"
	"log"
	"sort"
	"strings"

	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/util"
)

var epubsNumberOfPages map[string]int

func Load() error {
	epubsNumberOfPages = make(map[string]int)
	for _, epubNumberOfPages := range db.GetEpubsNumberOfPages() {
		epubsNumberOfPages[ epubNumberOfPages.Isbn ] = int(epubNumberOfPages.NumberOfPages)
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

func AddPage(page *db.Page) error {
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

	SortTopicNamesForPage(pageTopicNames)

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

func SortTopicNamesForPage(topicNamesForPage []*db.TopicNamesForPage) {
	sort.Slice(topicNamesForPage, func(i, j int) bool {
		a := util.GetNormalizedTopicNameForSorting(topicNamesForPage[i].TopicDisplayName) +
			util.GetNormalizedTopicNameForSorting(topicNamesForPage[i].TopicName)

		b := util.GetNormalizedTopicNameForSorting(topicNamesForPage[j].TopicDisplayName) +
			util.GetNormalizedTopicNameForSorting(topicNamesForPage[j].TopicName)

		return util.CompareUsingEnglishCollation(a,b) == -1
	} )
}

func pageNumberForDisplay(pageLocalId string) (pageNumber string) {
	return strings.Replace(pageLocalId, "page_", "", 1)
}