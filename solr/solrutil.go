package solrutil

import (
	"fmt"

	"github.com/nyulibraries/dlts-enm/db/models"
	solr "github.com/rtt/Go-Solr"
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
	pageNumberForDisplay := page.PageLocalid

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
				"pageNumberForDisplay": pageNumberForDisplay,
				"pageSequenceNumber": page.PageSequence,
				"pageText": page.PageText,
				"preferredTopicName": "TBD",
				"topicNames": "TBD",
			},
		},
	}

	resp, err := conn.Update(doc, true)
	if err != nil {
		return err
	}

	fmt.Println("%v", resp)

	return nil
}
