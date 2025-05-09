package port

import (
	"context"

	"github.com/Pan-1245/evently/service/booking/domain"
	"github.com/google/uuid"
)

type EventRepository interface {
	ListPaginated(ctx context.Context, page, limit int) ([]*domain.Event, int64, error)
	GetByID(ctx context.Context, id uuid.UUID) (*domain.Event, error)
	GetByOrganizerID(ctx context.Context, organizerID uuid.UUID) ([]*domain.Event, error)

	Create(ctx context.Context, event *domain.Event) error
	Update(ctx context.Context, event *domain.Event) error
	Delete(ctx context.Context, id uuid.UUID) error
}
