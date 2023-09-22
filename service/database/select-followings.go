package database

import "database/sql"

func (db *appdbimpl) SelectFollowings(username string, page uint64, limit uint64) ([]Following, error) {
	var ret []Following
	var offset = (page - 1) * limit
	var userId uint64

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

	const query = `
	SELECT users.username FROM users INNER JOIN followers ON users.id=followers.user_id 
	WHERE followers.follower_id=?
	LIMIT ?
	OFFSET ?`
	rows2, err2 := db.c.Query(query, userId, limit, offset)
	if err2 != nil {
		return nil, err2
	}
	defer func() { _ = rows2.Close() }()

	for rows2.Next() {
		var f Following
		err2 = rows2.Scan(&f.Username)
		if err2 != nil {
			return nil, err2
		}
		ret = append(ret, f)
	}

	if err2 := rows2.Err(); err2 != nil {
		return nil, err2
	}

	return ret, nil
}
