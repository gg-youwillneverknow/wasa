package database

import "database/sql"

func (db *appdbimpl) UpdateBan(username string, bannedusername string) error {
	var userId uint64
	var bannedId uint64

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)

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
	if err2 := row2.Scan(&bannedId); err2 != nil {
		if err2 == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err2
	}
	if err2 := row2.Err(); err2 != nil {
		return err2
	}

	res, err := db.c.Exec(`INSERT INTO bans (user_id, banned_id) VALUES (?, ?)`,
		userId, bannedId)

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrBanAlreadyExist
	}
	return nil
}
