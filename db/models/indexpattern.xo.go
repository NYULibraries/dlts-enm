// Package models contains the types for schema 'enm'.
package models

// GENERATED BY XO. DO NOT EDIT.

import (
	"errors"
)

// Indexpattern represents a row from 'enm.indexpatterns'.
type Indexpattern struct {
	TctID                               int    `json:"tct_id"`                                  // tct_id
	Name                                string `json:"name"`                                    // name
	Description                         string `json:"description"`                             // description
	PagenumberPreStrings                string `json:"pagenumber_pre_strings"`                  // pagenumber_pre_strings
	PagenumberCSSSelectorPattern        string `json:"pagenumber_css_selector_pattern"`         // pagenumber_css_selector_pattern
	PagenumberXpathPattern              string `json:"pagenumber_xpath_pattern"`                // pagenumber_xpath_pattern
	XpathEntry                          string `json:"xpath_entry"`                             // xpath_entry
	SeeSplitStrings                     string `json:"see_split_strings"`                       // see_split_strings
	SeeAlsoSplitStrings                 string `json:"see_also_split_strings"`                  // see_also_split_strings
	XpathSee                            string `json:"xpath_see"`                               // xpath_see
	XpathSeeAlso                        string `json:"xpath_see_also"`                          // xpath_see_also
	SeparatorBetweenSees                string `json:"separator_between_sees"`                  // separator_between_sees
	SeparatorBetweenSeealsos            string `json:"separator_between_seealsos"`              // separator_between_seealsos
	SeparatorSeeSubentry                string `json:"separator_see_subentry"`                  // separator_see_subentry
	InlineSeeStart                      string `json:"inline_see_start"`                        // inline_see_start
	InlineSeeAlsoStart                  string `json:"inline_see_also_start"`                   // inline_see_also_start
	InlineSeeEnd                        string `json:"inline_see_end"`                          // inline_see_end
	InlineSeeAlsoEnd                    string `json:"inline_see_also_end"`                     // inline_see_also_end
	SubentryClasses                     string `json:"subentry_classes"`                        // subentry_classes
	SeparatorBetweenSubentries          string `json:"separator_between_subentries"`            // separator_between_subentries
	SeparatorBetweenEntryAndOccurrences string `json:"separator_between_entry_and_occurrences"` // separator_between_entry_and_occurrences
	SeparatorBeforeFirstSubentry        string `json:"separator_before_first_subentry"`         // separator_before_first_subentry
	XpathOccurrenceLink                 string `json:"xpath_occurrence_link"`                   // xpath_occurrence_link
	IndicatorsOfOccurrenceRange         string `json:"indicators_of_occurrence_range"`          // indicators_of_occurrence_range

	// xo fields
	_exists, _deleted bool
}

// Exists determines if the Indexpattern exists in the database.
func (i *Indexpattern) Exists() bool {
	return i._exists
}

// Deleted provides information if the Indexpattern has been deleted from the database.
func (i *Indexpattern) Deleted() bool {
	return i._deleted
}

// Insert inserts the Indexpattern to the database.
func (i *Indexpattern) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if i._exists {
		return errors.New("insert failed: already exists")
	}

	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO enm.indexpatterns (` +
		`tct_id, name, description, pagenumber_pre_strings, pagenumber_css_selector_pattern, pagenumber_xpath_pattern, xpath_entry, see_split_strings, see_also_split_strings, xpath_see, xpath_see_also, separator_between_sees, separator_between_seealsos, separator_see_subentry, inline_see_start, inline_see_also_start, inline_see_end, inline_see_also_end, subentry_classes, separator_between_subentries, separator_between_entry_and_occurrences, separator_before_first_subentry, xpath_occurrence_link, indicators_of_occurrence_range` +
		`) VALUES (` +
		`?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, i.TctID, i.Name, i.Description, i.PagenumberPreStrings, i.PagenumberCSSSelectorPattern, i.PagenumberXpathPattern, i.XpathEntry, i.SeeSplitStrings, i.SeeAlsoSplitStrings, i.XpathSee, i.XpathSeeAlso, i.SeparatorBetweenSees, i.SeparatorBetweenSeealsos, i.SeparatorSeeSubentry, i.InlineSeeStart, i.InlineSeeAlsoStart, i.InlineSeeEnd, i.InlineSeeAlsoEnd, i.SubentryClasses, i.SeparatorBetweenSubentries, i.SeparatorBetweenEntryAndOccurrences, i.SeparatorBeforeFirstSubentry, i.XpathOccurrenceLink, i.IndicatorsOfOccurrenceRange)
	_, err = db.Exec(sqlstr, i.TctID, i.Name, i.Description, i.PagenumberPreStrings, i.PagenumberCSSSelectorPattern, i.PagenumberXpathPattern, i.XpathEntry, i.SeeSplitStrings, i.SeeAlsoSplitStrings, i.XpathSee, i.XpathSeeAlso, i.SeparatorBetweenSees, i.SeparatorBetweenSeealsos, i.SeparatorSeeSubentry, i.InlineSeeStart, i.InlineSeeAlsoStart, i.InlineSeeEnd, i.InlineSeeAlsoEnd, i.SubentryClasses, i.SeparatorBetweenSubentries, i.SeparatorBetweenEntryAndOccurrences, i.SeparatorBeforeFirstSubentry, i.XpathOccurrenceLink, i.IndicatorsOfOccurrenceRange)
	if err != nil {
		return err
	}

	// set existence
	i._exists = true

	return nil
}

// Update updates the Indexpattern in the database.
func (i *Indexpattern) Update(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return errors.New("update failed: does not exist")
	}

	// if deleted, bail
	if i._deleted {
		return errors.New("update failed: marked for deletion")
	}

	// sql query
	const sqlstr = `UPDATE enm.indexpatterns SET ` +
		`name = ?, description = ?, pagenumber_pre_strings = ?, pagenumber_css_selector_pattern = ?, pagenumber_xpath_pattern = ?, xpath_entry = ?, see_split_strings = ?, see_also_split_strings = ?, xpath_see = ?, xpath_see_also = ?, separator_between_sees = ?, separator_between_seealsos = ?, separator_see_subentry = ?, inline_see_start = ?, inline_see_also_start = ?, inline_see_end = ?, inline_see_also_end = ?, subentry_classes = ?, separator_between_subentries = ?, separator_between_entry_and_occurrences = ?, separator_before_first_subentry = ?, xpath_occurrence_link = ?, indicators_of_occurrence_range = ?` +
		` WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, i.Name, i.Description, i.PagenumberPreStrings, i.PagenumberCSSSelectorPattern, i.PagenumberXpathPattern, i.XpathEntry, i.SeeSplitStrings, i.SeeAlsoSplitStrings, i.XpathSee, i.XpathSeeAlso, i.SeparatorBetweenSees, i.SeparatorBetweenSeealsos, i.SeparatorSeeSubentry, i.InlineSeeStart, i.InlineSeeAlsoStart, i.InlineSeeEnd, i.InlineSeeAlsoEnd, i.SubentryClasses, i.SeparatorBetweenSubentries, i.SeparatorBetweenEntryAndOccurrences, i.SeparatorBeforeFirstSubentry, i.XpathOccurrenceLink, i.IndicatorsOfOccurrenceRange, i.TctID)
	_, err = db.Exec(sqlstr, i.Name, i.Description, i.PagenumberPreStrings, i.PagenumberCSSSelectorPattern, i.PagenumberXpathPattern, i.XpathEntry, i.SeeSplitStrings, i.SeeAlsoSplitStrings, i.XpathSee, i.XpathSeeAlso, i.SeparatorBetweenSees, i.SeparatorBetweenSeealsos, i.SeparatorSeeSubentry, i.InlineSeeStart, i.InlineSeeAlsoStart, i.InlineSeeEnd, i.InlineSeeAlsoEnd, i.SubentryClasses, i.SeparatorBetweenSubentries, i.SeparatorBetweenEntryAndOccurrences, i.SeparatorBeforeFirstSubentry, i.XpathOccurrenceLink, i.IndicatorsOfOccurrenceRange, i.TctID)
	return err
}

// Save saves the Indexpattern to the database.
func (i *Indexpattern) Save(db XODB) error {
	if i.Exists() {
		return i.Update(db)
	}

	return i.Insert(db)
}

// Delete deletes the Indexpattern from the database.
func (i *Indexpattern) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !i._exists {
		return nil
	}

	// if deleted, bail
	if i._deleted {
		return nil
	}

	// sql query
	const sqlstr = `DELETE FROM enm.indexpatterns WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, i.TctID)
	_, err = db.Exec(sqlstr, i.TctID)
	if err != nil {
		return err
	}

	// set deleted
	i._deleted = true

	return nil
}

// IndexpatternByTctID retrieves a row from 'enm.indexpatterns' as a Indexpattern.
//
// Generated from index 'indexpatterns_tct_id_pkey'.
func IndexpatternByTctID(db XODB, tctID int) (*Indexpattern, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, name, description, pagenumber_pre_strings, pagenumber_css_selector_pattern, pagenumber_xpath_pattern, xpath_entry, see_split_strings, see_also_split_strings, xpath_see, xpath_see_also, separator_between_sees, separator_between_seealsos, separator_see_subentry, inline_see_start, inline_see_also_start, inline_see_end, inline_see_also_end, subentry_classes, separator_between_subentries, separator_between_entry_and_occurrences, separator_before_first_subentry, xpath_occurrence_link, indicators_of_occurrence_range ` +
		`FROM enm.indexpatterns ` +
		`WHERE tct_id = ?`

	// run query
	XOLog(sqlstr, tctID)
	i := Indexpattern{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, tctID).Scan(&i.TctID, &i.Name, &i.Description, &i.PagenumberPreStrings, &i.PagenumberCSSSelectorPattern, &i.PagenumberXpathPattern, &i.XpathEntry, &i.SeeSplitStrings, &i.SeeAlsoSplitStrings, &i.XpathSee, &i.XpathSeeAlso, &i.SeparatorBetweenSees, &i.SeparatorBetweenSeealsos, &i.SeparatorSeeSubentry, &i.InlineSeeStart, &i.InlineSeeAlsoStart, &i.InlineSeeEnd, &i.InlineSeeAlsoEnd, &i.SubentryClasses, &i.SeparatorBetweenSubentries, &i.SeparatorBetweenEntryAndOccurrences, &i.SeparatorBeforeFirstSubentry, &i.XpathOccurrenceLink, &i.IndicatorsOfOccurrenceRange)
	if err != nil {
		return nil, err
	}

	return &i, nil
}

// IndexpatternsByName retrieves a row from 'enm.indexpatterns' as a Indexpattern.
//
// Generated from index 'name'.
func IndexpatternsByName(db XODB, name string) ([]*Indexpattern, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`tct_id, name, description, pagenumber_pre_strings, pagenumber_css_selector_pattern, pagenumber_xpath_pattern, xpath_entry, see_split_strings, see_also_split_strings, xpath_see, xpath_see_also, separator_between_sees, separator_between_seealsos, separator_see_subentry, inline_see_start, inline_see_also_start, inline_see_end, inline_see_also_end, subentry_classes, separator_between_subentries, separator_between_entry_and_occurrences, separator_before_first_subentry, xpath_occurrence_link, indicators_of_occurrence_range ` +
		`FROM enm.indexpatterns ` +
		`WHERE name = ?`

	// run query
	XOLog(sqlstr, name)
	q, err := db.Query(sqlstr, name)
	if err != nil {
		return nil, err
	}
	defer q.Close()

	// load results
	res := []*Indexpattern{}
	for q.Next() {
		i := Indexpattern{
			_exists: true,
		}

		// scan
		err = q.Scan(&i.TctID, &i.Name, &i.Description, &i.PagenumberPreStrings, &i.PagenumberCSSSelectorPattern, &i.PagenumberXpathPattern, &i.XpathEntry, &i.SeeSplitStrings, &i.SeeAlsoSplitStrings, &i.XpathSee, &i.XpathSeeAlso, &i.SeparatorBetweenSees, &i.SeparatorBetweenSeealsos, &i.SeparatorSeeSubentry, &i.InlineSeeStart, &i.InlineSeeAlsoStart, &i.InlineSeeEnd, &i.InlineSeeAlsoEnd, &i.SubentryClasses, &i.SeparatorBetweenSubentries, &i.SeparatorBetweenEntryAndOccurrences, &i.SeparatorBeforeFirstSubentry, &i.XpathOccurrenceLink, &i.IndicatorsOfOccurrenceRange)
		if err != nil {
			return nil, err
		}

		res = append(res, &i)
	}

	return res, nil
}
