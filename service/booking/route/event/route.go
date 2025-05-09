package route

import (
	handler "github.com/Pan-1245/evently/service/booking/adapter/http/handler/event"
	"github.com/gin-gonic/gin"
)

func Register(r *gin.Engine, h *handler.EventHandler) {
	group := r.Group("/events")
	{
		group.GET("", h.ListPaginated)
		group.GET("/:id", h.GetByID)
		group.GET("/organizers/:organizer_id", h.GetByOrganizerID)
		group.POST("", h.Create)
		group.PUT("/:id", h.Update)
		group.DELETE("/:id", h.Delete)
	}

	r.GET("/organizers/:id/events", h.GetByOrganizerID)
}
