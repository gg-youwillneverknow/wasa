package database
import "database/sql"

func (db *appdbimpl) SelectUser(username string) (uint64, error) {
	var ret uint64

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)
	if err := row.Scan(&ret); err != nil {
		if err == sql.ErrNoRows {
			return ret, ErrUserDoesNotExist
		}
		return ret, err
	}

	if err := row.Err(); err!= nil {
		return ret, err
	}

	return ret, nil
	
}
