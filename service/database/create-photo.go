package database

func (db *appdbimpl) CreatePhoto(photo Photo) (Photo, error) {
	res, err := db.c.Exec(`INSERT INTO photos (datetime, user_id, likes_num, comments_num, image) VALUES (datetime("now"), ?, 0, 0,?)`,
		photo.UserID, photo.Image)
	if err != nil {
		return photo, err
	}
	lastInsertID, err2 := res.LastInsertId()
	if err2 != nil {
		return photo, err2
	}
	photo.ID = uint64(lastInsertID)

	row := db.c.QueryRow(`SELECT datetime FROM photos WHERE id=?`, photo.ID)
	if err3 := row.Scan(&photo.Datetime); err3 != nil {
		return photo, err3
	}

	if err4 := row.Err(); err4 != nil {
		return photo, err4
	}
	return photo, nil
}
