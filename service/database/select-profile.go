package database

func (db *appdbimpl) SelectProfile(username string) (uint32, uint32, uint32, error) {
	var followers uint32
	var followings uint32
	var posts uint32
	// Plain simple SELECT query
	rows, err := db.c.Query(`SELECT followers_num,followings_num,posts_num FROM users WHERE users.username = ?`, username)
	if err != nil {
		return 0, 0, 0, err
	}

	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	err = rows.Scan(&followers, &followings, &posts)
	if err != nil {
		return 0, 0, 0, err
	}

	if rows.Err() != nil {
		return 0, 0, 0, err
	}

	return followers, followings, posts, nil
}
