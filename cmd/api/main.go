package main

import (
	"github.com/sebenitezg/hotel-service/config"
	hotelRepo "github.com/sebenitezg/hotel-service/internal/hotel/repository"
	roomRepo "github.com/sebenitezg/hotel-service/internal/room/repository"
	roomTypeRepo "github.com/sebenitezg/hotel-service/internal/roomtype/repository"
	hotelSrv "github.com/sebenitezg/hotel-service/internal/hotel/application"
	roomSrv "github.com/sebenitezg/hotel-service/internal/room/application"
	roomTypeSrv "github.com/sebenitezg/hotel-service/internal/roomtype/application"
	hotelCtrl "github.com/sebenitezg/hotel-service/internal/hotel/controller"
	roomCtrl "github.com/sebenitezg/hotel-service/internal/room/controller"
	roomTypeCtrl "github.com/sebenitezg/hotel-service/internal/roomtype/controller"
	
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
	roomRepository := roomRepo.NewRepository(database)
	roomTypeRepository := roomTypeRepo.NewRepository(database)
	hotelRepository := hotelRepo.NewRepository(database)

	// Setup Services
	hotelService := hotelSrv.NewService(hotelRepository)
	roomTypeService := roomTypeSrv.NewService(roomTypeRepository, hotelService)
	roomService := roomSrv.NewService(roomRepository, hotelService, roomTypeService)

	// Initialize Controllers
	hotelCtrl.NewController(httpServer, validatorInstance, hotelService)
	roomTypeCtrl.NewController(httpServer, validatorInstance, roomTypeService)
	roomCtrl.NewController(httpServer, validatorInstance, roomService)

	httpServer.Start()
}
