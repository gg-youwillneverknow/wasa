package database

func (db *appdbimpl) DeleteBan(username string, bannedusername string) error {
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

	rows2, err2 := db.c.Query(`SELECT id FROM users WHERE username=?`, bannedusername)
	if err2 != nil {
		return err2
	}
	defer func() { _ = rows2.Close() }()

	// Here we read the resultset and we build the list to be returned
	var bannedId uint64
	err2 = rows2.Scan(&bannedId)
	if err2 != nil {
		return err2
	}

	if rows.Err() != nil {
		return err2
	}

	res, err := db.c.Exec(`DELETE FROM bans (user_id, banned_id) WHERE user_id=? AND banned_id=?)`,
		userId, bannedId)

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrBanDoesNotExist
	}
	return nil
}
