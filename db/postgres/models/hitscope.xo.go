// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// HitScope represents a row from 'public.hit_scope'.
type HitScope struct {
	ID          int            `json:"id"`          // id
	Scope       string         `json:"scope"`       // scope
	Description sql.NullString `json:"description"` // description

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the HitScope exists in the database.
func (hs *HitScope) Exists() bool {
	return hs._exists
}

// Deleted provides information if the HitScope has been deleted from the database.
func (hs *HitScope) Deleted() bool {
	return hs._deleted
}

// Insert inserts the HitScope to the database.
func (hs *HitScope) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if hs._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.hit_scope (` +
		`scope, description` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, hs.Scope, hs.Description)
	err = db.QueryRow(sqlstr, hs.Scope, hs.Description).Scan(&hs.ID)
	if err != nil {
		return err
	}

	// set existence
	hs._exists = true

	return nil
}

// Update updates the HitScope in the database.
func (hs *HitScope) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hs._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if hs._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.hit_scope SET (` +
		`scope, description` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE id = $3`

	// run query
	XOLog(sqlstr, hs.Scope, hs.Description, hs.ID)
	_, err = db.Exec(sqlstr, hs.Scope, hs.Description, hs.ID)
	return err
}

// Save saves the HitScope to the database.
func (hs *HitScope) Save(db XODB) error {
	if hs.Exists() {
		return hs.Update(db)
	}

	return hs.Insert(db)
}

// Upsert performs an upsert for HitScope.
//
// NOTE: PostgreSQL 9.5+ only
func (hs *HitScope) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if hs._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.hit_scope (` +
		`id, scope, description` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, scope, description` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.scope, EXCLUDED.description` +
		`)`

	// run query
	XOLog(sqlstr, hs.ID, hs.Scope, hs.Description)
	_, err = db.Exec(sqlstr, hs.ID, hs.Scope, hs.Description)
	if err != nil {
		return err
	}

	// set existence
	hs._exists = true

	return nil
}

// Delete deletes the HitScope from the database.
func (hs *HitScope) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !hs._exists {
		return nil
	}

	// if deleted, bail
	if hs._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.hit_scope WHERE id = $1`

	// run query
	XOLog(sqlstr, hs.ID)
	_, err = db.Exec(sqlstr, hs.ID)
	if err != nil {
		return err
	}

	// set deleted
	hs._deleted = true

	return nil
}

// HitScopeByID retrieves a row from 'public.hit_scope' as a HitScope.
//
// Generated from index 'hit_scope_pkey'.
func HitScopeByID(db XODB, id int) (*HitScope, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, scope, description ` +
		`FROM public.hit_scope ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	hs := HitScope{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&hs.ID, &hs.Scope, &hs.Description)
	if err != nil {
		return nil, err
	}

	return &hs, nil
}

// HitScopesByScope retrieves a row from 'public.hit_scope' as a HitScope.
//
// Generated from index 'hit_scope_scope_e1a2f469_like'.
func HitScopesByScope(db XODB, scope string) ([]*HitScope, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, scope, description ` +
		`FROM public.hit_scope ` +
		`WHERE scope = $1`

	// run query
	XOLog(sqlstr, scope)
	q, err := db.Query(sqlstr, scope)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*HitScope{}
	for q.Next() {
		hs := HitScope{
			_exists: true,
		}

		// scan
		err = q.Scan(&hs.ID, &hs.Scope, &hs.Description)
		if err != nil {
			return nil, err
		}

		res = append(res, &hs)
	}

	return res, nil
}

// HitScopeByScope retrieves a row from 'public.hit_scope' as a HitScope.
//
// Generated from index 'hit_scope_scope_key'.
func HitScopeByScope(db XODB, scope string) (*HitScope, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, scope, description ` +
		`FROM public.hit_scope ` +
		`WHERE scope = $1`

	// run query
	XOLog(sqlstr, scope)
	hs := HitScope{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, scope).Scan(&hs.ID, &hs.Scope, &hs.Description)
	if err != nil {
		return nil, err
	}

	return &hs, nil
}
