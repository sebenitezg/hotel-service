package room

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"
)

// --------------------
// DB models
// --------------------
type Room struct {
	bun.BaseModel `bun:"table:rooms"`
	ID            uuid.UUID `bun:"id"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
	HotelID       uuid.UUID `bun:"hotel_id"`
	RoomTypeID    uuid.UUID `bun:"room_type_id"`
	Floor         int       `bun:"floor"`
	Number        int       `bun:"number"`
	Name          string    `bun:"name"`
	Status        string    `bun:"status"`
}

type Rooms []Room

func NewRoom(
	hotelID uuid.UUID,
	roomTypeID uuid.UUID,
	floor int,
	number int,
	name string,
	status string,
) (*Room, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}
	return &Room{
		ID:         id,
		CreatedAt:  time.Now().UTC(),
		UpdatedAt:  time.Now().UTC(),
		HotelID:    hotelID,
		RoomTypeID: roomTypeID,
		Floor:      floor,
		Number:     number,
		Name:       name,
		Status:     status,
	}, nil
}
