package database

func (db *appdbimpl) UpdateAccount(updatedInfo User) (User, error) {

	res, err := db.c.Exec(`UPDATE users SET username=? WHERE id=?`, updatedInfo.Username, updatedInfo.ID)
	if err != nil {
		return updatedInfo, err
	}

	affected, err := res.RowsAffected()

	if err != nil {
		return updatedInfo, err
	} else if affected == 0 {
		// If we didn't update any row, then the user didn't exist
		return updatedInfo, ErrUserDoesNotExist
	}
	return updatedInfo, nil
}
