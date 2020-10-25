package entity

import (
	"time"
)

// Character represents an album record.
type Character struct {
	ID             string    `json:"id"`
	Name           string    `json:"name"`
	CharacterCode  int64     `json:"character_code"`
	CharacterPower int64     `json:"character_power"`
	CharacterValue int64     `json:"character_value"`
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}
