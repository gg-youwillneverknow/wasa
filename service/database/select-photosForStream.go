package database

func (db *appdbimpl) SelectPhotosForStream(username string, page uint64, limit uint64) ([]Photo, error) {
	const query = `
	DECLARE @PageNumber AS INT
	DECLARE @RowsOfPage AS INT
	SET @PageNumber=?
	SET @RowsOfPage=?
	SELECT photos.id, photos.user_id, photos.datetime, photos.comments_num, photos.likes_num FROM photos WHERE photos.user_id IN (SELECT followers.user_id INNER JOIN users ON followers.follower_id=users.id WHERE users.username=?)  
	ORDER BY photos.datetime
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
		err = rows.Scan(&p.ID, &p.UserID, &p.Datetime, &p.NumComments, &p.NumLikes)
		if err != nil {
			return nil, err
		}
	}
	if rows.Err() != nil {
		return nil, err
	}

	return ret, nil
}
