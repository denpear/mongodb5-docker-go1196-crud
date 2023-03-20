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
	Applicant struct {
		ID                  int                 `gorm:"primaryKey"`
		UUID                string              `gorm:"column:uuid"`
		LastName            string              `gorm:"column:lastName"`
		FirstName           string              `gorm:"column:firstName"`
		MiddleName          string              `gorm:"column:middleName"`
		ChangeNameFlag      string              `gorm:"column:changeNameFlag"`
		BirthPlace          string              `gorm:"column:birthPlace"`
		BirthDate           string              `gorm:"column:birthDate"`
		SexCode             string              `gorm:"column:sexCode"`
		PassportSeries      string              `gorm:"column:passportSeries"`
		PassportNumber      string              `gorm:"column:passportNumber"`
		CreateDate          string              `gorm:"column:createDate"`
		UpdateDate          string              `gorm:"column:updateDate"`
		CreatedBy           string              `gorm:"column:createdby"`
		ApplicationDateTime time.Time           `gorm:"column:applicationDateTime"`
		AdditionalOptions   []AdditionalOptions //`json:"AdditionalOptions" gorm:"foreignKey:applicant id"`
	}
)

type AdditionalOptions struct {
	ID          int     `json:"id"`
	UUID        string  `gorm:"column:uuid"`
	Type        string  `gorm:"column:type"`
	Name        string  `gorm:"column:name"`
	Vendor      string  `gorm:"column:vendor"`
	Price       float32 `gorm:"column:price"`
	Term        int     `gorm:"column:term"`
	ApplicantID int     `gorm:"column:applicant_id"`
}

//func (Applicant) TableName() string {
//	return "cms.applications"
//}
