package main

import (
	"log"
	"os"

	handler "github.com/Pan-1245/evently/service/booking/adapter/http/handler/event"
	store "github.com/Pan-1245/evently/service/booking/adapter/store/event"
	"github.com/Pan-1245/evently/service/booking/config"
	"github.com/Pan-1245/evently/service/booking/infra"
	port "github.com/Pan-1245/evently/service/booking/port/event"
	route "github.com/Pan-1245/evently/service/booking/route/event"
	usecase "github.com/Pan-1245/evently/service/booking/usecase/event"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := infra.NewDB()

	var (
		repo    port.EventRepository  = store.NewEventRepository(db)
		usecase *usecase.EventUseCase = usecase.NewEventUsecase(repo)
		handler *handler.EventHandler = handler.NewEventHandler(usecase)
	)

	r := gin.Default()
	route.Register(r, handler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
