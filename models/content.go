package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/pop/v5"
	"github.com/gobuffalo/validate/v3"
	"github.com/gofrs/uuid"
)

// Content is used by pop to map your contents database table to your go code.
type Content struct {
	ID        int64     `json:"id" db:"id"`
	Tag       string    `json:"tag" db:"tag"`
	Type      string    `json:"type" db:"type"`
	Time      int64     `json:"time" db:"time"`
	UserID    uuid.UUID `json:"userId" db:"user_id"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
	UpdatetAt time.Time `json:"updated_at" db:"updated_at"`
}

// String is not required by pop and may be deleted
func (c Content) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Contents is not required by pop and may be deleted
type Contents []Content

// String is not required by pop and may be deleted
func (c Contents) String() string {
	jc, _ := json.Marshal(c)
	return string(jc)
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (c *Content) Validate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (c *Content) ValidateCreate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (c *Content) ValidateUpdate(tx *pop.Connection) (*validate.Errors, error) {
	return validate.NewErrors(), nil
}

// GetUniqueTypes return types from given dataset of content
func (c Contents) GetUniqueTypes() []string {
	var types []string

LOOP:
	for _, i := range []Content(c) {
		for _, t := range types {
			if i.Type == t {
				continue LOOP
			}
		}
		types = append(types, i.Type)
	}
	return types
}
