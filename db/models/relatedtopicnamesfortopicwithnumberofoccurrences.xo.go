// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

// RelatedTopicNamesForTopicWithNumberOfOccurrences represents a row from '[custom related_topic_names_for_topic_with_number_of_occurrences]'.
type RelatedTopicNamesForTopicWithNumberOfOccurrences struct {
	Topic2ID            int    // topic2_id
	DisplayNameDoNotUse string // display_name_do_not_use
	NumberOfOccurrences int64  // number_of_occurrences
}

// RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id runs a custom query, returning results as RelatedTopicNamesForTopicWithNumberOfOccurrences.
func RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id(db XODB, topic_id int) ([]*RelatedTopicNamesForTopicWithNumberOfOccurrences, error) {
	var err error

	// sql query
	const sqlstr = `SELECT trs.topic2_id, t.display_name_do_not_use, COUNT( * ) AS number_of_occurrences ` +
		` ` +
		`FROM topic_relations_simple trs INNER JOIN topics t ON trs.topic2_id = t.tct_id ` +
		`INNER JOIN occurrences o ON o.topic_id = trs.topic2_id ` +
		` ` +
		`WHERE topic1_id = ? ` +
		` ` +
		`GROUP BY trs.topic2_id, t.display_name_do_not_use ` +
		` ` +
		`ORDER BY t.display_name_do_not_use`

	// run query
	XOLog(sqlstr, topic_id)
	q, err := db.Query(sqlstr, topic_id)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*RelatedTopicNamesForTopicWithNumberOfOccurrences{}
	for q.Next() {
		rtnftwnoo := RelatedTopicNamesForTopicWithNumberOfOccurrences{}

		// scan
		err = q.Scan(&rtnftwnoo.Topic2ID, &rtnftwnoo.DisplayNameDoNotUse, &rtnftwnoo.NumberOfOccurrences)
		if err != nil {
			return nil, err
		}

		res = append(res, &rtnftwnoo)
	}

	return res, nil
}