// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// DjangoContentType represents a row from 'public.django_content_type'.
type DjangoContentType struct {
	ID       int    `json:"id"`        // id
	AppLabel string `json:"app_label"` // app_label
	Model    string `json:"model"`     // model

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoContentType exists in the database.
func (dct *DjangoContentType) Exists() bool {
	return dct._exists
}

// Deleted provides information if the DjangoContentType has been deleted from the database.
func (dct *DjangoContentType) Deleted() bool {
	return dct._deleted
}

// Insert inserts the DjangoContentType to the database.
func (dct *DjangoContentType) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dct._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.django_content_type (` +
		`app_label, model` +
		`) VALUES (` +
		`$1, $2` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, dct.AppLabel, dct.Model)
	err = db.QueryRow(sqlstr, dct.AppLabel, dct.Model).Scan(&dct.ID)
	if err != nil {
		return err
	}

	// set existence
	dct._exists = true

	return nil
}

// Update updates the DjangoContentType in the database.
func (dct *DjangoContentType) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dct._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dct._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.django_content_type SET (` +
		`app_label, model` +
		`) = ( ` +
		`$1, $2` +
		`) WHERE id = $3`

	// run query
	XOLog(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	_, err = db.Exec(sqlstr, dct.AppLabel, dct.Model, dct.ID)
	return err
}

// Save saves the DjangoContentType to the database.
func (dct *DjangoContentType) Save(db XODB) error {
	if dct.Exists() {
		return dct.Update(db)
	}

	return dct.Insert(db)
}

// Upsert performs an upsert for DjangoContentType.
//
// NOTE: PostgreSQL 9.5+ only
func (dct *DjangoContentType) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if dct._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.django_content_type (` +
		`id, app_label, model` +
		`) VALUES (` +
		`$1, $2, $3` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, app_label, model` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.app_label, EXCLUDED.model` +
		`)`

	// run query
	XOLog(sqlstr, dct.ID, dct.AppLabel, dct.Model)
	_, err = db.Exec(sqlstr, dct.ID, dct.AppLabel, dct.Model)
	if err != nil {
		return err
	}

	// set existence
	dct._exists = true

	return nil
}

// Delete deletes the DjangoContentType from the database.
func (dct *DjangoContentType) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dct._exists {
		return nil
	}

	// if deleted, bail
	if dct._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.django_content_type WHERE id = $1`

	// run query
	XOLog(sqlstr, dct.ID)
	_, err = db.Exec(sqlstr, dct.ID)
	if err != nil {
		return err
	}

	// set deleted
	dct._deleted = true

	return nil
}

// DjangoContentTypeByAppLabelModel retrieves a row from 'public.django_content_type' as a DjangoContentType.
//
// Generated from index 'django_content_type_app_label_76bd3d3b_uniq'.
func DjangoContentTypeByAppLabelModel(db XODB, appLabel string, model string) (*DjangoContentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM public.django_content_type ` +
		`WHERE app_label = $1 AND model = $2`

	// run query
	XOLog(sqlstr, appLabel, model)
	dct := DjangoContentType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, appLabel, model).Scan(&dct.ID, &dct.AppLabel, &dct.Model)
	if err != nil {
		return nil, err
	}

	return &dct, nil
}

// DjangoContentTypeByID retrieves a row from 'public.django_content_type' as a DjangoContentType.
//
// Generated from index 'django_content_type_pkey'.
func DjangoContentTypeByID(db XODB, id int) (*DjangoContentType, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, app_label, model ` +
		`FROM public.django_content_type ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	dct := DjangoContentType{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dct.ID, &dct.AppLabel, &dct.Model)
	if err != nil {
		return nil, err
	}

	return &dct, nil
}
