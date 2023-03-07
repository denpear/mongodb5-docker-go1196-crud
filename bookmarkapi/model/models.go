package model

import (
	"time"

	"gopkg.in/mgo.v2/bson"
)

type (
	// User type represents the registered user.
	User struct {
		ID           bson.ObjectId `bson:"_id,omitempty" json:"id"`
		FirstName    string        `json:"firstname"`
		LastName     string        `json:"lastname"`
		Email        string        `json:"email"`
		HashPassword []byte        `json:"hashpassword,omitempty"`
	}
	// Bookmark type represents the metadata of a bookmark.
	Bookmark struct {
		ID          bson.ObjectId `bson:"_id,omitempty"`
		Name        string        `json:"name"`
		Description string        `json:"description"`
		Location    string        `json:"location"`
		Priority    int           `json:"priority"` // Priority (1 -5)
		CreatedBy   string        `json:"createdby"`
		CreatedOn   time.Time     `json:"createdon,omitempty"`
		Tags        []string      `json:"tags,omitempty"`
	}

	// Bookmark type represents the metadata of a bookmark.
	BookmarkPdb struct {
		UUID        string    `gorm:"column:uuid,omitempty"`
		Name        string    `gorm:"column:name"`
		Description string    `gorm:"column:description"`
		Location    string    `gorm:"column:location"`
		Priority    int       `gorm:"column:priority"` // Priority (1 -5)
		CreatedBy   string    `gorm:"column:createdby"`
		CreatedOn   time.Time `gorm:"column:createdon,omitempty"`
		//Tags        []string  `gorm:"column:tags,omitempty"`
	}
)

//func (BookmarkPdb) TableName() string {
//	return "cms.applications"
//}
