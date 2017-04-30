// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// Location represents a row from 'enm.locations'.
type Location struct {
	TctID                   int    `json:"tct_id"`                    // tct_id
	EpubID                  int    `json:"epub_id"`                   // epub_id
	Localid                 int    `json:"localid"`                   // localid
	SequenceNumber          int    `json:"sequence_number"`           // sequence_number
	ContentUniqueDescriptor string `json:"content_unique_descriptor"` // content_unique_descriptor
	ContentDescriptor       string `json:"content_descriptor"`        // content_descriptor
	ContentText             string `json:"content_text"`              // content_text
	PagenumberFilepath      string `json:"pagenumber_filepath"`       // pagenumber_filepath
	PagenumberTag           string `json:"pagenumber_tag"`            // pagenumber_tag
	PagenumberCSSSelector   string `json:"pagenumber_css_selector"`   // pagenumber_css_selector
	PagenumberXpath         string `json:"pagenumber_xpath"`          // pagenumber_xpath
	NextLocationID          int    `json:"next_location_id"`          // next_location_id
	PreviousLocationID      int    `json:"previous_location_id"`      // previous_location_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Location exists in the database.
func (l *Location) Exists() bool {
	return l._exists
}

// Deleted provides information if the Location has been deleted from the database.
func (l *Location) Deleted() bool {
	return l._deleted
}

// Insert inserts the Location to the database.
func (l *Location) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if l._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO enm.locations (` +
		`tct_id, epub_id, localid, sequence_number, content_unique_descriptor, content_descriptor, content_text, pagenumber_filepath, pagenumber_tag, pagenumber_css_selector, pagenumber_xpath, next_location_id, previous_location_id` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, l.TctID, l.EpubID, l.Localid, l.SequenceNumber, l.ContentUniqueDescriptor, l.ContentDescriptor, l.ContentText, l.PagenumberFilepath, l.PagenumberTag, l.PagenumberCSSSelector, l.PagenumberXpath, l.NextLocationID, l.PreviousLocationID)
	_, err = db.Exec(sqlstr, l.TctID, l.EpubID, l.Localid, l.SequenceNumber, l.ContentUniqueDescriptor, l.ContentDescriptor, l.ContentText, l.PagenumberFilepath, l.PagenumberTag, l.PagenumberCSSSelector, l.PagenumberXpath, l.NextLocationID, l.PreviousLocationID)
	if err != nil {
		return err
	}

	// set existence
	l._exists = true

	return nil
}

// Update updates the Location in the database.
func (l *Location) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if l._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE enm.locations SET ` +
		`epub_id = ?, localid = ?, sequence_number = ?, content_unique_descriptor = ?, content_descriptor = ?, content_text = ?, pagenumber_filepath = ?, pagenumber_tag = ?, pagenumber_css_selector = ?, pagenumber_xpath = ?, next_location_id = ?, previous_location_id = ?` +
		` WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, l.EpubID, l.Localid, l.SequenceNumber, l.ContentUniqueDescriptor, l.ContentDescriptor, l.ContentText, l.PagenumberFilepath, l.PagenumberTag, l.PagenumberCSSSelector, l.PagenumberXpath, l.NextLocationID, l.PreviousLocationID, l.TctID)
	_, err = db.Exec(sqlstr, l.EpubID, l.Localid, l.SequenceNumber, l.ContentUniqueDescriptor, l.ContentDescriptor, l.ContentText, l.PagenumberFilepath, l.PagenumberTag, l.PagenumberCSSSelector, l.PagenumberXpath, l.NextLocationID, l.PreviousLocationID, l.TctID)
	return err
}

// Save saves the Location to the database.
func (l *Location) Save(db XODB) error {
	if l.Exists() {
		return l.Update(db)
	}

	return l.Insert(db)
}

// Delete deletes the Location from the database.
func (l *Location) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !l._exists {
		return nil
	}

	// if deleted, bail
	if l._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM enm.locations WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, l.TctID)
	_, err = db.Exec(sqlstr, l.TctID)
	if err != nil {
		return err
	}

	// set deleted
	l._deleted = true

	return nil
}

// Epub returns the Epub associated with the Location's EpubID (epub_id).
//
// Generated from foreign key 'fk__locations__epubs'.
func (l *Location) Epub(db XODB) (*Epub, error) {
	return EpubByTctID(db, l.EpubID)
}

// LocationByNextLocationID returns the Location associated with the Location's NextLocationID (next_location_id).
//
// Generated from foreign key 'fk__locations__next_location_id__locations__tct_id'.
func (l *Location) LocationByNextLocationID(db XODB) (*Location, error) {
	return LocationByTctID(db, l.NextLocationID)
}

// LocationByPreviousLocationID returns the Location associated with the Location's PreviousLocationID (previous_location_id).
//
// Generated from foreign key 'fk__locations__previous_location_id__locations__tct_id'.
func (l *Location) LocationByPreviousLocationID(db XODB) (*Location, error) {
	return LocationByTctID(db, l.PreviousLocationID)
}

// LocationsByEpubID retrieves a row from 'enm.locations' as a Location.
//
// Generated from index 'epub_id'.
func LocationsByEpubID(db XODB, epubID int) ([]*Location, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, epub_id, localid, sequence_number, content_unique_descriptor, content_descriptor, content_text, pagenumber_filepath, pagenumber_tag, pagenumber_css_selector, pagenumber_xpath, next_location_id, previous_location_id ` +
		`FROM enm.locations ` +
		`WHERE epub_id = ?`

	// run query
	XOLog(sqlstr, epubID)
	q, err := db.Query(sqlstr, epubID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Location{}
	for q.Next() {
		l := Location{
			_exists: true,
		}

		// scan
		err = q.Scan(&l.TctID, &l.EpubID, &l.Localid, &l.SequenceNumber, &l.ContentUniqueDescriptor, &l.ContentDescriptor, &l.ContentText, &l.PagenumberFilepath, &l.PagenumberTag, &l.PagenumberCSSSelector, &l.PagenumberXpath, &l.NextLocationID, &l.PreviousLocationID)
		if err != nil {
			return nil, err
		}

		res = append(res, &l)
	}

	return res, nil
}

// LocationByTctID retrieves a row from 'enm.locations' as a Location.
//
// Generated from index 'locations_tct_id_pkey'.
func LocationByTctID(db XODB, tctID int) (*Location, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, epub_id, localid, sequence_number, content_unique_descriptor, content_descriptor, content_text, pagenumber_filepath, pagenumber_tag, pagenumber_css_selector, pagenumber_xpath, next_location_id, previous_location_id ` +
		`FROM enm.locations ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	l := Location{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, tctID).Scan(&l.TctID, &l.EpubID, &l.Localid, &l.SequenceNumber, &l.ContentUniqueDescriptor, &l.ContentDescriptor, &l.ContentText, &l.PagenumberFilepath, &l.PagenumberTag, &l.PagenumberCSSSelector, &l.PagenumberXpath, &l.NextLocationID, &l.PreviousLocationID)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// LocationByNextLocationID retrieves a row from 'enm.locations' as a Location.
//
// Generated from index 'next_location_id'.
func LocationByNextLocationID(db XODB, nextLocationID int) (*Location, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, epub_id, localid, sequence_number, content_unique_descriptor, content_descriptor, content_text, pagenumber_filepath, pagenumber_tag, pagenumber_css_selector, pagenumber_xpath, next_location_id, previous_location_id ` +
		`FROM enm.locations ` +
		`WHERE next_location_id = ?`

	// run query
	XOLog(sqlstr, nextLocationID)
	l := Location{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, nextLocationID).Scan(&l.TctID, &l.EpubID, &l.Localid, &l.SequenceNumber, &l.ContentUniqueDescriptor, &l.ContentDescriptor, &l.ContentText, &l.PagenumberFilepath, &l.PagenumberTag, &l.PagenumberCSSSelector, &l.PagenumberXpath, &l.NextLocationID, &l.PreviousLocationID)
	if err != nil {
		return nil, err
	}

	return &l, nil
}

// LocationByPreviousLocationID retrieves a row from 'enm.locations' as a Location.
//
// Generated from index 'previous_location_id'.
func LocationByPreviousLocationID(db XODB, previousLocationID int) (*Location, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, epub_id, localid, sequence_number, content_unique_descriptor, content_descriptor, content_text, pagenumber_filepath, pagenumber_tag, pagenumber_css_selector, pagenumber_xpath, next_location_id, previous_location_id ` +
		`FROM enm.locations ` +
		`WHERE previous_location_id = ?`

	// run query
	XOLog(sqlstr, previousLocationID)
	l := Location{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, previousLocationID).Scan(&l.TctID, &l.EpubID, &l.Localid, &l.SequenceNumber, &l.ContentUniqueDescriptor, &l.ContentDescriptor, &l.ContentText, &l.PagenumberFilepath, &l.PagenumberTag, &l.PagenumberCSSSelector, &l.PagenumberXpath, &l.NextLocationID, &l.PreviousLocationID)
	if err != nil {
		return nil, err
	}

	return &l, nil
}
