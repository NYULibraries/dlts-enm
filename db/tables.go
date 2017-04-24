/* DO NOT EDIT: this file was generated by db.GenerateDbCode() */

// Copyright © 2017 NYU
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package db

import (
	"database/sql"
)

var insertStmt_Topics *sql.Stmt
var insertStmt_Scopes *sql.Stmt
var insertStmt_RelationType *sql.Stmt
var insertStmt_RelationDirection *sql.Stmt
var insertStmt_Relations *sql.Stmt
var insertStmt_Indexpatterns *sql.Stmt
var insertStmt_Epubs *sql.Stmt
var insertStmt_Locations *sql.Stmt
var insertStmt_Names *sql.Stmt
var insertStmt_Occurrences *sql.Stmt

func prepareInsertStmts() {
	var err error

	insertStmt_Topics, err = DB.Prepare("INSERT INTO topics (tct_id) VALUES (?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Scopes, err = DB.Prepare("INSERT INTO scopes (scope, tct_id) VALUES (?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_RelationType, err = DB.Prepare("INSERT INTO relation_type (role_from, role_to, rtype, symmetrical, tct_id) VALUES (?, ?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_RelationDirection, err = DB.Prepare("INSERT INTO relation_direction (direction, id) VALUES (?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Relations, err = DB.Prepare("INSERT INTO relations (relation_direction_id, relation_type_id, tct_id, topic_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Indexpatterns, err = DB.Prepare("INSERT INTO indexpatterns (description, indicators_of_occurrence_range, inline_see_also_end, inline_see_also_start, inline_see_end, inline_see_start, name, pagenumber_css_selector_pattern, pagenumber_pre_strings, pagenumber_xpath_pattern, see_also_split_strings, see_split_strings, separator_before_first_subentry, separator_between_entry_and_occurrences, separator_between_seealsos, separator_between_sees, separator_between_subentries, separator_see_subentry, subentry_classes, tct_id, xpath_entry, xpath_occurrence_link, xpath_see, xpath_see_also) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Epubs, err = DB.Prepare("INSERT INTO epubs (author, indexpattern_id, isbn, publisher, tct_id, title) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Locations, err = DB.Prepare("INSERT INTO locations (content_descriptor, content_text, content_unique_descriptor, context, epub_id, localid, next_location_id, pagenumber_css_selector, pagenumber_filepath, pagenumber_tag, pagenumber_xpath, previous_location_id, sequence_number, tct_id) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Names, err = DB.Prepare("INSERT INTO names (bypass, hidden, name, preferred, scope_id, tct_id, topic_id) VALUES (?, ?, ?, ?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
	insertStmt_Occurrences, err = DB.Prepare("INSERT INTO occurrences (ring_next, ring_prev, tct_id, topic_id) VALUES (?, ?, ?, ?)")
	if err != nil {
		panic("db.prepareInsertStatements: " + err.Error())
	}
		
}
