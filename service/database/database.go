/*
Package database is the middleware between the app database and the code. All data (de)serialization (save/load) from a
persistent database are handled here. Database specific logic should never escape this package.

To use this package you need to apply migrations to the database if needed/wanted, connect to it (using the database
data source name from config), and then initialize an instance of AppDatabase from the DB connection.

For example, this code adds a parameter in `webapi` executable for the database data source name (add it to the
main.WebAPIConfiguration structure):

	DB struct {
		Filename string `conf:""`
	}

This is an example on how to migrate the DB and connect to it:

	// Start Database
	logger.Println("initializing database support")
	db, err := sql.Open("sqlite3", "./foo.db")
	if err != nil {
		logger.WithError(err).Error("error opening SQLite DB")
		return fmt.Errorf("opening SQLite: %w", err)
	}
	defer func() {
		logger.Debug("database stopping")
		_ = db.Close()
	}()

Then you can initialize the AppDatabase and pass it to the api package.
*/
package database

import (
	"database/sql"
	"errors"
	"fmt"
)

// Fountain struct represent a fountain in every API call between this package and the outside world.
// Note that the internal representation of fountain in the database might be different.
type Comment struct {
	ID        uint64
	Text      string
	Commenter string
}

var ErrCommentDoesNotExist = errors.New("comment does not exist")

type Like struct {
	Liker string
}

var ErrLikeAlreadyExist = errors.New("user already put like")
var ErrLikeDoesNotExist = errors.New("like does not exist")

type Photo struct {
	ID          uint64
	Datetime    string
	UserID      uint64
	Owner 		string
	NumLikes    uint32
	NumComments uint32
	Image       []byte
}

var ErrPhotoDoesNotExist = errors.New("photo does not exist")

type User struct {
	ID       uint64
	Username string
}

var ErrUserDoesNotExist = errors.New("user does not exist")

type Following struct {
	Username string
}

var ErrFollowingAlreadyExist = errors.New("following already exist")
var ErrFollowingDoesNotExist = errors.New("following does not exist")

type Follower struct {
	Username string
}

type Ban struct {
	Username string
}

var ErrBanAlreadyExist = errors.New("ban does not exist")
var ErrBanDoesNotExist = errors.New("ban does not exist")

// AppDatabase is the high level interface for the DB
type AppDatabase interface {
	// ListFountains returns the list of fountains with their status
	UpdateAccount(updatedInfo User) (User, error)

	SelectProfile(username string) (uint32, uint32, uint32, error)

	SelectPhotos(username string, page uint64, limit uint64) ([]Photo, error)
	
	SelectPhoto(photoId uint64) (Photo, error)

	CreatePhoto(photo Photo) (Photo, error)

	SelectPhotosForStream(username string, page uint64, limit uint64) ([]Photo, error)

	DeletePhoto(photoId uint64) error

	SelectImage(photoId uint64) ([]byte, error)

	SelectComments(photoId uint64, page uint64, limit uint64) ([]Comment, error)

	CreateComment(photoId uint64, comment Comment) (Comment, error)

	DeleteComment(commentId uint64) error

	UpdateLike(photoId uint64, likerId uint64) error

	DeleteLike(photoId uint64, likerId uint64) error

	SelectLikes(photoId uint64, page uint64, limit uint64) ([]Like, error)

	SelectFollowers(username string, page uint64, limit uint64) ([]Follower, error)

	SelectFollowings(username string, page uint64, limit uint64) ([]Following, error)

	UpdateFollowings(username string, followingusername string) error

	DeleteFollowing(username string, followingusername string) error

	SelectBans(username string, page uint64, limit uint64) ([]Ban, error)

	UpdateBan(username string, bannedusername string) error

	DeleteBan(username string, bannedusername string) error

	CreateUser(username string) (uint64, error)

	SelectUser(username string) (uint64, error)
	// Ping checks whether the database is available or not (in that case, an error will be returned)
	Ping() error
}

type appdbimpl struct {
	c *sql.DB
}

// New returns a new instance of AppDatabase based on the SQLite connection `db`.
// `db` is required - an error will be returned if `db` is `nil`.
func New(db *sql.DB) (AppDatabase, error) {
	if db == nil {
		return nil, errors.New("database is required when building a AppDatabase")
	}
	var err error
	_,_=db.Exec("PRAGMA foreign_keys = ON;")
	// Check if table exists. If not, the database is empty, and we need to create the structure
	var tableName string
	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='users';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE users (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	username TEXT NOT NULL UNIQUE,
		followers_num INTEGER NOT NULL,
		posts_num INTEGER NOT NULL,
		followings_num INTEGER NOT NULL
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='photos';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE photos (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
    	datetime TEXT NOT NULL,
		user_id INTEGER,
		likes_num INTEGER NOT NULL,
		comments_num INTEGER NOT NULL,
		image BLOB NOT NULL, 
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos: %w", err)
		}

		sqlStmt2 := `CREATE TRIGGER increase_posts AFTER INSERT ON photos FOR EACH ROW BEGIN UPDATE users SET posts_num = posts_num + 1 WHERE id=NEW.user_id; END;`

		_, err = db.Exec(sqlStmt2)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users trigger 1: %w", err)
		}

		sqlStmt3 := `CREATE TRIGGER decrease_posts AFTER DELETE ON photos FOR EACH ROW BEGIN UPDATE users SET posts_num = posts_num - 1 WHERE id=OLD.user_id; END;`

		_, err = db.Exec(sqlStmt3)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users trigger2: %w", err)
		}
	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='comments';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE comments (
    	id INTEGER PRIMARY KEY AUTOINCREMENT,
		commenter_id INTEGER,
		comment TEXT NOT NULL,
		photo_id INTEGER,
		FOREIGN KEY (commenter_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure comments: %w", err)
		}

		sqlStmt2 := `CREATE TRIGGER increase_comments AFTER INSERT ON comments FOR EACH ROW 
					BEGIN 
						UPDATE photos SET comments_num = comments_num + 1 WHERE photos.id=NEW.photo_id;
					END;`

		_, err = db.Exec(sqlStmt2)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos trigger1: %w", err)
		}

		sqlStmt3 := `CREATE TRIGGER decrease_comments AFTER DELETE ON comments FOR EACH ROW 
					BEGIN 
						UPDATE photos SET comments_num = comments_num - 1 WHERE photos.id=OLD.photo_id;
					END;`

		_, err = db.Exec(sqlStmt3)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos trigger2: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='likes';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE likes (
		liker_id INTEGER,
		photo_id INTEGER,
		FOREIGN KEY (photo_id) REFERENCES photos(id) ON DELETE CASCADE,
		FOREIGN KEY (liker_id) REFERENCES users(id) ON DELETE CASCADE,
		PRIMARY KEY (liker_id,photo_id)
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure likes: %w", err)
		}

		sqlStmt2 := `CREATE TRIGGER increase_likes AFTER INSERT ON likes FOR EACH ROW 
					BEGIN 
						UPDATE photos SET likes_num = likes_num + 1 WHERE id=NEW.photo_id;
					END;`

		_, err = db.Exec(sqlStmt2)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure photos trigger3: %w", err)
		}

		sqlStmt3 := `CREATE TRIGGER decrease_likes AFTER DELETE ON likes FOR EACH ROW 
		BEGIN 
			UPDATE photos SET likes_num = likes_num - 1 WHERE id=OLD.photo_id;
		END;`

		_, err = db.Exec(sqlStmt3)
		if err != nil {
		return nil, fmt.Errorf("error creating database structure photos trigger4: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='followers';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE followers (
		follower_id INTEGER,
		user_id INTEGER,
		PRIMARY KEY (follower_id,user_id),
		FOREIGN KEY (follower_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure followers: %w", err)
		}

		sqlStmt2 := `CREATE TRIGGER increase_followers AFTER INSERT ON followers FOR EACH ROW 
					BEGIN 
						UPDATE users SET followers_num = followers_num + 1 WHERE id=NEW.user_id;
					END;`

		_, err = db.Exec(sqlStmt2)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users trigger3: %w", err)
		}

		sqlStmt3 := `CREATE TRIGGER decrease_followers AFTER DELETE ON followers FOR EACH ROW 
		BEGIN 
			UPDATE users SET followers_num = followers_num - 1 WHERE id=OLD.user_id;
		END;`

		_, err = db.Exec(sqlStmt3)
		if err != nil {
		return nil, fmt.Errorf("error creating database structure users trigger4: %w", err)
		}
		
		sqlStmt4 := `CREATE TRIGGER increase_followings AFTER INSERT ON followers FOR EACH ROW 
					BEGIN 
						UPDATE users SET followings_num = followings_num + 1 WHERE id=NEW.follower_id;
					END;`

		_, err = db.Exec(sqlStmt4)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure users trigger5: %w", err)
		}

		sqlStmt5 := `CREATE TRIGGER decrease_followings AFTER DELETE ON followers FOR EACH ROW 
		BEGIN 
			UPDATE users SET followings_num = followings_num - 1 WHERE id=OLD.follower_id;
		END;`

		_, err = db.Exec(sqlStmt5)
		if err != nil {
		return nil, fmt.Errorf("error creating database structure users trigger6: %w", err)
		}

	}

	err = db.QueryRow(`SELECT name FROM sqlite_master WHERE type='table' AND name='bans';`).Scan(&tableName)
	if errors.Is(err, sql.ErrNoRows) {
		sqlStmt := `CREATE TABLE bans (
		banned_id INTEGER,
		user_id INTEGER,
		PRIMARY KEY (banned_id,user_id)
		FOREIGN KEY (banned_id) REFERENCES users(id) ON DELETE CASCADE,
		FOREIGN KEY (user_id) REFERENCES users(id) ON DELETE CASCADE
		);`
		_, err = db.Exec(sqlStmt)
		if err != nil {
			return nil, fmt.Errorf("error creating database structure bans: %w", err)
		}
	}

	return &appdbimpl{c: db}, nil
}

func (db *appdbimpl) Ping() error {
	return db.c.Ping()
}
