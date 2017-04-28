// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// Topic represents a row from 'enm.topics'.
type Topic struct {
	TctID               int    `json:"tct_id"`                  // tct_id
	DisplayNameDoNotUse string `json:"display_name_do_not_use"` // display_name_do_not_use

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Topic exists in the database.
func (t *Topic) Exists() bool {
	return t._exists
}

// Deleted provides information if the Topic has been deleted from the database.
func (t *Topic) Deleted() bool {
	return t._deleted
}

// Insert inserts the Topic to the database.
func (t *Topic) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if t._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO enm.topics (` +
		`tct_id, display_name_do_not_use` +
		`) VALUES (` +
		`?, ?` +
		`)`

	// run query
	XOLog(sqlstr, t.TctID, t.DisplayNameDoNotUse)
	_, err = db.Exec(sqlstr, t.TctID, t.DisplayNameDoNotUse)
	if err != nil {
		return err
	}

	// set existence
	t._exists = true

	return nil
}

// Update updates the Topic in the database.
func (t *Topic) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if t._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE enm.topics SET ` +
		`display_name_do_not_use = ?` +
		` WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, t.DisplayNameDoNotUse, t.TctID)
	_, err = db.Exec(sqlstr, t.DisplayNameDoNotUse, t.TctID)
	return err
}

// Save saves the Topic to the database.
func (t *Topic) Save(db XODB) error {
	if t.Exists() {
		return t.Update(db)
	}

	return t.Insert(db)
}

// Delete deletes the Topic from the database.
func (t *Topic) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !t._exists {
		return nil
	}

	// if deleted, bail
	if t._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM enm.topics WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, t.TctID)
	_, err = db.Exec(sqlstr, t.TctID)
	if err != nil {
		return err
	}

	// set deleted
	t._deleted = true

	return nil
}

// TopicByTctID retrieves a row from 'enm.topics' as a Topic.
//
// Generated from index 'topics_tct_id_pkey'.
func TopicByTctID(db XODB, tctID int) (*Topic, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, display_name_do_not_use ` +
		`FROM enm.topics ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	t := Topic{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, tctID).Scan(&t.TctID, &t.DisplayNameDoNotUse)
	if err != nil {
		return nil, err
	}

	return &t, nil
}
