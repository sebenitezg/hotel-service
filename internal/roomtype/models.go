package roomtype

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/shopspring/decimal"
	"github.com/uptrace/bun"
)

type BedType string

const (
	QUEEN_SIZE BedType = "queen"
	KING_SIZE  BedType = "king"
	TWIN_SIZE  BedType = "twin"
)

type RoomType struct {
	bun.BaseModel `bun:"table:room_types"`
	ID            uuid.UUID       `bun:"id"`
	CreatedAt     time.Time       `bun:"created_at"`
	UpdatedAt     time.Time       `bun:"updated_at"`
	HotelID       uuid.UUID       `bun:"hotel_id"`
	Name          string          `bun:"name"`
	Description   string          `bun:"description"`
	NumberOfBeds  int             `bun:"number_of_beds"`
	BedType       string          `bun:"bed_type"`
	MaxOccupancy  int             `bun:"max_occupancy"`
	BasePrice     decimal.Decimal `bun:"base_price"`
}

type RoomTypes []RoomType

func NewRoomType(
	hotelID uuid.UUID,
	name string,
	description string,
	numberOfBeds int,
	bedType string,
	maxOccupancy int,
	basePrice decimal.Decimal,
) (*RoomType, error) {
	now := time.Now().UTC()

	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return &RoomType{
		ID:           id,
		CreatedAt:    now,
		UpdatedAt:    now,
		HotelID:      hotelID,
		Name:         name,
		Description:  description,
		NumberOfBeds: numberOfBeds,
		BedType:      bedType,
		MaxOccupancy: maxOccupancy,
		BasePrice:    basePrice,
	}, nil
}
