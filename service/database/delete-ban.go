package database

import "database/sql"

func (db *appdbimpl) DeleteBan(username string, bannedusername string) error {
	var userId uint64
	var bannedId uint64
	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)

	if err := row.Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err
	}

	if err2 := row.Err(); err2 != nil {
		return err2
	}

	row2 := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, bannedusername)

	if err3 := row2.Scan(&bannedId); err3 != nil {
		if err3 == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err3
	}

	if err4 := row2.Err(); err4 != nil {
		return err4
	}
	res, err5 := db.c.Exec(`DELETE FROM bans WHERE user_id=? AND banned_id=?`, userId, bannedId)
	if err5 != nil {
		return err5
	}

	affected, err6 := res.RowsAffected()
	if err6 != nil {
		return err6
	} else if affected == 0 {

		return ErrBanDoesNotExist
	}
	return nil
}
