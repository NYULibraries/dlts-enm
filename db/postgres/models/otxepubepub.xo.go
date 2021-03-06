// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// OtxEpubEpub represents a row from 'public.otx_epub_epub'.
type OtxEpubEpub struct {
	DocumentPtrID int    `json:"document_ptr_id"` // document_ptr_id
	Publisher     string `json:"publisher"`       // publisher
	Source        string `json:"source"`          // source
	OebpsFolder   string `json:"oebps_folder"`    // oebps_folder
	Manifest      string `json:"manifest"`        // manifest
	Contents      string `json:"contents"`        // contents

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the OtxEpubEpub exists in the database.
func (oee *OtxEpubEpub) Exists() bool {
	return oee._exists
}

// Deleted provides information if the OtxEpubEpub has been deleted from the database.
func (oee *OtxEpubEpub) Deleted() bool {
	return oee._deleted
}

// Insert inserts the OtxEpubEpub to the database.
func (oee *OtxEpubEpub) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if oee._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO public.otx_epub_epub (` +
		`document_ptr_id, publisher, source, oebps_folder, manifest, contents` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`)`

	// run query
	XOLog(sqlstr, oee.DocumentPtrID, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents)
	err = db.QueryRow(sqlstr, oee.DocumentPtrID, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents).Scan(&oee.DocumentPtrID)
	if err != nil {
		return err
	}

	// set existence
	oee._exists = true

	return nil
}

// Update updates the OtxEpubEpub in the database.
func (oee *OtxEpubEpub) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oee._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if oee._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.otx_epub_epub SET (` +
		`publisher, source, oebps_folder, manifest, contents` +
		`) = ( ` +
		`$1, $2, $3, $4, $5` +
		`) WHERE document_ptr_id = $6`

	// run query
	XOLog(sqlstr, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents, oee.DocumentPtrID)
	_, err = db.Exec(sqlstr, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents, oee.DocumentPtrID)
	return err
}

// Save saves the OtxEpubEpub to the database.
func (oee *OtxEpubEpub) Save(db XODB) error {
	if oee.Exists() {
		return oee.Update(db)
	}

	return oee.Insert(db)
}

// Upsert performs an upsert for OtxEpubEpub.
//
// NOTE: PostgreSQL 9.5+ only
func (oee *OtxEpubEpub) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if oee._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.otx_epub_epub (` +
		`document_ptr_id, publisher, source, oebps_folder, manifest, contents` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6` +
		`) ON CONFLICT (document_ptr_id) DO UPDATE SET (` +
		`document_ptr_id, publisher, source, oebps_folder, manifest, contents` +
		`) = (` +
		`EXCLUDED.document_ptr_id, EXCLUDED.publisher, EXCLUDED.source, EXCLUDED.oebps_folder, EXCLUDED.manifest, EXCLUDED.contents` +
		`)`

	// run query
	XOLog(sqlstr, oee.DocumentPtrID, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents)
	_, err = db.Exec(sqlstr, oee.DocumentPtrID, oee.Publisher, oee.Source, oee.OebpsFolder, oee.Manifest, oee.Contents)
	if err != nil {
		return err
	}

	// set existence
	oee._exists = true

	return nil
}

// Delete deletes the OtxEpubEpub from the database.
func (oee *OtxEpubEpub) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !oee._exists {
		return nil
	}

	// if deleted, bail
	if oee._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.otx_epub_epub WHERE document_ptr_id = $1`

	// run query
	XOLog(sqlstr, oee.DocumentPtrID)
	_, err = db.Exec(sqlstr, oee.DocumentPtrID)
	if err != nil {
		return err
	}

	// set deleted
	oee._deleted = true

	return nil
}

// OccurrenceDocument returns the OccurrenceDocument associated with the OtxEpubEpub's DocumentPtrID (document_ptr_id).
//
// Generated from foreign key 'otx_epub_epu_document_ptr_id_af55e311_fk_occurrence_document_id'.
func (oee *OtxEpubEpub) OccurrenceDocument(db XODB) (*OccurrenceDocument, error) {
	return OccurrenceDocumentByID(db, oee.DocumentPtrID)
}

// OtxEpubEpubByDocumentPtrID retrieves a row from 'public.otx_epub_epub' as a OtxEpubEpub.
//
// Generated from index 'otx_epub_epub_pkey'.
func OtxEpubEpubByDocumentPtrID(db XODB, documentPtrID int) (*OtxEpubEpub, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`document_ptr_id, publisher, source, oebps_folder, manifest, contents ` +
		`FROM public.otx_epub_epub ` +
		`WHERE document_ptr_id = $1`

	// run query
	XOLog(sqlstr, documentPtrID)
	oee := OtxEpubEpub{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, documentPtrID).Scan(&oee.DocumentPtrID, &oee.Publisher, &oee.Source, &oee.OebpsFolder, &oee.Manifest, &oee.Contents)
	if err != nil {
		return nil, err
	}

	return &oee, nil
}
