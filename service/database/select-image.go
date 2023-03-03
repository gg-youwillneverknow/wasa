package database

func (db *appdbimpl) SelectImage(photoId uint64) (Image, error) {
	var ret Image

	// Plain simple SELECT query
	row, err := db.c.Query(`SELECT photo_id, image FROM photos WHERE photo_id=?`, photoId)
	if err != nil {
		return ret, err
	}
	defer func() { _ = row.Close() }()

	// Here we read the resultset and we build the list to be returned

	err = row.Scan(&ret.ID, ret.Image)
	if err != nil {
		return ret, err
	}

	if row.Err() != nil {
		return ret, err
	}

	return ret, nil
}
