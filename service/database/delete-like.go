package database

func (db *appdbimpl) DeleteLike(photoId uint64, likerId uint64) error {
	res, err := db.c.Exec(`DELETE FROM likes WHERE user_id=? AND photo=?`, likerId, photoId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrLikeDoesNotExist
	}
	return nil
}
