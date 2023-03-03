package database

func (db *appdbimpl) CreatePhoto(username string, photo Photo) (Photo, error) {
	rows, err := db.c.Query(`SELECT id FROM users WHERE username=?`, username)
	if err != nil {
		return photo, err
	}
	defer func() { _ = rows.Close() }()

	// Here we read the resultset and we build the list to be returned
	var userId uint64
	err = rows.Scan(&userId)
	if err != nil {
		return photo, err
	}

	if rows.Err() != nil {
		return photo, err
	}
	photo.UserID = userId
	res, err := db.c.Exec(`INSERT INTO photos (id, datetime, user_id, likes_num, comments_num, image) VALUES (?, NOW(), ?, 0, 0,?)`,
		photo.ID, photo.UserID, photo.Image)
	if err != nil {
		return photo, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return photo, err
	}
	photo.ID = uint64(lastInsertID)

	row, err := db.c.Query(`SELECT datetime FROM photos WHERE id=?`, photo.ID)
	if err != nil {
		return photo, err
	}
	defer func() { _ = row.Close() }()

	// Here we read the resultset and we build the list to be returned

	err = row.Scan(&photo.Datetime)
	if err != nil {
		return photo, err
	}

	if row.Err() != nil {
		return photo, err
	}
	return photo, nil
}
