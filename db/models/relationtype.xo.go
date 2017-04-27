// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// RelationType represents a row from 'enm.relation_type'.
type RelationType struct {
	TctID       int    `json:"tct_id"`      // tct_id
	Rtype       string `json:"rtype"`       // rtype
	RoleFrom    string `json:"role_from"`   // role_from
	RoleTo      string `json:"role_to"`     // role_to
	Symmetrical bool   `json:"symmetrical"` // symmetrical

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the RelationType exists in the database.
func (rt *RelationType) Exists() bool {
	return rt._exists
}

// Deleted provides information if the RelationType has been deleted from the database.
func (rt *RelationType) Deleted() bool {
	return rt._deleted
}

// Insert inserts the RelationType to the database.
func (rt *RelationType) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if rt._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO enm.relation_type (` +
		`tct_id, rtype, role_from, role_to, symmetrical` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, rt.TctID, rt.Rtype, rt.RoleFrom, rt.RoleTo, rt.Symmetrical)
	_, err = db.Exec(sqlstr, rt.TctID, rt.Rtype, rt.RoleFrom, rt.RoleTo, rt.Symmetrical)
	if err != nil {
		return err
	}

	// set existence
	rt._exists = true

	return nil
}

// Update updates the RelationType in the database.
func (rt *RelationType) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rt._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if rt._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE enm.relation_type SET ` +
		`rtype = ?, role_from = ?, role_to = ?, symmetrical = ?` +
		` WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, rt.Rtype, rt.RoleFrom, rt.RoleTo, rt.Symmetrical, rt.TctID)
	_, err = db.Exec(sqlstr, rt.Rtype, rt.RoleFrom, rt.RoleTo, rt.Symmetrical, rt.TctID)
	return err
}

// Save saves the RelationType to the database.
func (rt *RelationType) Save(db XODB) error {
	if rt.Exists() {
		return rt.Update(db)
	}

	return rt.Insert(db)
}

// Delete deletes the RelationType from the database.
func (rt *RelationType) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rt._exists {
		return nil
	}

	// if deleted, bail
	if rt._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM enm.relation_type WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, rt.TctID)
	_, err = db.Exec(sqlstr, rt.TctID)
	if err != nil {
		return err
	}

	// set deleted
	rt._deleted = true

	return nil
}

// RelationTypeByTctID retrieves a row from 'enm.relation_type' as a RelationType.
//
// Generated from index 'relation_type_tct_id_pkey'.
func RelationTypeByTctID(db XODB, tctID int) (*RelationType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, rtype, role_from, role_to, symmetrical ` +
		`FROM enm.relation_type ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	rt := RelationType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, tctID).Scan(&rt.TctID, &rt.Rtype, &rt.RoleFrom, &rt.RoleTo, &rt.Symmetrical)
	if err != nil {
		return nil, err
	}

	return &rt, nil
}
