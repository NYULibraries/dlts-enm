// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// EpubsForTopicWithNumberOfMatchedPages represents a row from '[custom epubs_for_topic_with_number_of_matched_pages]'.
type EpubsForTopicWithNumberOfMatchedPages struct {
	Title               string // title
	Author              string // author
	Publisher           string // publisher
	Isbn                string // isbn
	NumberOfOccurrences int64  // number_of_occurrences
}

// EpubsForTopicWithNumberOfMatchedPagesByTopic_id runs a custom query, returning results as EpubsForTopicWithNumberOfMatchedPages.
func EpubsForTopicWithNumberOfMatchedPagesByTopic_id(db XODB, topic_id int) ([]*EpubsForTopicWithNumberOfMatchedPages, error) {
	var err error

	// sql query
	const sqlstr = `SELECT od.title, od.author, oee.publisher, LEFT(RIGHT(oee.source, 18), 13) AS isbn, ` +
		`COUNT(oo.*) number_of_occurrences ` +
		` ` +
		`FROM occurrence_document od ` +
		`INNER JOIN otx_epub_epub oee ON oee.document_ptr_id = od.id ` +
		`INNER JOIN occurrence_location ol on od.id = ol.document_id ` +
		`INNER JOIN occurrence_occurrence oo ON oo.location_id = ol.id ` +
		`INNER JOIN hit_basket hb ON hb.id = oo.basket_id ` +
		` ` +
		`WHERE hb.id = $1 ` +
		` ` +
		`GROUP BY od.title, od.author, oee.publisher, isbn ` +
		` ` +
		`ORDER BY number_of_occurrences DESC`

	// run query
	XOLog(sqlstr, topic_id)
	q, err := db.Query(sqlstr, topic_id)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*EpubsForTopicWithNumberOfMatchedPages{}
	for q.Next() {
		eftwnomp := EpubsForTopicWithNumberOfMatchedPages{}

		// scan
		err = q.Scan(&eftwnomp.Title, &eftwnomp.Author, &eftwnomp.Publisher, &eftwnomp.Isbn, &eftwnomp.NumberOfOccurrences)
		if err != nil {
			return nil, err
		}

		res = append(res, &eftwnomp)
	}

	return res, nil
}