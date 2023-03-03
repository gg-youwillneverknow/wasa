package database

func (db *appdbimpl) DeleteComment(commentId uint64) error {
	res, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, commentId)
	if err != nil {
		return err
	}

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} else if affected == 0 {
		// If we didn't delete any row, then the fountain didn't exist
		return ErrCommentDoesNotExist
	}
	return nil
}
