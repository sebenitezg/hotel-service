package core

import "github.com/gofrs/uuid/v5"

type HotelValidator interface {
	ValidateHotelExists(id uuid.UUID) (bool, error)
}

type RoomTypeValidator interface {
	ValidateRoomTypeExists(id uuid.UUID) (bool, error)
}
