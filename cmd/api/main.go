package main

import (
	"github.com/sebenitezg/hotel-service/config"
	"github.com/sebenitezg/hotel-service/internal/hotel"
	"github.com/sebenitezg/hotel-service/internal/room"
	"github.com/sebenitezg/hotel-service/internal/roomtype"
	"github.com/sebenitezg/hotel-service/pkg/db"
	"github.com/sebenitezg/hotel-service/pkg/logger"
	"github.com/sebenitezg/hotel-service/pkg/server/rest"

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
	roomTypeRepository := roomtype.NewRepository(database)
	hotelRepository := hotel.NewRepository(database)

	// Setup Services
	hotelService := hotel.NewService(hotelRepository)
	roomTypeService := roomtype.NewService(roomTypeRepository, hotelService)
	roomService := room.NewService(roomRepository, hotelService, roomTypeService)

	// Initialize Controllers
	hotel.NewController(httpServer, validatorInstance, hotelService)
	roomtype.NewController(httpServer, validatorInstance, roomTypeService)
	room.NewController(httpServer, validatorInstance, roomService)

	httpServer.Start()
}
