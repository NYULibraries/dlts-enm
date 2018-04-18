// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
)

// OccurrenceOccurrence represents a row from 'public.occurrence_occurrence'.
type OccurrenceOccurrence struct {
	ID           int            `json:"id"`             // id
	HitInContent sql.NullString `json:"hit_in_content"` // hit_in_content
	BasketID     sql.NullInt64  `json:"basket_id"`      // basket_id
	LocationID   sql.NullInt64  `json:"location_id"`    // location_id
	Order        int            `json:"_order"`         // _order

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OccurrenceOccurrence exists in the database.
func (oo *OccurrenceOccurrence) Exists() bool {
	return oo._exists
}

// Deleted provides information if the OccurrenceOccurrence has been deleted from the database.
func (oo *OccurrenceOccurrence) Deleted() bool {
	return oo._deleted
}

// Insert inserts the OccurrenceOccurrence to the database.
func (oo *OccurrenceOccurrence) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if oo._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.occurrence_occurrence (` +
		`hit_in_content, basket_id, location_id, _order` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order)
	err = db.QueryRow(sqlstr, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order).Scan(&oo.ID)
	if err != nil {
		return err
	}

	// set existence
	oo._exists = true

	return nil
}

// Update updates the OccurrenceOccurrence in the database.
func (oo *OccurrenceOccurrence) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oo._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if oo._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.occurrence_occurrence SET (` +
		`hit_in_content, basket_id, location_id, _order` +
		`) = ( ` +
		`$1, $2, $3, $4` +
		`) WHERE id = $5`

	// run query
	XOLog(sqlstr, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order, oo.ID)
	_, err = db.Exec(sqlstr, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order, oo.ID)
	return err
}

// Save saves the OccurrenceOccurrence to the database.
func (oo *OccurrenceOccurrence) Save(db XODB) error {
	if oo.Exists() {
		return oo.Update(db)
	}

	return oo.Insert(db)
}

// Upsert performs an upsert for OccurrenceOccurrence.
//
// NOTE: PostgreSQL 9.5+ only
func (oo *OccurrenceOccurrence) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if oo._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.occurrence_occurrence (` +
		`id, hit_in_content, basket_id, location_id, _order` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, hit_in_content, basket_id, location_id, _order` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.hit_in_content, EXCLUDED.basket_id, EXCLUDED.location_id, EXCLUDED._order` +
		`)`

	// run query
	XOLog(sqlstr, oo.ID, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order)
	_, err = db.Exec(sqlstr, oo.ID, oo.HitInContent, oo.BasketID, oo.LocationID, oo.Order)
	if err != nil {
		return err
	}

	// set existence
	oo._exists = true

	return nil
}

// Delete deletes the OccurrenceOccurrence from the database.
func (oo *OccurrenceOccurrence) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oo._exists {
		return nil
	}

	// if deleted, bail
	if oo._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.occurrence_occurrence WHERE id = $1`

	// run query
	XOLog(sqlstr, oo.ID)
	_, err = db.Exec(sqlstr, oo.ID)
	if err != nil {
		return err
	}

	// set deleted
	oo._deleted = true

	return nil
}

// OccurrenceLocation returns the OccurrenceLocation associated with the OccurrenceOccurrence's LocationID (location_id).
//
// Generated from foreign key 'occurrence_occur_location_id_0469bb62_fk_occurrence_location_id'.
func (oo *OccurrenceOccurrence) OccurrenceLocation(db XODB) (*OccurrenceLocation, error) {
	return OccurrenceLocationByID(db, int(oo.LocationID.Int64))
}

// HitBasket returns the HitBasket associated with the OccurrenceOccurrence's BasketID (basket_id).
//
// Generated from foreign key 'occurrence_occurrence_basket_id_3a5fce9a_fk_hit_basket_id'.
func (oo *OccurrenceOccurrence) HitBasket(db XODB) (*HitBasket, error) {
	return HitBasketByID(db, int(oo.BasketID.Int64))
}

// OccurrenceOccurrencesByBasketID retrieves a row from 'public.occurrence_occurrence' as a OccurrenceOccurrence.
//
// Generated from index 'occurrence_occurrence_afdeaea9'.
func OccurrenceOccurrencesByBasketID(db XODB, basketID sql.NullInt64) ([]*OccurrenceOccurrence, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, hit_in_content, basket_id, location_id, _order ` +
		`FROM public.occurrence_occurrence ` +
		`WHERE basket_id = $1`

	// run query
	XOLog(sqlstr, basketID)
	q, err := db.Query(sqlstr, basketID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*OccurrenceOccurrence{}
	for q.Next() {
		oo := OccurrenceOccurrence{
			_exists: true,
		}

		// scan
		err = q.Scan(&oo.ID, &oo.HitInContent, &oo.BasketID, &oo.LocationID, &oo.Order)
		if err != nil {
			return nil, err
		}

		res = append(res, &oo)
	}

	return res, nil
}

// OccurrenceOccurrencesByLocationID retrieves a row from 'public.occurrence_occurrence' as a OccurrenceOccurrence.
//
// Generated from index 'occurrence_occurrence_b171ba63'.
func OccurrenceOccurrencesByLocationID(db XODB, locationID sql.NullInt64) ([]*OccurrenceOccurrence, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, hit_in_content, basket_id, location_id, _order ` +
		`FROM public.occurrence_occurrence ` +
		`WHERE location_id = $1`

	// run query
	XOLog(sqlstr, locationID)
	q, err := db.Query(sqlstr, locationID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*OccurrenceOccurrence{}
	for q.Next() {
		oo := OccurrenceOccurrence{
			_exists: true,
		}

		// scan
		err = q.Scan(&oo.ID, &oo.HitInContent, &oo.BasketID, &oo.LocationID, &oo.Order)
		if err != nil {
			return nil, err
		}

		res = append(res, &oo)
	}

	return res, nil
}

// OccurrenceOccurrenceByID retrieves a row from 'public.occurrence_occurrence' as a OccurrenceOccurrence.
//
// Generated from index 'occurrence_occurrence_pkey'.
func OccurrenceOccurrenceByID(db XODB, id int) (*OccurrenceOccurrence, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, hit_in_content, basket_id, location_id, _order ` +
		`FROM public.occurrence_occurrence ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	oo := OccurrenceOccurrence{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&oo.ID, &oo.HitInContent, &oo.BasketID, &oo.LocationID, &oo.Order)
	if err != nil {
		return nil, err
	}

	return &oo, nil
}
