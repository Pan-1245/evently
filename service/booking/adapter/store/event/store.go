package store

import (
	"context"

	"github.com/Pan-1245/evently/service/booking/domain"
	port "github.com/Pan-1245/evently/service/booking/port/event"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type eventRepository struct {
	db *gorm.DB
}

func NewEventRepository(db *gorm.DB) port.EventRepository {
	return &eventRepository{db: db}
}

// ListPaginated implements port.EventRepository.
func (r *eventRepository) ListPaginated(ctx context.Context, page int, limit int) ([]*domain.Event, int64, error) {
	var events []*domain.Event
	var total int64

	db := r.db.WithContext(ctx).Model(&domain.Event{})

	// Count total first
	if err := db.Count(&total).Error; err != nil {
		return nil, 0, err
	}

	query := db.Order("created_at DESC")
	if limit > 0 {
		offset := (page - 1) * limit
		query = query.Offset(offset).Limit(limit)
	}

	err := query.Find(&events).Error
	return events, total, err
}

// GetByID implements port.EventRepository.
func (r *eventRepository) GetByID(ctx context.Context, id uuid.UUID) (*domain.Event, error) {
	var event domain.Event
	err := r.db.WithContext(ctx).First(&event, "id = ?", id).Error
	if err != nil {
		return nil, err
	}
	return &event, nil
}

// GetByOrganizerID implements port.EventRepository.
func (r *eventRepository) GetByOrganizerID(ctx context.Context, organizerID uuid.UUID) ([]*domain.Event, error) {
	var events []*domain.Event
	err := r.db.WithContext(ctx).Where("organizer_id = ?", organizerID).Find(&events).Error
	return events, err
}

// Create implements port.EventRepository.
func (r *eventRepository) Create(ctx context.Context, event *domain.Event) error {
	return r.db.WithContext(ctx).Create(event).Error
}

// Update implements port.EventRepository.
func (r *eventRepository) Update(ctx context.Context, event *domain.Event) error {
	return r.db.WithContext(ctx).Save(event).Error
}

// Delete implements port.EventRepository.
func (r *eventRepository) Delete(ctx context.Context, id uuid.UUID) error {
	return r.db.WithContext(ctx).Delete(&domain.Event{}, "id = ?", id).Error
}
