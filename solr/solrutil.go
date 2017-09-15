package solrutil

import (
	"log"
	"strings"

	"github.com/nyulibraries/dlts-enm/db/models"
	solr "github.com/rtt/Go-Solr"
	"github.com/nyulibraries/dlts-enm/db"
)

var Port int
var Server string
var conn *solr.Connection

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
	var topicNamesForDisplay string

	pageTopicNames := db.GetPageTopicNamesByPageId(page.ID)

	for _, pageTopic := range pageTopicNames {
		topicNames = append(topicNames, pageTopic.PreferredTopicName)
	}

	topicNamesForDisplay = "TBD"

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
				"topicNamesForDisplay": topicNamesForDisplay,
			},
		},
	}

	_, err := conn.Update(doc, true)
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