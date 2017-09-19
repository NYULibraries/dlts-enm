package solrutil

import (
	"log"
	"strings"

	"github.com/nyulibraries/dlts-enm/db/models"
	solr "github.com/rtt/Go-Solr"
	"github.com/nyulibraries/dlts-enm/db"
	"encoding/json"
	"fmt"
)

var Port int
var Server string
var conn *solr.Connection

type topicNamesForDisplay map[string][]string

func Init(server string, port int) error {
	var err error
	conn, err = solr.Init(server, port, "enm-pages")
	if err != nil {
		return err
	}

	return nil
}

func AddPage(page *models.Page) error {
	var topicNames []string
	var topicNamesForDisplay topicNamesForDisplay = make(map[string][]string)

	pageTopicNames := db.GetPageTopicNamesByPageId(page.ID)

	// We are assuming that the pageTopicNames are alpha-sorted by PreferredTopicName
	// and then TopicName within each topic.
	for _, pageTopic := range pageTopicNames {
		if _, ok := topicNamesForDisplay[pageTopic.PreferredTopicName]; ! ok {
			topicNamesForDisplay[pageTopic.PreferredTopicName] = []string{}
		}

		if pageTopic.TopicName != pageTopic.PreferredTopicName {
			topicNamesForDisplay[pageTopic.PreferredTopicName] =
				append(topicNamesForDisplay[pageTopic.PreferredTopicName], pageTopic.TopicName)
			topicNames = append(topicNames, pageTopic.TopicName)
		} else {
			topicNames = append(topicNames, pageTopic.PreferredTopicName)
		}
	}

	topicNamesForDisplayBytes, err := json.Marshal(topicNamesForDisplay)
	if err != nil {
		panic(fmt.Sprintf("ERROR: couldn't marshal topicNamesForDisplay for %s", page.ID))
	}

	doc := map[string]interface{}{
		"add": []interface{}{
			map[string]interface{}{
				"id": page.ID,
				"isbn": page.Isbn,
				"authors": page.Authors,
				"publisher": page.Publisher,
				"title": page.Title,
				"yearOfPublication": 0,
				"pageLocalId": page.PageLocalid,
				"pageNumberForDisplay": pageNumberForDisplay(page.PageLocalid, page.PagePattern),
				"pageSequenceNumber": page.PageSequence,
				"pageText": page.PageText,
				"topicNames": topicNames,
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

func pageNumberForDisplay(pageLocalId string, pagePattern string) (pageNumber string) {
	pagePattern = strings.Replace(pagePattern, "{}", "", 1)
	pagePattern = strings.Replace(pagePattern, "#", "", 1)

	pageNumber = strings.Replace(pageLocalId, pagePattern, "", 1)

	return
}