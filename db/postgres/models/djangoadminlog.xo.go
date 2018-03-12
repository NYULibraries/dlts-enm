// Package models contains the types for schema 'public'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"database/sql"
	"errors"
	"time"
)

// DjangoAdminLog represents a row from 'public.django_admin_log'.
type DjangoAdminLog struct {
	ID            int            `json:"id"`              // id
	ActionTime    time.Time      `json:"action_time"`     // action_time
	ObjectID      sql.NullString `json:"object_id"`       // object_id
	ObjectRepr    string         `json:"object_repr"`     // object_repr
	ActionFlag    int16          `json:"action_flag"`     // action_flag
	ChangeMessage string         `json:"change_message"`  // change_message
	ContentTypeID sql.NullInt64  `json:"content_type_id"` // content_type_id
	UserID        int            `json:"user_id"`         // user_id

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the DjangoAdminLog exists in the database.
func (dal *DjangoAdminLog) Exists() bool {
	return dal._exists
}

// Deleted provides information if the DjangoAdminLog has been deleted from the database.
func (dal *DjangoAdminLog) Deleted() bool {
	return dal._deleted
}

// Insert inserts the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if dal._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key provided by sequence
	const sqlstr = `INSERT INTO public.django_admin_log (` +
		`action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) RETURNING id`

	// run query
	XOLog(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	err = db.QueryRow(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID).Scan(&dal.ID)
	if err != nil {
		return err
	}

	// set existence
	dal._exists = true

	return nil
}

// Update updates the DjangoAdminLog in the database.
func (dal *DjangoAdminLog) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if dal._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE public.django_admin_log SET (` +
		`action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) = ( ` +
		`$1, $2, $3, $4, $5, $6, $7` +
		`) WHERE id = $8`

	// run query
	XOLog(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID)
	_, err = db.Exec(sqlstr, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID, dal.ID)
	return err
}

// Save saves the DjangoAdminLog to the database.
func (dal *DjangoAdminLog) Save(db XODB) error {
	if dal.Exists() {
		return dal.Update(db)
	}

	return dal.Insert(db)
}

// Upsert performs an upsert for DjangoAdminLog.
//
// NOTE: PostgreSQL 9.5+ only
func (dal *DjangoAdminLog) Upsert(db XODB) error {
	var err error

	// if already exist, bail
	if dal._exists {
		return errors.New("insert failed: already exists")
	}

	// sql query
	const sqlstr = `INSERT INTO public.django_admin_log (` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) VALUES (` +
		`$1, $2, $3, $4, $5, $6, $7, $8` +
		`) ON CONFLICT (id) DO UPDATE SET (` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id` +
		`) = (` +
		`EXCLUDED.id, EXCLUDED.action_time, EXCLUDED.object_id, EXCLUDED.object_repr, EXCLUDED.action_flag, EXCLUDED.change_message, EXCLUDED.content_type_id, EXCLUDED.user_id` +
		`)`

	// run query
	XOLog(sqlstr, dal.ID, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	_, err = db.Exec(sqlstr, dal.ID, dal.ActionTime, dal.ObjectID, dal.ObjectRepr, dal.ActionFlag, dal.ChangeMessage, dal.ContentTypeID, dal.UserID)
	if err != nil {
		return err
	}

	// set existence
	dal._exists = true

	return nil
}

// Delete deletes the DjangoAdminLog from the database.
func (dal *DjangoAdminLog) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !dal._exists {
		return nil
	}

	// if deleted, bail
	if dal._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM public.django_admin_log WHERE id = $1`

	// run query
	XOLog(sqlstr, dal.ID)
	_, err = db.Exec(sqlstr, dal.ID)
	if err != nil {
		return err
	}

	// set deleted
	dal._deleted = true

	return nil
}

// DjangoContentType returns the DjangoContentType associated with the DjangoAdminLog's ContentTypeID (content_type_id).
//
// Generated from foreign key 'django_admin_content_type_id_c4bce8eb_fk_django_content_type_id'.
func (dal *DjangoAdminLog) DjangoContentType(db XODB) (*DjangoContentType, error) {
	return DjangoContentTypeByID(db, int(dal.ContentTypeID.Int64))
}

// AuthUser returns the AuthUser associated with the DjangoAdminLog's UserID (user_id).
//
// Generated from foreign key 'django_admin_log_user_id_c564eba6_fk_auth_user_id'.
func (dal *DjangoAdminLog) AuthUser(db XODB) (*AuthUser, error) {
	return AuthUserByID(db, dal.UserID)
}

// DjangoAdminLogsByContentTypeID retrieves a row from 'public.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_417f1b1c'.
func DjangoAdminLogsByContentTypeID(db XODB, contentTypeID sql.NullInt64) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM public.django_admin_log ` +
		`WHERE content_type_id = $1`

	// run query
	XOLog(sqlstr, contentTypeID)
	q, err := db.Query(sqlstr, contentTypeID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogsByUserID retrieves a row from 'public.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_e8701ad4'.
func DjangoAdminLogsByUserID(db XODB, userID int) ([]*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM public.django_admin_log ` +
		`WHERE user_id = $1`

	// run query
	XOLog(sqlstr, userID)
	q, err := db.Query(sqlstr, userID)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*DjangoAdminLog{}
	for q.Next() {
		dal := DjangoAdminLog{
			_exists: true,
		}

		// scan
		err = q.Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
		if err != nil {
			return nil, err
		}

		res = append(res, &dal)
	}

	return res, nil
}

// DjangoAdminLogByID retrieves a row from 'public.django_admin_log' as a DjangoAdminLog.
//
// Generated from index 'django_admin_log_pkey'.
func DjangoAdminLogByID(db XODB, id int) (*DjangoAdminLog, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, action_time, object_id, object_repr, action_flag, change_message, content_type_id, user_id ` +
		`FROM public.django_admin_log ` +
		`WHERE id = $1`

	// run query
	XOLog(sqlstr, id)
	dal := DjangoAdminLog{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&dal.ID, &dal.ActionTime, &dal.ObjectID, &dal.ObjectRepr, &dal.ActionFlag, &dal.ChangeMessage, &dal.ContentTypeID, &dal.UserID)
	if err != nil {
		return nil, err
	}

	return &dal, nil
}
