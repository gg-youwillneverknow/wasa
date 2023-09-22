package database

func (db *appdbimpl) CreatePhoto(photo Photo) (Photo, error) {
	res, err := db.c.Exec(`INSERT INTO photos (datetime, user_id, likes_num, comments_num, image) VALUES (datetime("now"), ?, 0, 0,?)`,
		photo.UserID, photo.Image)
	if err != nil {
		return photo, err
	}
	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return photo, err
	}
	photo.ID = uint64(lastInsertID)

	row := db.c.QueryRow(`SELECT datetime FROM photos WHERE id=?`, photo.ID)
	if err = row.Scan(&photo.Datetime); err != nil {
		return photo, err
	}

	if row.Err() != nil {
		return photo, err
	}
	return photo, nil
}
