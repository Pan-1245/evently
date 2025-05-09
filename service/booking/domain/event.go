package domain

import (
	"time"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Event struct {
	ID          uuid.UUID      `gorm:"type:uuid;default:gen_random_uuid();primaryKey" json:"id"`
	Title       string         `gorm:"not null" json:"title"`
	Description string         `json:"description"`
	StartTime   time.Time      `gorm:"not null" json:"start_time"`
	EndTime     time.Time      `gorm:"not null" json:"end_time"`
	Location    string         `gorm:"not null" json:"location"`
	IsActive    bool           `gorm:"not null" json:"is_active"`
	CreatedAt   time.Time      `gorm:"not null" json:"created_at"`
	UpdatedAt   time.Time      `gorm:"not null" json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"-"`
	OrganizerID uuid.UUID      `gorm:"not null" json:"organizer_id"`
}

func (Event) TableName() string {
	return "booking.event"
}
