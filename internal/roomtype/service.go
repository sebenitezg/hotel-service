package roomtype

import (
	"errors"

	"github.com/sebenitezg/hotel-service/internal/core"
	"github.com/sebenitezg/hotel-service/pkg/logger"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"
	"go.uber.org/zap"
)

type RoomTypeService struct {
	roomTypeRepo   *RoomTypeRepository
	hotelValidator core.HotelValidator
	log            *zap.SugaredLogger
}

func NewService(
	roomTypeRepo *RoomTypeRepository,
	hotelValidator core.HotelValidator,
) *RoomTypeService {
	return &RoomTypeService{
		roomTypeRepo:   roomTypeRepo,
		hotelValidator: hotelValidator,
		log:            logger.GetLogger(),
	}
}

func (s *RoomTypeService) ListRoomTypesByHotelID(hotelID uuid.UUID) (RoomTypes, error) {
	rooms, err := s.roomTypeRepo.GetByHotelID(hotelID)
	if err != nil {
		s.log.Errorw(
			"error retrieving room types by hotel ID",
			"hotelID", hotelID, "error", err,
		)
		return nil, err
	}
	return rooms, nil
}

func (s *RoomTypeService) RetrieveRoomTypeByHotelRoomTypeID(
	hotelID uuid.UUID, roomTypeID uuid.UUID,
) (*RoomType, error) {
	roomType, err := s.roomTypeRepo.GetByID(roomTypeID)
	if err != nil {
		return nil, err
	}

	if roomType == nil {
		s.log.Error("room type not found", "roomTypeID", roomTypeID)
		return nil, ErrRoomTypeNotFound
	}

	if roomType.HotelID != hotelID {
		s.log.Errorw(
			"error retrieving the room by hotel and room type IDs",
			"hotelID", hotelID, "roomTypeID", roomTypeID, "error", err,
		)
		return nil, errors.New("hotel doesn't have any room type with the provided identifier")
	}

	return roomType, nil
}

func (s *RoomTypeService) CreateRoomType(r *RoomType) (*RoomType, error) {

	hotelExist, err := s.hotelValidator.ValidateHotelExists(r.HotelID)
	if err != nil {
		s.log.Errorw(
			"error validating hotel existence",
			"hotelID", r.HotelID, "error", err,
		)
		return nil, err
	}
	if !hotelExist {
		s.log.Errorw("hotel does not exist", "hotelID", r.HotelID)
		return nil, errors.New("hotel does not exist")
	}

	if err := s.roomTypeRepo.Save(r); err != nil {
		s.log.Errorw("error creating new room", "error", err)
		return nil, err
	}

	s.log.Infow("room created successfully", "hotelID", r.ID, "roomTypeID", r.ID)

	return r, nil
}

func (s *RoomTypeService) UpdatePartiallyRoomType(
	roomTypeID uuid.UUID,
	uuidHotelID uuid.UUID,
	name *string,
	description *string,
	numberOfBeds *int,
	bedType *string,
	maxOccupancy *int,
	basePrice *decimal.Decimal,
) (*RoomType, error) {
	roomType, err := s.roomTypeRepo.GetByID(roomTypeID)
	if err != nil {
		s.log.Errorw(
			"failure geting RoomType entity to update partially it",
			"roomTypeID", roomTypeID, "error", err,
		)
		return nil, err
	}
	if roomType == nil {
		s.log.Errorw("room type entity not found", "roomTypeID", roomTypeID)
		return nil, ErrRoomTypeNotFound
	}

	if roomType.HotelID != uuidHotelID {
		s.log.Errorw(
			"hotel does not have the room with the provided ID",
			"hotelID", uuidHotelID, "roomTypeID", roomTypeID,
		)
		return nil, errors.New("hotel does not have the room with the provided ID")
	}

	if name != nil {
		roomType.Name = *name
	}
	if description != nil {
		roomType.Description = *description
	}
	if numberOfBeds != nil {
		roomType.NumberOfBeds = *numberOfBeds
	}
	if bedType != nil {
		roomType.BedType = *bedType
	}
	if maxOccupancy != nil {
		roomType.MaxOccupancy = *maxOccupancy
	}
	if basePrice != nil {
		roomType.BasePrice = *basePrice
	}

	err = s.roomTypeRepo.Update(roomType)
	if err != nil {
		s.log.Errorw(
			"failure updating partially room",
			"roomTypeID", roomTypeID, "error", err,
		)
		return nil, err
	}

	return roomType, nil
}

func (s RoomTypeService) ValidateRoomTypeExists(id uuid.UUID) (bool, error) {
	roomType, err := s.roomTypeRepo.GetByID(id)
	if err != nil {
		return false, err
	}
	return roomType != nil, nil
}
