package database

func (db *appdbimpl) CreateComment(photoId uint64, comment Comment) (Comment, error) {
	res, err := db.c.Exec(`INSERT INTO comments (id, photo_id, commenter_id, commenter) VALUES (?, ?, ?, ?)`,
		comment.ID, photoId, comment.Text, comment.Commenter)
	if err != nil {
		return comment, err
	}

	lastInsertID, err := res.LastInsertId()

	if err != nil {
		return comment, err
	}

	comment.ID = uint64(lastInsertID)
	return comment, nil
}
