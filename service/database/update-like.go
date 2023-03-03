package database

func (db *appdbimpl) UpdateLike(photoId uint64, likerId uint64) error {
	res, err := db.c.Exec(`INSERT INTO likes (liker_id, photo_id) VALUES (?, ?) WHERE NOT EXISTS (SELECT * FROM likes WHERE liker_id=? AND photo_id=?)`,
		likerId, photoId)
	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrLikeAlreadyExist
	}
	return nil
}
