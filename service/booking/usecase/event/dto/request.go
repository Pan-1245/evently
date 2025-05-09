package dto

import (
	"time"

	"github.com/google/uuid"
)

type UpsertEventRequest struct {
	Title       string    `json:"title" validate:"required"`
	Description string    `json:"description"`
	StartTime   time.Time `json:"start_time" validate:"required"`
	EndTime     time.Time `json:"end_time" validate:"required"`
	Location    string    `json:"location" valudate:"required"`
	OrganizerID uuid.UUID `json:"organizer_id" validate:"required"`
	IsActive    bool      `json:"is_active" validate:"required"`
}
