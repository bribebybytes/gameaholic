package models

import (
	"encoding/json"
	"github.com/gobuffalo/pop"
	"github.com/gobuffalo/validate"
	"github.com/gofrs/uuid"
	"time"
	"github.com/gobuffalo/validate/validators"
)
// PlayerTimeline is used by pop to map your .model.Name.Proper.Pluralize.Underscore database table to your go code.
type PlayerTimeline struct {
    ID uuid.UUID `json:"id" db:"id"`
    PlayerID int `json:"player_id" db:"player_id"`
    Event string `json:"event" db:"event"`
    EventProvider string `json:"event_provider" db:"event_provider"`
    EventURL string `json:"event_url" db:"event_url"`
    CreationMonthYear string `json:"creation_month_year" db:"creation_month_year"`
    CreatedAt time.Time `json:"created_at" db:"created_at"`
    UpdatedAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (p PlayerTimeline) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// PlayerTimelines is not required by pop and may be deleted
type PlayerTimelines []PlayerTimeline

// String is not required by pop and may be deleted
func (p PlayerTimelines) String() string {
	jp, _ := json.Marshal(p)
	return string(jp)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (p *PlayerTimeline) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.Validate(
		&validators.IntIsPresent{Field: p.PlayerID, Name: "PlayerID"},
		&validators.StringIsPresent{Field: p.Event, Name: "Event"},
		&validators.StringIsPresent{Field: p.EventProvider, Name: "EventProvider"},
		&validators.StringIsPresent{Field: p.EventURL, Name: "EventURL"},
		&validators.StringIsPresent{Field: p.CreationMonthYear, Name: "CreationMonthYear"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (p *PlayerTimeline) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (p *PlayerTimeline) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}
