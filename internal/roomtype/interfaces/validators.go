package interfaces

import "github.com/gofrs/uuid/v5"

type RoomTypeValidator interface {
	ValidateRoomTypeExists(id uuid.UUID) (bool, error)
}
