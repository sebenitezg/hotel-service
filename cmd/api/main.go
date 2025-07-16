package main

import (
	"hotel-service/config"
	"hotel-service/internal/hotel"
	"hotel-service/internal/room"
	"hotel-service/pkg/db"
	"hotel-service/pkg/logger"
	"hotel-service/pkg/server/rest"

	"log"

	"github.com/go-playground/validator/v10"
)

func main() {

	// Load global configurations
	configs, err := config.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading configurations: %v", err)
	}

	// Initialize Logger
	logger.GetLogger(configs.Server.DebugMode)
	defer logger.CloseLogger()

	// Validator service
	validatorInstance := validator.New()

	// Create DB connection
	database := db.NewConnection(configs.Database)

	// Initialize HTTP Server
	httpServer := rest.NewHTTPServer(configs.Server)

	// Setup Repositories
	roomRepository := room.NewRepository(database)
	hotelRepository := hotel.NewRepository(database)

	// Setup Services
	hotelService := hotel.NewService(hotelRepository)
	roomService := room.NewService(roomRepository)

	// Initialize Controllers
	hotel.NewController(httpServer, validatorInstance, hotelService)
	room.NewController(httpServer, validatorInstance, roomService)

	httpServer.Start()
}
