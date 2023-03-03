package database

func (db *appdbimpl) UpdateFollowings(username string, followingusername string) error {

	rows, err := db.c.Query(`SELECT id FROM users WHERE username=?`, username)
	if err != nil {
		return err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	var userId uint64
	err = rows.Scan(&userId)
	if err != nil {
		return err
	}

	if rows.Err() != nil {
		return err
	}

	rows2, err2 := db.c.Query(`SELECT id FROM users WHERE username=?`, followingusername)
	if err2 != nil {
		return err2
	}
	defer func() { _ = rows2.Close() }()

	// Here we read the resultset and we build the list to be returned
	var followingId uint64
	err2 = rows2.Scan(&followingId)
	if err2 != nil {
		return err2
	}

	if rows.Err() != nil {
		return err2
	}

	res, err := db.c.Exec(`INSERT INTO followers (user_id, follower_id) VALUES (?, ?) WHERE NOT EXISTS (SELECT * FROM followers WHERE user_id=? AND follower_id=?)`,
		followingId, userId)

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return nil
	}
	return nil
}
