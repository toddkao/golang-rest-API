package models

import "time"

// LocationGroup struct
type LocationGroup struct {
	ID          string    `json:"id,omitempty" gorethink:"id,omitempty"`
	Name        string    `validate:"nonzero" json:"name" gorethink:"name"`
	Description string    `validate:"nonzero" json:"description" gorethink:"description"`
	Status      byte      `validate:"nonzero" json:"status" gorethink:"status"`
	Created     time.Time `json:"created_at" gorethink:"created_at"`
	Updated     time.Time `json:"updated_at" gorethink:"updated_at"`
}
