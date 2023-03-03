package database

func (db *appdbimpl) SelectLikes(photoId uint64, page uint64, limit uint64) ([]Like, error) {

	const query = `
	DECLARE @PageNumber AS INT
	DECLARE @RowsOfPage AS INT
	SET @PageNumber=?
	SET @RowsOfPage=?
	SELECT users.username FROM users INNER JOIN likes ON users.id=likes.liker_id WHERE likes.photo_id=?
	OFFSET (@PageNumber-1)*@RowsOfPage ROWS
	FETCH NEXT @RowsOfPage ROWS ONLY`

	var ret []Like

	rows, err := db.c.Query(query, page, limit, photoId)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	for rows.Next() {
		var c Like
		err = rows.Scan(&c.Liker)
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
