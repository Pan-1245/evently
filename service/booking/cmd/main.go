package main

import (
	"log"
	"os"

	"github.com/Pan-1245/evently/service/booking/config"
	"github.com/Pan-1245/evently/service/booking/domain"
	"github.com/Pan-1245/evently/service/booking/infra"
	"github.com/gin-gonic/gin"
)

func main() {
	config.LoadEnv()

	db := infra.NewDB()

	if err := db.AutoMigrate(&domain.Event{}); err != nil {
		log.Fatalf("Failed to auto-migrate: %v", err)
	}

	// eventRepo := store.NewEventRepository(db)
	// eventUsecase := usecase.NewEventUsecase(eventRepo)
	// eventHandler := http.NewEventHandler(eventUsecase)

	r := gin.Default()
	// route.SetupRoutes(r, eventHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	log.Printf("🚀 Server running on port %s", port)
	if err := r.Run(":" + port); err != nil {
		log.Fatalf("Could not start server: %v", err)
	}
}
