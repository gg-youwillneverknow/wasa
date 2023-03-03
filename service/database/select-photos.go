package database

func (db *appdbimpl) SelectPhotos(username string, page uint64, limit uint64, sort string) ([]Photo, error) {

	const query = `
	DECLARE @PageNumber AS INT
	DECLARE @RowsOfPage AS INT
	SET @PageNumber=?
	SET @RowsOfPage=?
	SELECT photos.id, photos.datetime, photos.user_id, photos.comments_num, photos.likes_num FROM photos INNER JOIN users ON photos.user_id=users.id WHERE users.username=?
	ORDER BY photos.datetime DESC
	OFFSET (@PageNumber-1)*@RowsOfPage ROWS
	FETCH NEXT @RowsOfPage ROWS ONLY`
	var ret []Photo

	// Issue the query, using the bounding box as filter
	rows, err := db.c.Query(query, page, limit, username)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Read all fountains in the resultset
	for rows.Next() {
		var p Photo
		err = rows.Scan(&p.ID, &p.Datetime, &p.UserID, &p.NumComments, &p.NumLikes)
		if err != nil {
			return nil, err
		}
		ret = append(ret, p)
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
