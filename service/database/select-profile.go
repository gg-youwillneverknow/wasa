package database

import "database/sql"

func (db *appdbimpl) SelectProfile(username string) (uint32, uint32, uint32, error) {
	var followers uint32
	var followings uint32
	var posts uint32

	rows := db.c.QueryRow(`SELECT followers_num,followings_num,posts_num FROM users WHERE users.username = ?`, username)

	if err := rows.Scan(&followers, &followings, &posts); err != nil {
		if err == sql.ErrNoRows {
			return 0, 0, 0, ErrUserDoesNotExist
		}
		return 0, 0, 0, err
	}

	if err := rows.Err(); err != nil {
		return 0, 0, 0, err
	}

	return followers, followings, posts, nil
}
