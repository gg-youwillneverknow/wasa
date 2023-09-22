package database

func (db *appdbimpl) SelectPhotos(username string, page uint64, limit uint64) ([]Photo, error) {
	var offset = (page - 1) * limit
	var ret []Photo

	const query = `
	SELECT photos.id, photos.datetime, photos.comments_num, photos.likes_num FROM photos INNER JOIN users ON photos.user_id=users.id WHERE users.username=?
	ORDER BY photos.datetime DESC
	LIMIT ?
	OFFSET ?`

	rows, err := db.c.Query(query, username, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.Datetime, &p.NumComments, &p.NumLikes)
		if err != nil {
			return nil, err
		}
		p.Owner = username
		ret = append(ret, p)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
