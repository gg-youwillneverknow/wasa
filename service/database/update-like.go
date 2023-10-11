package database

func (db *appdbimpl) UpdateLike(photoId uint64, likerId uint64) error {
	res, err := db.c.Exec(`INSERT INTO likes (liker_id, photo_id) VALUES (?, ?)`, likerId, photoId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {

		return ErrLikeAlreadyExist
	}
	return nil
}
