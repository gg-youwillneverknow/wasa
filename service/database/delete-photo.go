package database

func (db *appdbimpl) DeletePhoto(photoId uint64) error {
	res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, photoId)
	if err != nil {
		return err
	}

	affected, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	} else if affected == 0 {
		// If we didn't delete any row, then the phot0 didn't exist
		return ErrPhotoDoesNotExist
	}
	return nil
}
