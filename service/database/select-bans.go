package database

import "database/sql"

func (db *appdbimpl) SelectBans(username string, page uint64, limit uint64) ([]Ban, error) {
	var userId uint64
	var offset = (page - 1) * limit
	var ret []Ban

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)
	if err := row.Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			return ret, ErrUserDoesNotExist
		}
		return ret, err
	}
	if err2 := row.Err(); err2 != nil {
		return ret, err2
	}

	const query = `SELECT users.username FROM users INNER JOIN bans ON users.id=bans.banned_id 
	WHERE bans.user_id=?
	LIMIT ? 
	OFFSET ?`
	rows2, err3 := db.c.Query(query, userId, limit, offset)
	if err3 != nil {
		return ret, err3
	}
	defer func() { _ = rows2.Close() }()

	for rows2.Next() {
		var b Ban
		err3 = rows2.Scan(&b.Username)
		if err3 != nil {
			return ret, err3
		}
		ret = append(ret, b)
	}
	if err4 := rows2.Err(); err4 != nil {
		return ret, err4
	}
	return ret, nil

}
