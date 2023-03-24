package api

import "git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"

// Fountain struct represent a fountain in every data exchange with the external world via REST API. JSON tags have been
// added to the struct to conform to the OpenAPI specifications regarding JSON key names.
// Note: there is a similar struct in the database package. See Fountain.FromDatabase (below) to understand why.

type User struct {
	ID       uint64 `json:"userId"`
	Username string `json:"username"`
}

// ToDatabase returns the user in a database-compatible representation
func (f *User) ToDatabase() database.User {
	return database.User{
		ID:       f.ID,
		Username: f.Username,
	}
}

func (f *User) FromDatabase(fountain database.User) {
	f.ID = fountain.ID
	f.Username = fountain.Username
}

type Profile struct {
	Username   string  `json:"username"`
	Followers  uint32  `json:"followers"`
	Followings uint32  `json:"followings"`
	Posts      uint32  `json:"posts"`
	Photos     []Photo `json:"photos"`
}

type Photo struct {
	ID          uint64    `json:"ID"`
	Datetime    string    `json:"dateTime"`
	Likes       []Like    `json:"likes"`
	Comments    []Comment `json:"comments"`
	NumLikes    uint32    `json:"likesnum"`
	NumComments uint32    `json:"commentsnum"`
}

func (f *Photo) ToDatabase(img []byte) database.Photo {
	return database.Photo{
		ID:          f.ID,
		Datetime:    f.Datetime,
		UserID:      0,
		NumLikes:    0,
		NumComments: 0,
		Image:       img,
	}
}

func (f *Photo) FromDatabase(fountain database.Photo) {
	f.ID = fountain.ID
	f.Datetime = fountain.Datetime
	f.Likes = nil
	f.Comments = nil
	f.NumLikes = fountain.NumLikes
	f.NumComments = fountain.NumLikes
}

type Image struct {
	ID    uint64
	Image []byte
}

func (f *Image) ToDatabase() database.Image {
	return database.Image{
		ID:    f.ID,
		Image: f.Image,
	}
}

func (f *Image) FromDatabase(fountain database.Image) {
	f.ID = fountain.ID
	f.Image = fountain.Image
}

type Comment struct {
	ID        uint64 `json:"ID"`
	Text      string `json:"text"`
	Commenter string `json:"commenter"`
}

func (f *Comment) ToDatabase() database.Comment {
	return database.Comment{
		ID:        f.ID,
		Text:      f.Text,
		Commenter: f.Commenter,
	}
}

func (f *Comment) FromDatabase(fountain database.Comment) {
	f.ID = fountain.ID
	f.Text = fountain.Text
	f.Commenter = fountain.Commenter
}

func (f *Comment) IsValid() bool {
	return 0 <= len(f.Text) && len(f.Text) <= 90
}

type Like struct {
	Liker string `json:"liker"`
}

func (f *Like) ToDatabase() database.Like {
	return database.Like{
		Liker: f.Liker,
	}
}

func (f *Like) FromDatabase(fountain database.Like) {
	f.Liker = fountain.Liker
}

// FromDatabase populates the struct with data from the database, overwriting all values.
// You might think this is code duplication, which is correct. However, it's "good" code duplication because it allows
// us to uncouple the database and API packages.
// Suppose we were using the "database.Fountain" struct inside the API package; in that case, we were forced to conform
// either the API specifications to the database package or the other way around. However, very often, the database
// structure is different from the structure of the REST API.
// Also, in this way the database package is freely usable by other packages without the assumption that structs from
// the database should somehow be JSON-serializable (or, in general, serializable).
