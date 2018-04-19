// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// Name represents a row from 'enm.names'.
type Name struct {
	TctID     int    `json:"tct_id"`    // tct_id
	TopicID   int    `json:"topic_id"`  // topic_id
	Name      string `json:"name"`      // name
	ScopeID   int    `json:"scope_id"`  // scope_id
	Bypass    bool   `json:"bypass"`    // bypass
	Hidden    bool   `json:"hidden"`    // hidden
	Preferred bool   `json:"preferred"` // preferred

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Name exists in the database.
func (n *Name) Exists() bool {
	return n._exists
}

// Deleted provides information if the Name has been deleted from the database.
func (n *Name) Deleted() bool {
	return n._deleted
}

// Insert inserts the Name to the database.
func (n *Name) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if n._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO enm.names (` +
		`tct_id, topic_id, name, scope_id, bypass, hidden, preferred` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, n.TctID, n.TopicID, n.Name, n.ScopeID, n.Bypass, n.Hidden, n.Preferred)
	_, err = db.Exec(sqlstr, n.TctID, n.TopicID, n.Name, n.ScopeID, n.Bypass, n.Hidden, n.Preferred)
	if err != nil {
		return err
	}

	// set existence
	n._exists = true

	return nil
}

// Update updates the Name in the database.
func (n *Name) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !n._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if n._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE enm.names SET ` +
		`topic_id = ?, name = ?, scope_id = ?, bypass = ?, hidden = ?, preferred = ?` +
		` WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, n.TopicID, n.Name, n.ScopeID, n.Bypass, n.Hidden, n.Preferred, n.TctID)
	_, err = db.Exec(sqlstr, n.TopicID, n.Name, n.ScopeID, n.Bypass, n.Hidden, n.Preferred, n.TctID)
	return err
}

// Save saves the Name to the database.
func (n *Name) Save(db XODB) error {
	if n.Exists() {
		return n.Update(db)
	}

	return n.Insert(db)
}

// Delete deletes the Name from the database.
func (n *Name) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !n._exists {
		return nil
	}

	// if deleted, bail
	if n._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM enm.names WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, n.TctID)
	_, err = db.Exec(sqlstr, n.TctID)
	if err != nil {
		return err
	}

	// set deleted
	n._deleted = true

	return nil
}

// Scope returns the Scope associated with the Name's ScopeID (scope_id).
//
// Generated from foreign key 'fk__names__scopes'.
func (n *Name) Scope(db XODB) (*Scope, error) {
	return ScopeByTctID(db, n.ScopeID)
}

// Topic returns the Topic associated with the Name's TopicID (topic_id).
//
// Generated from foreign key 'fk__names__topics'.
func (n *Name) Topic(db XODB) (*Topic, error) {
	return TopicByTctID(db, n.TopicID)
}

// NamesByName retrieves a row from 'enm.names' as a Name.
//
// Generated from index 'name'.
func NamesByName(db XODB, name string) ([]*Name, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, topic_id, name, scope_id, bypass, hidden, preferred ` +
		`FROM enm.names ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	q, err := db.Query(sqlstr, name)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Name{}
	for q.Next() {
		n := Name{
			_exists: true,
		}

		// scan
		err = q.Scan(&n.TctID, &n.TopicID, &n.Name, &n.ScopeID, &n.Bypass, &n.Hidden, &n.Preferred)
		if err != nil {
			return nil, err
		}

		res = append(res, &n)
	}

	return res, nil
}

// NameByTctID retrieves a row from 'enm.names' as a Name.
//
// Generated from index 'names_tct_id_pkey'.
func NameByTctID(db XODB, tctID int) (*Name, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, topic_id, name, scope_id, bypass, hidden, preferred ` +
		`FROM enm.names ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	n := Name{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, tctID).Scan(&n.TctID, &n.TopicID, &n.Name, &n.ScopeID, &n.Bypass, &n.Hidden, &n.Preferred)
	if err != nil {
		return nil, err
	}

	return &n, nil
}

// NamesByScopeID retrieves a row from 'enm.names' as a Name.
//
// Generated from index 'scope_id'.
func NamesByScopeID(db XODB, scopeID int) ([]*Name, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, topic_id, name, scope_id, bypass, hidden, preferred ` +
		`FROM enm.names ` +
		`WHERE scope_id = ?`

	// run query
	XOLog(sqlstr, scopeID)
	q, err := db.Query(sqlstr, scopeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Name{}
	for q.Next() {
		n := Name{
			_exists: true,
		}

		// scan
		err = q.Scan(&n.TctID, &n.TopicID, &n.Name, &n.ScopeID, &n.Bypass, &n.Hidden, &n.Preferred)
		if err != nil {
			return nil, err
		}

		res = append(res, &n)
	}

	return res, nil
}

// NamesByTopicID retrieves a row from 'enm.names' as a Name.
//
// Generated from index 'topic_id'.
func NamesByTopicID(db XODB, topicID int) ([]*Name, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, topic_id, name, scope_id, bypass, hidden, preferred ` +
		`FROM enm.names ` +
		`WHERE topic_id = ?`

	// run query
	XOLog(sqlstr, topicID)
	q, err := db.Query(sqlstr, topicID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Name{}
	for q.Next() {
		n := Name{
			_exists: true,
		}

		// scan
		err = q.Scan(&n.TctID, &n.TopicID, &n.Name, &n.ScopeID, &n.Bypass, &n.Hidden, &n.Preferred)
		if err != nil {
			return nil, err
		}

		res = append(res, &n)
	}

	return res, nil
}