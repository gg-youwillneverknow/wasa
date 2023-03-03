package database

func (db *appdbimpl) DeletePhoto(photoId string) error {
	res, err := db.c.Exec(`DELETE FROM photos WHERE id=?`, photoId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrPhotoDoesNotExist
	}
	return nil
}
