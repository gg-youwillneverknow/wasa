package database
import "database/sql"

func (db *appdbimpl) SelectImage(photoId uint64) ([]byte, error) {
	var ret []byte

	row := db.c.QueryRow(`SELECT image FROM photos WHERE id=?`, photoId)

	if err := row.Scan(&ret); err != nil {
		if err == sql.ErrNoRows {
			return ret, ErrPhotoDoesNotExist
		}
		return ret, err
	}

	if err = row.Err(); err!= nil {
		return ret, err
	}

	return ret, nil
}
