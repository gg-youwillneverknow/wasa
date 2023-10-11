package database

func (db *appdbimpl) SelectPhotosForStream(username string, page uint64, limit uint64) ([]Photo, error) {
	var offset = (page - 1) * limit
	const query = `
	SELECT photos.id, photos.datetime, photos.comments_num, photos.likes_num, users.username 
	FROM photos INNER JOIN users ON photos.user_id=users.id 
	WHERE photos.user_id IN 
	(SELECT followers.user_id FROM followers INNER JOIN users ON followers.follower_id=users.id 
		WHERE users.username=?) AND photos.user_id NOT IN 
	(SELECT bans.user_id FROM bans INNER JOIN users ON bans.banned_id=users.id 
		WHERE users.username=?)
	ORDER BY photos.datetime DESC
	LIMIT ?
	OFFSET ?`
	var ret []Photo

	// Issue the query, using the bounding box as filter
	rows, err := db.c.Query(query, username, username, limit, offset)
	if err != nil {
		return ret, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.Datetime, &p.NumComments, &p.NumLikes, &p.Owner)
		if err != nil {
			return ret, err
		}
		ret = append(ret, p)
	}
	if err2 := rows.Err(); err2 != nil {
		return ret, err2
	}

	return ret, nil
}
