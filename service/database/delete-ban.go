package database

import "database/sql"

func (db *appdbimpl) DeleteBan(username string, bannedusername string) error {

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)
	var userId uint64
	if err := row.Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err
	}

	if err := row.Err(); err != nil {
		return err
	}

	row2 := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, bannedusername)
	var bannedId uint64
	if err2 := row2.Scan(&bannedId); err2 != nil {
		if err2 == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err2
	}

	if err2 := row2.Err(); err2 != nil {
		return err2
	}
	res, err := db.c.Exec(`DELETE FROM bans WHERE user_id=? AND banned_id=?`, userId, bannedId)

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrBanDoesNotExist
	}
	return nil
}
