// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// RelatedTopicNamesForTopicWithNumberOfOccurrences represents a row from '[custom related_topic_names_for_topic_with_number_of_occurrences]'.
type RelatedTopicNamesForTopicWithNumberOfOccurrences struct {
	ID                  int    // id
	DisplayName         string // display_name
	NumberOfOccurrences int64  // number_of_occurrences
}

// RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id_1Topic_id_2Topic_id_3Topic_id_4 runs a custom query, returning results as RelatedTopicNamesForTopicWithNumberOfOccurrences.
func RelatedTopicNamesForTopicWithNumberOfOccurrencesByTopic_id_1Topic_id_2Topic_id_3Topic_id_4(db XODB, topic_id_1 int, topic_id_2 int, topic_id_3 int, topic_id_4 int) ([]*RelatedTopicNamesForTopicWithNumberOfOccurrences, error) {
	var err error

	// sql query
	const sqlstr = `SELECT hb.id, hb.display_name, COUNT(o.*) AS number_of_occurrences ` +
		`FROM ` +
		`( ` +
		`SELECT rrb.destination_id AS related_topic_id ` +
		`FROM relation_relatedbasket rrb ` +
		`WHERE rrb.source_id = $1 ` +
		`AND rrb.forbidden = FALSE ` +
		` ` +
		`UNION ` +
		` ` +
		`SELECT rrb.source_id AS related_topic_id ` +
		`FROM relation_relatedbasket rrb ` +
		`WHERE rrb.destination_id = $2 ` +
		`AND rrb.forbidden = FALSE ` +
		`) AS r INNER JOIN hit_basket hb ON r.related_topic_id = hb.id ` +
		`INNER JOIN occurrence_occurrence o on hb.id = o.basket_id ` +
		`GROUP BY hb.id, hb.display_name ` +
		` ` +
		`UNION ` +
		` ` +
		`SELECT hb.id, hb.display_name, 0 AS number_of_occurrences ` +
		`FROM ` +
		`( ` +
		`SELECT rrb.destination_id AS related_topic_id ` +
		`FROM relation_relatedbasket rrb ` +
		`WHERE rrb.source_id = $3 ` +
		`AND rrb.forbidden = FALSE ` +
		` ` +
		`UNION ` +
		` ` +
		`SELECT rrb.source_id AS related_topic_id ` +
		`FROM relation_relatedbasket rrb ` +
		`WHERE rrb.destination_id = $4 ` +
		`AND rrb.forbidden = FALSE ` +
		`) AS r INNER JOIN hit_basket hb ON r.related_topic_id = hb.id ` +
		`LEFT OUTER JOIN occurrence_occurrence o on hb.id = o.basket_id ` +
		`WHERE o.basket_id IS NULL ` +
		`GROUP BY hb.id, hb.display_name ` +
		` ` +
		`ORDER BY display_name`

	// run query
	XOLog(sqlstr, topic_id_1, topic_id_2, topic_id_3, topic_id_4)
	q, err := db.Query(sqlstr, topic_id_1, topic_id_2, topic_id_3, topic_id_4)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*RelatedTopicNamesForTopicWithNumberOfOccurrences{}
	for q.Next() {
		rtnftwnoo := RelatedTopicNamesForTopicWithNumberOfOccurrences{}

		// scan
		err = q.Scan(&rtnftwnoo.ID, &rtnftwnoo.DisplayName, &rtnftwnoo.NumberOfOccurrences)
		if err != nil {
			return nil, err
		}

		res = append(res, &rtnftwnoo)
	}

	return res, nil
}
