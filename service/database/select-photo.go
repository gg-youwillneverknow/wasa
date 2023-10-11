package database

import "database/sql"

func (db *appdbimpl) SelectPhoto(photoId uint64) (Photo, error) {
	var p Photo

	const query = `
	SELECT photos.id, photos.datetime, photos.user_id, photos.comments_num, photos.likes_num 
	FROM photos WHERE photos.id=?`
	row := db.c.QueryRow(query, photoId)
	if err := row.Scan(&p.ID, &p.Datetime, &p.UserID, &p.NumComments, &p.NumLikes); err != nil {
		if err == sql.ErrNoRows {
			return p, ErrPhotoDoesNotExist
		}
		return p, err
	}

	if err2 := row.Err(); err2 != nil {
		return p, err2
	}

	return p, nil

}
