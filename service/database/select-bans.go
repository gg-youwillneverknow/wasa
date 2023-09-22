package database

import "database/sql"

func (db *appdbimpl) SelectBans(username string, page uint64, limit uint64) ([]Ban, error) {
	var userId uint64
	var offset = (page - 1) * limit
	var ret []Ban

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)
	if err := row.Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			return nil, ErrUserDoesNotExist
		}
		return nil, err
	}
	if err := row.Err(); err != nil {
		return nil, err
	}

	const query = `SELECT users.username FROM users INNER JOIN bans ON users.id=bans.banned_id 
	WHERE bans.user_id=?
	LIMIT ? 
	OFFSET ?`
	rows2, err2 := db.c.Query(query, userId, limit, offset)
	if err2 != nil {
		return nil, err2
	}
	defer func() { _ = rows2.Close() }()

	for rows2.Next() {
		var b Ban
		err2 = rows2.Scan(&b.Username)
		if err2 != nil {
			return nil, err2
		}
		ret = append(ret, b)
	}
	if rows2.Err() != nil {
		return nil, err2
	}
	return ret, nil

}
