package solr

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/nyulibraries/dlts-enm/cache"
	"github.com/nyulibraries/dlts-enm/db"
	"github.com/nyulibraries/dlts-enm/util"
)

type SolrDoc map[string]interface{}

var epubsNumberOfPages map[string]int

func Load() error {
	if Source == "database" {
		return LoadFromDatabase()
	} else if Source == "cache" {
		return LoadFromCache()
	} else {
		panic(fmt.Sprintf("ERROR: \"%s\" is not a valid --source option.", Source))
	}
}

func LoadFromCache() error {
	err := filepath.Walk(cache.SolrLoadCache, func(path string, info os.FileInfo, err error) error {
		if (err != nil) {
			return err
		}
		if (! info.IsDir() && strings.HasSuffix(info.Name(), ".json")) {
			var solrDocWithFloat64Values SolrDoc

			jsonBytes, err := ioutil.ReadFile(path)
			if err != nil {
				panic(err)
			}

			err = json.Unmarshal(jsonBytes, &solrDocWithFloat64Values)
			if err != nil {
				panic(err)
			}

			addDoc := solrDocWithFloat64Values["add"].([]interface{})[0].(map[string]interface{})

			solrDoc := SolrDoc{
				"add": []interface{}{
					map[string]interface{}{
						"id": int(addDoc["id"].(float64)),
						"isbn": addDoc["isbn"],
						"authors": addDoc["authors"],
						"epubNumberOfPages": addDoc["epubNumberOfPages"],
						"publisher": addDoc["publisher"],
						"title": addDoc["title"],
						"yearOfPublication": 0,
						"pageLocalId": addDoc["pageLocalId"],
						"pageNumberForDisplay": pageNumberForDisplay(addDoc["pageLocalId"].(string)),
						"pageSequenceNumber": addDoc["pageSequenceNumber"],
						"pageText": addDoc["pageText"],
						"topicNames": addDoc["topicNames"],
						"topicNames_facet": addDoc["topicNames_facet"],
						"topicNamesForDisplay": addDoc["topicNamesForDisplay"],
					},
				},
			}

			Update(solrDoc)
		}

		return nil
	})
	if (err != nil) {
		return err
	}

	return nil
}

func LoadFromDatabase() error {
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

	doc := SolrDoc{
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

	WriteCacheFile(doc)

	err = Update(doc)
	if err != nil {
		return err
	}

	log.Printf("Added page %d\n", page.ID)

	return nil
}

func SortTopicNamesForPage(topicNamesForPage []*db.TopicNamesForPage) {
	sort.SliceStable(topicNamesForPage, func(i, j int) bool {
		// Note that i is the higher index, it is j+1
		a := util.GetNormalizedTopicNameForSorting(topicNamesForPage[i].TopicDisplayName) +
			util.GetNormalizedTopicNameForSorting(topicNamesForPage[i].TopicName)

		b := util.GetNormalizedTopicNameForSorting(topicNamesForPage[j].TopicDisplayName) +
			util.GetNormalizedTopicNameForSorting(topicNamesForPage[j].TopicName)

		return util.CompareUsingEnglishCollation(a,b) == -1
	} )
}

func Update(doc SolrDoc) error {
	_, err := conn.Update(doc, true)
	if err != nil {
		return err
	}

	return nil
}

func WriteCacheFile(solrDoc SolrDoc) (err error){
	solrDocJSON, err := json.MarshalIndent(solrDoc,"","    ")
	if err != nil {
		panic(err)
	}

	addDocArray := solrDoc["add"].([]interface{})
	doc := addDocArray[0].(map[string]interface{})

	cacheFile := cache.SolrLoadCacheFile(
		doc["isbn"].(string), doc["pageNumberForDisplay"].(string), doc["id"].(int))
	f, err := util.CreateFileWithAllParentDirectories(cacheFile)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(solrDocJSON)
	if err != nil {
		panic(err)
	}

	return nil
}

func pageNumberForDisplay(pageLocalId string) (pageNumber string) {
	return strings.Replace(pageLocalId, "page_", "", 1)
}
