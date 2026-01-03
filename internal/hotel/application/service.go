package application

import (
	"errors"

	"github.com/sebenitezg/hotel-service/pkg/logger"
	"github.com/sebenitezg/hotel-service/internal/hotel/repository"
	"github.com/sebenitezg/hotel-service/internal/hotel/models"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type HotelService struct {
	hotelRepo *repository.HotelRepository
	log       *zap.SugaredLogger
}

func NewService(
	hotelRepo *repository.HotelRepository,
) *HotelService {
	return &HotelService{
		hotelRepo: hotelRepo,
		log:       logger.GetLogger(),
	}
}

func (s *HotelService) ListHotels() (models.Hotels, error) {
	s.log.Infof("fetching all hotels")

	hotels, err := s.hotelRepo.GetAll()
	if err != nil {
		s.log.Errorw("error getting hotels information", "error", err)
		return models.Hotels{}, err
	}

	return hotels, nil
}

func (s *HotelService) GetHotelByID(id uuid.UUID) (*models.Hotel, error) {
	s.log.Infof("fetching hotel by id: %s", id)

	hotel, err := s.hotelRepo.GetByID(id)
	if err != nil {
		s.log.Errorw("error fetching hotel by ID", "hotelID", id, "reason", err)
		return nil, errors.New("unexpected error fetching hotel")
	}
	if hotel == nil {
		s.log.Errorw("hotel not found", "hotelID", id)
		return nil, models.ErrHotelNotFound
	}

	return hotel, nil
}

func (s *HotelService) CreateHotel(hotel *models.Hotel) (*models.Hotel, error) {
	s.log.Infof("creating a new hotel: %s", hotel.Name)

	if err := s.hotelRepo.Save(hotel); err != nil {
		s.log.Errorf("failed to create hotel: %v", err)
		return nil, err
	}

	s.log.Infow("hotel created successfully", "hotel_id", hotel.ID)
	return hotel, nil
}

func (s *HotelService) UpdatePartiallyHotel(
	id uuid.UUID,
	name *string,
	address *string,
	status *string,
	description *string,
) (*models.Hotel, error) {
	s.log.Infof("updating hotel instance with id: %v", id)

	hotel, err := s.hotelRepo.GetByID(id)
	if err != nil {
		return nil, nil
	}

	if hotel == nil {
		s.log.Infof("hotel not found: %v", id)
		return nil, models.ErrHotelNotFound
	}

	if name != nil {
		hotel.Name = *name
	}
	if address != nil {
		hotel.Address = *address
	}
	if status != nil {
		hotel.Status = *status
	}
	if description != nil {
		hotel.Status = *description
	}

	if err := s.hotelRepo.Update(hotel); err != nil {
		s.log.Errorf("failed updating hotel information", "error", err)
		return nil, err
	}

	s.log.Infow("updated hotel information successfully", "hotel_id", hotel.ID)
	return hotel, nil
}

func (s *HotelService) ValidateHotelExists(id uuid.UUID) (bool, error) {
	hotel, err := s.hotelRepo.GetByID(id)
	if err != nil {
		return false, err
	}
	return hotel != nil, nil
}
