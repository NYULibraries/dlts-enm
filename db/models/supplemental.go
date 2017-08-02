package models

// MySQL views do not have indexes, so xo doesn't automatically generate utility
// methods for fetching records by index column value.

// PageByID retrieves a row from 'enm.pages' as a Page.
// Modeled after xo template code.
// Note that it is not possible to add _exists_ field to Page type, so we omit it.
func (*Page) PageByID(db XODB, id int) (*Page, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, title, authors, publisher, isbn, page_pattern, page_localid, page_sequence ` +
		`FROM enm.pages ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	p := Page{}

	err = db.QueryRow(sqlstr, id).Scan(
		&p.ID, &p.Title, &p.Authors, &p.Publisher, &p.Isbn,
		&p.PagePattern, &p.PageLocalid, &p.PageSequence,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}

// PageByID retrieves a row from 'enm.pages' as a Page.
// Modeled after xo template code.
// Note that it is not possible to add _exists_ field to Page type, so we omit it.
func (*PageTopic) PageTopicByID(db XODB, id int) (*PageTopic, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`page_id, topic_id, preferred_topic_name,  ` +
		`FROM enm.page_topics ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	p := PageTopic{}

	err = db.QueryRow(sqlstr, id).Scan(
		&p.PageID, &p.TopicID, &p.PreferredTopicName,
	)
	if err != nil {
		return nil, err
	}

	return &p, nil
}