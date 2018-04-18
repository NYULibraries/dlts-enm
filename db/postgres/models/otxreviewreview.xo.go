// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"

	"github.com/lib/pq"
)

// OtxReviewReview represents a row from 'public.otx_review_review'.
type OtxReviewReview struct {
	ID         int           `json:"id"`          // id
	Time       pq.NullTime   `json:"time"`        // time
	Reviewed   bool          `json:"reviewed"`    // reviewed
	BasketID   int           `json:"basket_id"`   // basket_id
	ReviewerID sql.NullInt64 `json:"reviewer_id"` // reviewer_id
	Changed    bool          `json:"changed"`     // changed

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OtxReviewReview exists in the database.
func (orr *OtxReviewReview) Exists() bool {
	return orr._exists
}

// Deleted provides information if the OtxReviewReview has been deleted from the database.
func (orr *OtxReviewReview) Deleted() bool {
	return orr._deleted
}

// Insert inserts the OtxReviewReview to the database.
func (orr *OtxReviewReview) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if orr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.otx_review_review (` +
		`time, reviewed, basket_id, reviewer_id, changed` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed)
	err = db.QueryRow(sqlstr, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed).Scan(&orr.ID)
	if err != nil {
		return err
	}

	// set existence
	orr._exists = true

	return nil
}

// Update updates the OtxReviewReview in the database.
func (orr *OtxReviewReview) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !orr._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if orr._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.otx_review_review SET (` +
		`time, reviewed, basket_id, reviewer_id, changed` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE id = $6`

	// run query
	XOLog(sqlstr, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed, orr.ID)
	_, err = db.Exec(sqlstr, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed, orr.ID)
	return err
}

// Save saves the OtxReviewReview to the database.
func (orr *OtxReviewReview) Save(db XODB) error {
	if orr.Exists() {
		return orr.Update(db)
	}

	return orr.Insert(db)
}

// Upsert performs an upsert for OtxReviewReview.
//
// NOTE: PostgreSQL 9.5+ only
func (orr *OtxReviewReview) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if orr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.otx_review_review (` +
		`id, time, reviewed, basket_id, reviewer_id, changed` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, time, reviewed, basket_id, reviewer_id, changed` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.time, EXCLUDED.reviewed, EXCLUDED.basket_id, EXCLUDED.reviewer_id, EXCLUDED.changed` +
		`)`

	// run query
	XOLog(sqlstr, orr.ID, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed)
	_, err = db.Exec(sqlstr, orr.ID, orr.Time, orr.Reviewed, orr.BasketID, orr.ReviewerID, orr.Changed)
	if err != nil {
		return err
	}

	// set existence
	orr._exists = true

	return nil
}

// Delete deletes the OtxReviewReview from the database.
func (orr *OtxReviewReview) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !orr._exists {
		return nil
	}

	// if deleted, bail
	if orr._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.otx_review_review WHERE id = $1`

	// run query
	XOLog(sqlstr, orr.ID)
	_, err = db.Exec(sqlstr, orr.ID)
	if err != nil {
		return err
	}

	// set deleted
	orr._deleted = true

	return nil
}

// HitBasket returns the HitBasket associated with the OtxReviewReview's BasketID (basket_id).
//
// Generated from foreign key 'otx_review_review_basket_id_0bcdb7dc_fk_hit_basket_id'.
func (orr *OtxReviewReview) HitBasket(db XODB) (*HitBasket, error) {
	return HitBasketByID(db, orr.BasketID)
}

// AuthUser returns the AuthUser associated with the OtxReviewReview's ReviewerID (reviewer_id).
//
// Generated from foreign key 'otx_review_review_reviewer_id_1c99d78e_fk_auth_user_id'.
func (orr *OtxReviewReview) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, int(orr.ReviewerID.Int64))
}

// OtxReviewReviewsByReviewerID retrieves a row from 'public.otx_review_review' as a OtxReviewReview.
//
// Generated from index 'otx_review_review_071d8141'.
func OtxReviewReviewsByReviewerID(db XODB, reviewerID sql.NullInt64) ([]*OtxReviewReview, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, time, reviewed, basket_id, reviewer_id, changed ` +
		`FROM public.otx_review_review ` +
		`WHERE reviewer_id = $1`

	// run query
	XOLog(sqlstr, reviewerID)
	q, err := db.Query(sqlstr, reviewerID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*OtxReviewReview{}
	for q.Next() {
		orr := OtxReviewReview{
			_exists: true,
		}

		// scan
		err = q.Scan(&orr.ID, &orr.Time, &orr.Reviewed, &orr.BasketID, &orr.ReviewerID, &orr.Changed)
		if err != nil {
			return nil, err
		}

		res = append(res, &orr)
	}

	return res, nil
}

// OtxReviewReviewByBasketID retrieves a row from 'public.otx_review_review' as a OtxReviewReview.
//
// Generated from index 'otx_review_review_basket_id_key'.
func OtxReviewReviewByBasketID(db XODB, basketID int) (*OtxReviewReview, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, time, reviewed, basket_id, reviewer_id, changed ` +
		`FROM public.otx_review_review ` +
		`WHERE basket_id = $1`

	// run query
	XOLog(sqlstr, basketID)
	orr := OtxReviewReview{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, basketID).Scan(&orr.ID, &orr.Time, &orr.Reviewed, &orr.BasketID, &orr.ReviewerID, &orr.Changed)
	if err != nil {
		return nil, err
	}

	return &orr, nil
}

// OtxReviewReviewByID retrieves a row from 'public.otx_review_review' as a OtxReviewReview.
//
// Generated from index 'otx_review_review_pkey'.
func OtxReviewReviewByID(db XODB, id int) (*OtxReviewReview, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, time, reviewed, basket_id, reviewer_id, changed ` +
		`FROM public.otx_review_review ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	orr := OtxReviewReview{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&orr.ID, &orr.Time, &orr.Reviewed, &orr.BasketID, &orr.ReviewerID, &orr.Changed)
	if err != nil {
		return nil, err
	}

	return &orr, nil
}