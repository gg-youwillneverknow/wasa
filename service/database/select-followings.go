package database

import "database/sql"

func (db *appdbimpl) SelectFollowings(username string, page uint64, limit uint64) ([]Following, error) {
	var ret []Following
	var offset = (page - 1) * limit
	var userId uint64

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
	SELECT users.username FROM users INNER JOIN followers ON users.id=followers.user_id 
	WHERE followers.follower_id=?
	LIMIT ?
	OFFSET ?`
	rows, err3 := db.c.Query(query, userId, limit, offset)
	if err3 != nil {
		return ret, err3
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var f Following
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
