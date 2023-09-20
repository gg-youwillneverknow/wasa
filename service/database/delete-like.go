package database

func (db *appdbimpl) DeleteLike(photoId uint64, likerId uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE liker_id=? AND photo_id=?`, likerId, photoId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		return ErrLikeDoesNotExist
	}
	return nil
}
