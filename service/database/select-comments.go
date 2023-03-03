package database

func (db *appdbimpl) SelectComments(photoId uint64, page uint64, limit uint64) ([]Comment, error) {

	const query = `
	DECLARE @PageNumber AS INT
	DECLARE @RowsOfPage AS INT
	SET @PageNumber=?
	SET @RowsOfPage=?
	SELECT comments.id, comments.comment, users.username FROM comments INNER JOIN users ON comments.commenter_id=users.id WHERE comments.photo_id=?
	OFFSET (@PageNumber-1)*@RowsOfPage ROWS
	FETCH NEXT @RowsOfPage ROWS ONLY`

	var ret []Comment

	rows, err := db.c.Query(query, page, limit, photoId)
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
