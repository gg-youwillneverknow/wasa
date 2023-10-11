package database

func (db *appdbimpl) DeleteComment(commentId uint64) error {
	res, err := db.c.Exec(`DELETE FROM comments WHERE id=?`, commentId)
	if err != nil {
		return err
	}

	affected, err2 := res.RowsAffected()
	if err2 != nil {
		return err2
	} else if affected == 0 {
		return ErrCommentDoesNotExist
	}
	return nil
}
