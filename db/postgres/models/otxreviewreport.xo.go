// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
	"time"
)

// OtxReviewReport represents a row from 'public.otx_review_report'.
type OtxReviewReport struct {
	ID       int       `json:"id"`        // id
	Time     time.Time `json:"time"`      // time
	TopicSet string    `json:"topic_set"` // topic_set
	Location string    `json:"location"`  // location

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OtxReviewReport exists in the database.
func (orr *OtxReviewReport) Exists() bool {
	return orr._exists
}

// Deleted provides information if the OtxReviewReport has been deleted from the database.
func (orr *OtxReviewReport) Deleted() bool {
	return orr._deleted
}

// Insert inserts the OtxReviewReport to the database.
func (orr *OtxReviewReport) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if orr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.otx_review_report (` +
		`time, topic_set, location` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, orr.Time, orr.TopicSet, orr.Location)
	err = db.QueryRow(sqlstr, orr.Time, orr.TopicSet, orr.Location).Scan(&orr.ID)
	if err != nil {
		return err
	}

	// set existence
	orr._exists = true

	return nil
}

// Update updates the OtxReviewReport in the database.
func (orr *OtxReviewReport) Update(db XODB) error {
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
	const sqlstr = `UPDATE public.otx_review_report SET (` +
		`time, topic_set, location` +
		`) = ( ` +
		`$1, $2, $3` +
		`) WHERE id = $4`

	// run query
	XOLog(sqlstr, orr.Time, orr.TopicSet, orr.Location, orr.ID)
	_, err = db.Exec(sqlstr, orr.Time, orr.TopicSet, orr.Location, orr.ID)
	return err
}

// Save saves the OtxReviewReport to the database.
func (orr *OtxReviewReport) Save(db XODB) error {
	if orr.Exists() {
		return orr.Update(db)
	}

	return orr.Insert(db)
}

// Upsert performs an upsert for OtxReviewReport.
//
// NOTE: PostgreSQL 9.5+ only
func (orr *OtxReviewReport) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if orr._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.otx_review_report (` +
		`id, time, topic_set, location` +
		`) VALUES (` +
		`$1, $2, $3, $4` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, time, topic_set, location` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.time, EXCLUDED.topic_set, EXCLUDED.location` +
		`)`

	// run query
	XOLog(sqlstr, orr.ID, orr.Time, orr.TopicSet, orr.Location)
	_, err = db.Exec(sqlstr, orr.ID, orr.Time, orr.TopicSet, orr.Location)
	if err != nil {
		return err
	}

	// set existence
	orr._exists = true

	return nil
}

// Delete deletes the OtxReviewReport from the database.
func (orr *OtxReviewReport) Delete(db XODB) error {
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
	const sqlstr = `DELETE FROM public.otx_review_report WHERE id = $1`

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

// OtxReviewReportByID retrieves a row from 'public.otx_review_report' as a OtxReviewReport.
//
// Generated from index 'otx_review_report_pkey'.
func OtxReviewReportByID(db XODB, id int) (*OtxReviewReport, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, time, topic_set, location ` +
		`FROM public.otx_review_report ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	orr := OtxReviewReport{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&orr.ID, &orr.Time, &orr.TopicSet, &orr.Location)
	if err != nil {
		return nil, err
	}

	return &orr, nil
}
