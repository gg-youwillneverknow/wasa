package database

import "database/sql"

func (db *appdbimpl) CreateComment(photoId uint64, comment Comment) (Comment, error) {
	var id uint64
	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, comment.Commenter)
	// Here we read the resultset and we build the list to be returned
	if err := row.Scan(&id); err != nil {
		if err == sql.ErrNoRows {
			return comment, ErrUserDoesNotExist
		}
		return comment, err
	}

	if err2 := row.Err(); err2 != nil {
		return comment, err2
	}

	res, err3 := db.c.Exec(`INSERT INTO comments (id, photo_id, comment, commenter_id) VALUES (NULL, ?, ?, ?)`,
		photoId, comment.Text, id)
	if err3 != nil {
		return comment, err3
	}

	lastInsertID, err4 := res.LastInsertId()

	if err4 != nil {
		return comment, err4
	}

	comment.ID = uint64(lastInsertID)
	return comment, nil
}
