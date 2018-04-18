// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// OtxWeblinkWeblink represents a row from 'public.otx_weblink_weblink'.
type OtxWeblinkWeblink struct {
	ID      int            `json:"id"`      // id
	URL     sql.NullString `json:"url"`     // url
	Content sql.NullString `json:"content"` // content

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OtxWeblinkWeblink exists in the database.
func (oww *OtxWeblinkWeblink) Exists() bool {
	return oww._exists
}

// Deleted provides information if the OtxWeblinkWeblink has been deleted from the database.
func (oww *OtxWeblinkWeblink) Deleted() bool {
	return oww._deleted
}

// Insert inserts the OtxWeblinkWeblink to the database.
func (oww *OtxWeblinkWeblink) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if oww._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.otx_weblink_weblink (` +
		`url, content` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, oww.URL, oww.Content)
	err = db.QueryRow(sqlstr, oww.URL, oww.Content).Scan(&oww.ID)
	if err != nil {
		return err
	}

	// set existence
	oww._exists = true

	return nil
}

// Update updates the OtxWeblinkWeblink in the database.
func (oww *OtxWeblinkWeblink) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oww._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if oww._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.otx_weblink_weblink SET (` +
		`url, content` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE id = $3`

	// run query
	XOLog(sqlstr, oww.URL, oww.Content, oww.ID)
	_, err = db.Exec(sqlstr, oww.URL, oww.Content, oww.ID)
	return err
}

// Save saves the OtxWeblinkWeblink to the database.
func (oww *OtxWeblinkWeblink) Save(db XODB) error {
	if oww.Exists() {
		return oww.Update(db)
	}

	return oww.Insert(db)
}

// Upsert performs an upsert for OtxWeblinkWeblink.
//
// NOTE: PostgreSQL 9.5+ only
func (oww *OtxWeblinkWeblink) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if oww._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.otx_weblink_weblink (` +
		`id, url, content` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, url, content` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.url, EXCLUDED.content` +
		`)`

	// run query
	XOLog(sqlstr, oww.ID, oww.URL, oww.Content)
	_, err = db.Exec(sqlstr, oww.ID, oww.URL, oww.Content)
	if err != nil {
		return err
	}

	// set existence
	oww._exists = true

	return nil
}

// Delete deletes the OtxWeblinkWeblink from the database.
func (oww *OtxWeblinkWeblink) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oww._exists {
		return nil
	}

	// if deleted, bail
	if oww._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.otx_weblink_weblink WHERE id = $1`

	// run query
	XOLog(sqlstr, oww.ID)
	_, err = db.Exec(sqlstr, oww.ID)
	if err != nil {
		return err
	}

	// set deleted
	oww._deleted = true

	return nil
}

// OtxWeblinkWeblinkByID retrieves a row from 'public.otx_weblink_weblink' as a OtxWeblinkWeblink.
//
// Generated from index 'otx_weblink_weblink_pkey'.
func OtxWeblinkWeblinkByID(db XODB, id int) (*OtxWeblinkWeblink, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, url, content ` +
		`FROM public.otx_weblink_weblink ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	oww := OtxWeblinkWeblink{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&oww.ID, &oww.URL, &oww.Content)
	if err != nil {
		return nil, err
	}

	return &oww, nil
}
