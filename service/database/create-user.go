package database

func (db *appdbimpl) CreateUser(username string) (uint64, error) {
	res, err := db.c.Exec(`INSERT INTO users (id,username,followers_num,posts_num,followings_num) VALUES (NULL, ?,0,0,0)`, username)
	if err != nil {
		return 0, err
	}

	lastInsertID, err := res.LastInsertId()
	if err != nil {
		return 0, err
	}

	userId := uint64(lastInsertID)
	return userId, nil
}
