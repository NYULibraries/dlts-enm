// Package models contains the types for schema 'public'.
package models

// Code generated by xo. DO NOT EDIT.

// TopicAlternateName represents a row from '[custom topic_alternate_name]'.
type TopicAlternateName struct {
	TctID       int    // tct_id
	DisplayName string // display_name
	Name        string // name
}

// GetTopicAlternateNames runs a custom query, returning results as TopicAlternateName.
func GetTopicAlternateNames(db XODB) ([]*TopicAlternateName, error) {
	var err error

	// sql query
	const sqlstr = `SELECT hb.id AS tct_id, hb.display_name, hh.name ` +
		`FROM hit_basket hb INNER JOIN hit_hit hh ON hb.id = hh.basket_id ` +
		`ORDER BY display_name, name`

	// run query
	XOLog(sqlstr)
	q, err := db.Query(sqlstr)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*TopicAlternateName{}
	for q.Next() {
		tan := TopicAlternateName{}

		// scan
		err = q.Scan(&tan.TctID, &tan.DisplayName, &tan.Name)
		if err != nil {
			return nil, err
		}

		res = append(res, &tan)
	}

	return res, nil
}
