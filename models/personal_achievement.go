package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// PersonalAchievement is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type PersonalAchievement struct {
    ID uuid.UUID `json:"id" db:"id"`
    PaID int `json:"pa_id" db:"pa_id"`
    PlayerID int `json:"player_id" db:"player_id"`
    AchievID int `json:"achiev_id" db:"achiev_id"`
    AchievDate time.Time `json:"achiev_date" db:"achiev_date"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p PersonalAchievement) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PersonalAchievements is not required by pop and may be deleted
type PersonalAchievements []PersonalAchievement

// String is not required by pop and may be deleted
func (p PersonalAchievements) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PersonalAchievement) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.PaID, Name: "PaID"},
		&validators.IntIsPresent{Field: p.PlayerID, Name: "PlayerID"},
		&validators.IntIsPresent{Field: p.AchievID, Name: "AchievID"},
		&validators.TimeIsPresent{Field: p.AchievDate, Name: "AchievDate"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PersonalAchievement) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PersonalAchievement) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
