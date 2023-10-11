package database

import "database/sql"

func (db *appdbimpl) SelectFollowers(username string, page uint64, limit uint64) ([]Follower, error) {
	var userId uint64
	var ret []Follower
	var offset = (page - 1) * limit

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

	const query = `
	SELECT users.username FROM users INNER JOIN followers ON users.id=followers.follower_id 
	WHERE followers.user_id=?
	LIMIT ?
	OFFSET ?`
	rows, err3 := db.c.Query(query, userId, limit, offset)
	if err3 != nil {
		return ret, err3
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var f Follower
		err3 = rows.Scan(&f.Username)
		if err3 != nil {
			return ret, err3
		}
		ret = append(ret, f)
	}
	if err4 := rows.Err(); err4 != nil {
		return ret, err4
	}
	return ret, nil
}
