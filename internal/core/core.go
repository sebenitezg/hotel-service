package core

import "github.com/gofrs/uuid/v5"

type HotelValidator interface {
	ValidateHotelExists(hotelID uuid.UUID) (bool, error)
}
