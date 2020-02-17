package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// Level is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type Level struct {
    ID uuid.UUID `json:"id" db:"id"`
    LID int `json:"l_id" db:"l_id"`
    Level int `json:"level" db:"level"`
    Score int `json:"score" db:"score"`
    Aavatar string `json:"aavatar" db:"aavatar"`
    LevelName string `json:"level_name" db:"level_name"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (l Level) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Levels is not required by pop and may be deleted
type Levels []Level

// String is not required by pop and may be deleted
func (l Levels) String() string {
	jl, _ := json.Marshal(l)
	return string(jl)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (l *Level) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: l.LID, Name: "LID"},
		&validators.IntIsPresent{Field: l.Level, Name: "Level"},
		&validators.IntIsPresent{Field: l.Score, Name: "Score"},
		&validators.StringIsPresent{Field: l.Aavatar, Name: "Aavatar"},
		&validators.StringIsPresent{Field: l.LevelName, Name: "LevelName"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (l *Level) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (l *Level) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
