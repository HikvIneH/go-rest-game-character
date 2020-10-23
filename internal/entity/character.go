package entity

import (
	"time"
)

// Character represents an album record.
type Character struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CharCode  string    `json:"character_code"`
	CharValue string    `json:"character_value"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
