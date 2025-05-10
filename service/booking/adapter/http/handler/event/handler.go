package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"

	usecase "github.com/Pan-1245/evently/service/booking/usecase/event"
	"github.com/Pan-1245/evently/service/booking/usecase/event/dto"
	"github.com/Pan-1245/evently/shared/response"
)

type EventHandler struct {
	usecase *usecase.EventUseCase
}

func NewEventHandler(usecase *usecase.EventUseCase) *EventHandler {
	return &EventHandler{usecase: usecase}
}

// GET /events?page=1&limit=10
func (h *EventHandler) ListPaginated(c *gin.Context) {
	page, err := strconv.Atoi(c.Query("page"))
	if err != nil || page <= 0 {
		page = 1
	}

	limit, err := strconv.Atoi(c.Query("limit"))
	if err != nil || limit <= 0 {
		limit = 0
	}

	res, err := h.usecase.ListPaginated(c.Request.Context(), page, limit)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Wrapper[any]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Failed to fetch events",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	if len(res.Data) == 0 {
		c.JSON(http.StatusNoContent, response.Wrapper[*response.PageResponse[*dto.EventResponse]]{
			StatusCode: http.StatusNoContent,
			Success:    true,
			Message:    "No events found",
			Data:       response.DataWrapper[*response.PageResponse[*dto.EventResponse]]{Data: &res},
		})
	}

	c.JSON(http.StatusOK, response.Wrapper[*response.PageResponse[*dto.EventResponse]]{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Events retrieved successfully",
		Data:       response.DataWrapper[*response.PageResponse[*dto.EventResponse]]{Data: &res},
	})
}

// GET /events/:id
func (h *EventHandler) GetByID(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid UUID",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	event, err := h.usecase.GetByID(c.Request.Context(), id)
	if err != nil {
		c.JSON(http.StatusNotFound, response.Wrapper[any]{
			StatusCode: http.StatusNotFound,
			Success:    false,
			Message:    "Event not found",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, response.Wrapper[*dto.EventResponse]{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Event retrieved successfully",
		Data:       response.DataWrapper[*dto.EventResponse]{Data: &event},
	})
}

// GET /organizers/:id/events
func (h *EventHandler) GetByOrganizerID(c *gin.Context) {
	organizerID, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid UUID",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	events, err := h.usecase.GetByOrganizerID(c.Request.Context(), organizerID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Wrapper[any]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Failed to fetch events",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusOK, response.Wrapper[[]*dto.EventResponse]{
		StatusCode: http.StatusOK,
		Success:    true,
		Message:    "Events retrieved successfully",
		Data:       response.DataWrapper[[]*dto.EventResponse]{Data: &events},
	})
}

// POST /events
func (h *EventHandler) Create(c *gin.Context) {
	var req dto.UpsertEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid input",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	res, err := h.usecase.Create(c.Request.Context(), req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, response.Wrapper[any]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Failed to create event",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusCreated, response.Wrapper[*dto.EventResponse]{
		StatusCode: http.StatusCreated,
		Success:    true,
		Message:    "Event created successfully",
		Data:       response.DataWrapper[*dto.EventResponse]{Data: &res},
	})
}

// PUT /events/:id
func (h *EventHandler) Update(c *gin.Context) {
	var req dto.UpsertEventRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid input",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid UUID",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	if err := h.usecase.Update(c.Request.Context(), id, req); err != nil {
		c.JSON(http.StatusInternalServerError, response.Wrapper[any]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Failed to update event",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusNoContent, response.Wrapper[any]{
		StatusCode: http.StatusNoContent,
		Success:    true,
		Message:    "Event updated successfully",
		Data:       response.DataWrapper[any]{},
	})
}

// DELETE /events/:id
func (h *EventHandler) Delete(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, response.Wrapper[any]{
			StatusCode: http.StatusBadRequest,
			Success:    false,
			Message:    "Invalid UUID",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	if err := h.usecase.Delete(c.Request.Context(), id); err != nil {
		c.JSON(http.StatusInternalServerError, response.Wrapper[any]{
			StatusCode: http.StatusInternalServerError,
			Success:    false,
			Message:    "Failed to delete event",
			Data:       response.DataWrapper[any]{Error: err.Error()},
		})
		return
	}

	c.JSON(http.StatusNoContent, response.Wrapper[any]{
		StatusCode: http.StatusNoContent,
		Success:    true,
		Message:    "Event deleted successfully",
		Data:       response.DataWrapper[any]{},
	})
}
