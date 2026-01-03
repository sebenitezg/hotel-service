package interfaces

import "github.com/gofrs/uuid/v5"

type HotelValidator interface {
	ValidateHotelExists(id uuid.UUID) (bool, error)
}
