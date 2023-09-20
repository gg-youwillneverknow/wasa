package database
import "database/sql"

func (db *appdbimpl) UpdateFollowings(username string, followingusername string) error {
	var userId uint64
	var followingId uint64

	row := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, username)
	if err := row.Scan(&userId); err != nil {
		if err == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err
	}
	if err := row.Err(); err!= nil {
		return err
	}

	row2 := db.c.QueryRow(`SELECT id FROM users WHERE username=?`, followingusername)
	if err2 := row2.Scan(&followingId); err2 != nil {
		if err2 == sql.ErrNoRows {
			return ErrUserDoesNotExist
		}
		return err2
	}
	if err2 := row2.Err(); err2!= nil {
		return err2
	}

	res, err := db.c.Exec(`INSERT INTO followers (user_id, follower_id) VALUES (?, ?)`,
		followingId, userId)
	if err != nil {
		return err
	} 

	affected, err := res.RowsAffected()
	if err != nil {
		return err
	} 
	if affected == 0 {
		return ErrFollowingAlreadyExist
	}
	return nil
}
