package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// FavoriteFood is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type FavoriteFood struct {
    ID uuid.UUID `json:"id" db:"id"`
    User uuid.UUID `json:"user" db:"user"`
    Food string `json:"food" db:"food"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (f FavoriteFood) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// FavoriteFoods is not required by pop and may be deleted
type FavoriteFoods []FavoriteFood

// String is not required by pop and may be deleted
func (f FavoriteFoods) String() string {
	jf, _ := json.Marshal(f)
	return string(jf)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (f *FavoriteFood) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: f.Food, Name: "Food"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (f *FavoriteFood) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (f *FavoriteFood) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
