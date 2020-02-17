package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/gofrs/uuid"
)

// Player is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Player struct {
	ID              uuid.UUID `json:"id" db:"id"`
	Pid             int       `json:"pid" db:"pid"`
	Pname           string    `json:"pname" db:"pname"`
	Pavatar         string    `json:"pavatar" db:"pavatar"`
	Pemailid        string    `json:"pemailid" db:"pemailid"`
	PcurAchive      string    `json:"pcur_achive" db:"pcur_achive"`
	Pscore          int       `json:"pscore" db:"pscore"`
	Plevel          int       `json:"plevel" db:"plevel"`
	BitbucketUname  string    `json:"bitbucket_uname" db:"bitbucket_uname"`
	JiraUname       string    `json:"jira_uname" db:"jira_uname"`
	ConfluenceUname string    `json:"confluence_uname" db:"confluence_uname"`
	CreatedAt       time.Time `json:"created_at" db:"created_at"`
	UpdatedAt       time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p Player) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Players is not required by pop and may be deleted
type Players []Player

// String is not required by pop and may be deleted
func (p Players) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *Player) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.Pid, Name: "Pid"},
		&validators.StringIsPresent{Field: p.Pname, Name: "Pname"},
		&validators.StringIsPresent{Field: p.Pavatar, Name: "Pavatar"},
		&validators.StringIsPresent{Field: p.Pemailid, Name: "Pemailid"},
		&validators.StringIsPresent{Field: p.PcurAchive, Name: "PcurAchive"},
		&validators.IntIsPresent{Field: p.Pscore, Name: "Pscore"},
		&validators.IntIsPresent{Field: p.Plevel, Name: "Plevel"},
		&validators.StringIsPresent{Field: p.BitbucketUname, Name: "BitbucketUname"},
		&validators.StringIsPresent{Field: p.JiraUname, Name: "JiraUname"},
		&validators.StringIsPresent{Field: p.ConfluenceUname, Name: "ConfluenceUname"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *Player) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *Player) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
