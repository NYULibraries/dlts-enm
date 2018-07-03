// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// RelationRelationtype represents a row from 'public.relation_relationtype'.
type RelationRelationtype struct {
	ID          int            `json:"id"`          // id
	Rtype       string         `json:"rtype"`       // rtype
	Description sql.NullString `json:"description"` // description
	RoleFrom    string         `json:"role_from"`   // role_from
	RoleTo      string         `json:"role_to"`     // role_to
	Symmetrical bool           `json:"symmetrical"` // symmetrical

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the RelationRelationtype exists in the database.
func (rr *RelationRelationtype) Exists() bool {
	return rr._exists
}

// Deleted provides information if the RelationRelationtype has been deleted from the database.
func (rr *RelationRelationtype) Deleted() bool {
	return rr._deleted
}

// Insert inserts the RelationRelationtype to the database.
func (rr *RelationRelationtype) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if rr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.relation_relationtype (` +
		`rtype, description, role_from, role_to, symmetrical` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical)
	err = db.QueryRow(sqlstr, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical).Scan(&rr.ID)
	if err != nil {
		return err
	}

	// set existence
	rr._exists = true

	return nil
}

// Update updates the RelationRelationtype in the database.
func (rr *RelationRelationtype) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if rr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.relation_relationtype SET (` +
		`rtype, description, role_from, role_to, symmetrical` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE id = $6`

	// run query
	XOLog(sqlstr, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical, rr.ID)
	_, err = db.Exec(sqlstr, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical, rr.ID)
	return err
}

// Save saves the RelationRelationtype to the database.
func (rr *RelationRelationtype) Save(db XODB) error {
	if rr.Exists() {
		return rr.Update(db)
	}

	return rr.Insert(db)
}

// Upsert performs an upsert for RelationRelationtype.
//
// NOTE: PostgreSQL 9.5+ only
func (rr *RelationRelationtype) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if rr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.relation_relationtype (` +
		`id, rtype, description, role_from, role_to, symmetrical` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, rtype, description, role_from, role_to, symmetrical` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.rtype, EXCLUDED.description, EXCLUDED.role_from, EXCLUDED.role_to, EXCLUDED.symmetrical` +
		`)`

	// run query
	XOLog(sqlstr, rr.ID, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical)
	_, err = db.Exec(sqlstr, rr.ID, rr.Rtype, rr.Description, rr.RoleFrom, rr.RoleTo, rr.Symmetrical)
	if err != nil {
		return err
	}

	// set existence
	rr._exists = true

	return nil
}

// Delete deletes the RelationRelationtype from the database.
func (rr *RelationRelationtype) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !rr._exists {
		return nil
	}

	// if deleted, bail
	if rr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.relation_relationtype WHERE id = $1`

	// run query
	XOLog(sqlstr, rr.ID)
	_, err = db.Exec(sqlstr, rr.ID)
	if err != nil {
		return err
	}

	// set deleted
	rr._deleted = true

	return nil
}

// RelationRelationtypeByID retrieves a row from 'public.relation_relationtype' as a RelationRelationtype.
//
// Generated from index 'relation_relationtype_pkey'.
func RelationRelationtypeByID(db XODB, id int) (*RelationRelationtype, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, rtype, description, role_from, role_to, symmetrical ` +
		`FROM public.relation_relationtype ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	rr := RelationRelationtype{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&rr.ID, &rr.Rtype, &rr.Description, &rr.RoleFrom, &rr.RoleTo, &rr.Symmetrical)
	if err != nil {
		return nil, err
	}

	return &rr, nil
}

// RelationRelationtypesByRtype retrieves a row from 'public.relation_relationtype' as a RelationRelationtype.
//
// Generated from index 'relation_relationtype_rtype_4542a47f_like'.
func RelationRelationtypesByRtype(db XODB, rtype string) ([]*RelationRelationtype, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, rtype, description, role_from, role_to, symmetrical ` +
		`FROM public.relation_relationtype ` +
		`WHERE rtype = $1`

	// run query
	XOLog(sqlstr, rtype)
	q, err := db.Query(sqlstr, rtype)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*RelationRelationtype{}
	for q.Next() {
		rr := RelationRelationtype{
			_exists: true,
		}

		// scan
		err = q.Scan(&rr.ID, &rr.Rtype, &rr.Description, &rr.RoleFrom, &rr.RoleTo, &rr.Symmetrical)
		if err != nil {
			return nil, err
		}

		res = append(res, &rr)
	}

	return res, nil
}

// RelationRelationtypeByRtypeRoleFromRoleTo retrieves a row from 'public.relation_relationtype' as a RelationRelationtype.
//
// Generated from index 'relation_relationtype_rtype_65db05ac_uniq'.
func RelationRelationtypeByRtypeRoleFromRoleTo(db XODB, rtype string, roleFrom string, roleTo string) (*RelationRelationtype, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, rtype, description, role_from, role_to, symmetrical ` +
		`FROM public.relation_relationtype ` +
		`WHERE rtype = $1 AND role_from = $2 AND role_to = $3`

	// run query
	XOLog(sqlstr, rtype, roleFrom, roleTo)
	rr := RelationRelationtype{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, rtype, roleFrom, roleTo).Scan(&rr.ID, &rr.Rtype, &rr.Description, &rr.RoleFrom, &rr.RoleTo, &rr.Symmetrical)
	if err != nil {
		return nil, err
	}

	return &rr, nil
}

// RelationRelationtypeByRtype retrieves a row from 'public.relation_relationtype' as a RelationRelationtype.
//
// Generated from index 'relation_relationtype_rtype_key'.
func RelationRelationtypeByRtype(db XODB, rtype string) (*RelationRelationtype, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, rtype, description, role_from, role_to, symmetrical ` +
		`FROM public.relation_relationtype ` +
		`WHERE rtype = $1`

	// run query
	XOLog(sqlstr, rtype)
	rr := RelationRelationtype{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, rtype).Scan(&rr.ID, &rr.Rtype, &rr.Description, &rr.RoleFrom, &rr.RoleTo, &rr.Symmetrical)
	if err != nil {
		return nil, err
	}

	return &rr, nil
}
