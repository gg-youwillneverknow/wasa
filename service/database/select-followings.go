package database

func (db *appdbimpl) SelectFollowings(username string, page uint64, limit uint64) ([]Following, error) {

	// Here we need to get all fountains inside a given range. One simple solution is to rely on GIS/Spatial functions
	// from the DB itself. GIS/Spatial functions are those dedicated to geometry/geography/space computation.
	//
	// However, some databases (like SQLite) do not support these functions. So, we use a naive approach: instead of
	// drawing a circle for a given range, we get slightly more fountains by retrieving a square area, and then we will
	// filter the result later ("cutting the corner").
	//
	// Steps are:
	// 1. We compute a square ("bounding box") that contains the circle. The square will have edges with the same length
	//    of the range of the circle.
	// 2. For each resulting fountain, we will check (using Go and some math) if it's inside the range or not.

	rows, err := db.c.Query(`SELECT id FROM users WHERE username=?`, username)
	if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	var followerId uint64
	err = rows.Scan(&followerId)
	if err != nil {
		return nil, err
	}

	if rows.Err() != nil {
		return nil, err
	}

	const query = `
	DECLARE @PageNumber AS INT
	DECLARE @RowsOfPage AS INT
	SET @PageNumber=?
	SET @RowsOfPage=?
	SELECT user_id FROM followers
	WHERE follower_id=?
	OFFSET (@PageNumber-1)*@RowsOfPage ROWS
	FETCH NEXT @RowsOfPage ROWS ONLY`

	var ret []Following

	// Issue the query, using the bounding box as filter
	rows2, err2 := db.c.Query(query, page, limit, followerId)
	if err2 != nil {
		return nil, err2
	}
	defer func() { _ = rows2.Close() }()

	// Read all fountains in the resultset
	for rows2.Next() {
		var f Following
		var followingId uint64
		err = rows2.Scan(followingId)
		if err != nil {
			return nil, err
		}
		rows3, err3 := db.c.Query(`SELECT username FROM users WHERE id=?`, followingId)
		if err3 != nil {
			return nil, err3
		}
		defer func() { _ = rows3.Close() }()

		err3 = rows3.Scan(&f.Following)
		if err3 != nil {
			return nil, err3
		}

		if rows3.Err() != nil {
			return nil, err3
		}
		ret = append(ret, f)

	}
	if rows2.Err() != nil {
		return nil, err2
	}

	return ret, nil
}
