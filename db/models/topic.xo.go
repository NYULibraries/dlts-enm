// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

// Topic represents a row from 'enm.topics'.
type Topic struct {
	TctID int `json:"tct_id"` // tct_id
}

// TopicByTctID retrieves a row from 'enm.topics' as a Topic.
//
// Generated from index 'topics_tct_id_pkey'.
func TopicByTctID(db XODB, tctID int) (*Topic, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id ` +
		`FROM enm.topics ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	t := Topic{}

	err = db.QueryRow(sqlstr, tctID).Scan(&t.TctID)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
