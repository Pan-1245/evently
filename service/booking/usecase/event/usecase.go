package usecase

import (
	"context"

	"github.com/Pan-1245/evently/service/booking/domain"
	port "github.com/Pan-1245/evently/service/booking/port/event"
	"github.com/Pan-1245/evently/service/booking/usecase/event/dto"
	"github.com/Pan-1245/evently/shared/response"
	"github.com/google/uuid"
)

type EventUseCase struct {
	repo port.EventRepository
}

func NewEventUsecase(repo port.EventRepository) *EventUseCase {
	return &EventUseCase{repo: repo}
}

func (uc *EventUseCase) ListPaginated(ctx context.Context, page int, limit int) (*response.PageResponse[*dto.EventResponse], error) {
	events, total, err := uc.repo.ListPaginated(ctx, page, limit)
	if err != nil {
		return nil, err
	}

	eventResponses := make([]*dto.EventResponse, 0, len(events))
	for _, e := range events {
		eventResponses = append(eventResponses, MapEventResponse(e))
	}

	// if limit <= 0, we treat it as "fetch all"
	if limit <= 0 {
		limit = len(events)
		page = 1
	}

	totalPages := int((total + int64(limit) - 1) / int64(limit))

	res := &response.PageResponse[*dto.EventResponse]{
		Page:       page,
		PageSize:   limit,
		TotalItems: total,
		TotalPages: totalPages,
		Data:       eventResponses,
	}
	return res, nil
}

func (uc *EventUseCase) GetByID(ctx context.Context, id uuid.UUID) (*dto.EventResponse, error) {
	event, err := uc.repo.GetByID(ctx, id)
	if err != nil {
		return nil, err
	}

	res := MapEventResponse(event)
	return res, nil
}

func (uc *EventUseCase) GetByOrganizerID(ctx context.Context, organizerID uuid.UUID) ([]*dto.EventResponse, error) {
	events, err := uc.repo.GetByOrganizerID(ctx, organizerID)
	if err != nil {
		return nil, err
	}

	var res []*dto.EventResponse
	for _, e := range events {
		res = append(res, MapEventResponse(e))
	}
	return res, nil
}

func (uc *EventUseCase) Create(ctx context.Context, req dto.UpsertEventRequest) (*dto.EventResponse, error) {
	event := &domain.Event{
		ID:          uuid.New(),
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		IsActive:    req.IsActive,
		OrganizerID: req.OrganizerID,
	}
	if err := uc.repo.Create(ctx, event); err != nil {
		return nil, err
	}

	res := MapEventResponse(event)
	return res, nil
}

func (uc *EventUseCase) Update(ctx context.Context, id uuid.UUID, req dto.UpsertEventRequest) error {
	event := &domain.Event{
		ID:          id,
		Title:       req.Title,
		Description: req.Description,
		StartTime:   req.StartTime,
		EndTime:     req.EndTime,
		Location:    req.Location,
		IsActive:    req.IsActive,
	}
	return uc.repo.Update(ctx, event)
}

func (uc *EventUseCase) Delete(ctx context.Context, id uuid.UUID) error {
	return uc.repo.Delete(ctx, id)
}

func MapEventResponse(e *domain.Event) *dto.EventResponse {
	return &dto.EventResponse{
		ID:          e.ID,
		Title:       e.Title,
		Description: e.Description,
		StartTime:   e.StartTime,
		EndTime:     e.EndTime,
		Location:    e.Location,
		IsActive:    e.IsActive,
		OrganizerID: e.OrganizerID,
	}
}
