package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Achievement is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Achievement struct {
    ID uuid.UUID `json:"id" db:"id"`
    Aid int `json:"aid" db:"aid"`
    Aname string `json:"aname" db:"aname"`
    Alevel int `json:"alevel" db:"alevel"`
    Aavatar string `json:"aavatar" db:"aavatar"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (a Achievement) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Achievements is not required by pop and may be deleted
type Achievements []Achievement

// String is not required by pop and may be deleted
func (a Achievements) String() string {
	ja, _ := json.Marshal(a)
	return string(ja)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (a *Achievement) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: a.Aid, Name: "Aid"},
		&validators.StringIsPresent{Field: a.Aname, Name: "Aname"},
		&validators.IntIsPresent{Field: a.Alevel, Name: "Alevel"},
		&validators.StringIsPresent{Field: a.Aavatar, Name: "Aavatar"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (a *Achievement) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (a *Achievement) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
