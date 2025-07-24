package room

import (
	"errors"

	"github.com/sebenitezg/hotel-service/internal/core"
	"github.com/sebenitezg/hotel-service/pkg/logger"

	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type RoomService struct {
	roomRepo          *RoomRepository
	hotelValidator    core.HotelValidator
	roomTypeValidator core.RoomTypeValidator
	log               *zap.SugaredLogger
}

func NewService(
	roomRepo *RoomRepository,
	hotelValidator core.HotelValidator,
	roomTypeValidator core.RoomTypeValidator,
) *RoomService {
	return &RoomService{
		roomRepo:          roomRepo,
		hotelValidator:    hotelValidator,
		roomTypeValidator: roomTypeValidator,
		log:               logger.GetLogger(),
	}
}

func (s *RoomService) ListRoomsByHotelID(hotelID uuid.UUID) (Rooms, error) {
	rooms, err := s.roomRepo.GetByHotelID(hotelID)
	if err != nil {
		s.log.Errorw("error retrieving rooms by hotel ID", "hotelID", hotelID, "error", err)
		return nil, err
	}
	return rooms, nil
}

func (s *RoomService) RetrieveRoomByHotelRoomID(
	hotelID uuid.UUID, roomID uuid.UUID,
) (*Room, error) {
	room, err := s.roomRepo.GetByID(roomID)
	if err != nil {
		return nil, err
	}

	if room.HotelID != hotelID {
		s.log.Errorw(
			"error retrieving the room by hotel and room IDs",
			"hotelID", hotelID, "roomID", roomID, "error", err,
		)
		return nil, errors.New("hotel doesn't have any room with the provided identifier")
	}

	return room, nil
}

func (s *RoomService) CreateRoom(r *Room) (*Room, error) {

	hotelExist, err := s.hotelValidator.ValidateHotelExists(r.HotelID)
	if err != nil {
		s.log.Errorw("error validating hotel existence", "hotelID", r.HotelID, "error", err)
		return nil, err
	}
	if !hotelExist {
		s.log.Errorw("hotel does not exist", "hotelID", r.HotelID)
		return nil, errors.New("hotel does not exist")
	}

	roomTypeExist, err := s.roomTypeValidator.ValidateRoomTypeExists(r.RoomTypeID)
	if err != nil {
		s.log.Errorw(
			"error validating room type existence",
			"roomTypeID", r.RoomTypeID, "error", err,
		)
		return nil, err
	}
	if !roomTypeExist {
		s.log.Errorw("room type does not exist", "roomTypeID", r.RoomTypeID)
		return nil, errors.New("room type does not exist")
	}

	if err := s.roomRepo.Save(r); err != nil {
		s.log.Errorw("error creating new room", "error", err)
		return nil, err
	}

	s.log.Infow("room created successfully", "hotel_id", r.ID)

	return r, nil
}

func (s *RoomService) UpdatePartiallyRoom(
	roomID uuid.UUID,
	uuidHotelID uuid.UUID,
	roomTypeID *uuid.UUID,
	floor *int,
	number *int,
	name *string,
	status *string,
) (*Room, error) {

	room, err := s.roomRepo.GetByID(roomID)
	if err != nil {
		s.log.Errorw("failure updating partially room", "roomID", roomID, "error", err)
		return nil, err
	}
	if room == nil {
		s.log.Errorw("room not found", "roomID", roomID)
		return nil, errors.New("room not found")
	}

	if room.HotelID != uuidHotelID {
		s.log.Errorw("hotel does not have the room with the provided ID", "hotelID", uuidHotelID, "roomID", roomID)
		return nil, errors.New("hotel does not have the room with the provided ID")
	}

	if roomTypeID != nil {
		room.RoomTypeID = *roomTypeID
	}
	if floor != nil {
		room.Floor = *floor
	}
	if number != nil {
		room.Number = *number
	}
	if name != nil {
		room.Name = *name
	}
	if status != nil {
		room.Status = *status
	}

	err = s.roomRepo.Update(room)
	if err != nil {
		s.log.Errorw("failure updating partially room", "roomID", roomID, "error", err)
		return nil, err
	}

	return room, nil
}
