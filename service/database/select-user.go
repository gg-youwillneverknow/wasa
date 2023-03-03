package database

func (db *appdbimpl) SelectUser(username string) (uint64, error) {
	var ret uint64

	// Plain simple SELECT query
	row, err := db.c.Query(`SELECT id FROM users WHERE username=?`, username)
	if err != nil {
		return 0, err
	}
	defer func() { _ = row.Close() }()

	// Here we read the resultset and we build the list to be returned

	err = row.Scan(&ret)
	if err != nil {
		return ret, err
	}

	if row.Err() != nil {
		return ret, err
	}

	return ret, nil
}
