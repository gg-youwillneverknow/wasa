package database

func (db *appdbimpl) SelectComments(photoId uint64, page uint64, limit uint64) ([]Comment, error) {
	var offset = (page - 1) * limit
	var ret []Comment

	const query = `
	SELECT comments.id, comments.comment, users.username FROM comments INNER JOIN users ON comments.commenter_id=users.id WHERE comments.photo_id=?
	LIMIT ?
	OFFSET ?`
	rows, err := db.c.Query(query, photoId, limit, offset)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var c Comment
		err = rows.Scan(&c.ID, &c.Text, &c.Commenter)
		if err != nil {
			return nil, err
		}
		ret = append(ret, c)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
