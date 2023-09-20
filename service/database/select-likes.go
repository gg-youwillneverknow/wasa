package database

func (db *appdbimpl) SelectLikes(photoId uint64, page uint64, limit uint64) ([]Like, error) {
	var offset = (page -1)*limit
	var ret []Like

	const query = `
	SELECT users.username FROM users INNER JOIN likes ON users.id=likes.liker_id WHERE likes.photo_id=?
	LIMIT ?
	OFFSET ?`

	rows, err := db.c.Query(query, photoId, limit, offset)
	if err != nil {
		return ret, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var l Like
		err = rows.Scan(&l.Liker)
		if err != nil {
			return ret, err
		}
		ret = append(ret, l)
	}
	if rows.Err() != nil {
		return ret, err
	}

	return ret, nil
}
