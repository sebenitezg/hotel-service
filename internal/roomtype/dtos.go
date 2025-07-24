package roomtype

import (
	"time"

	"github.com/shopspring/decimal"
)

type CreateRoomTypeRequest struct {
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	NumberOfBeds int             `json:"number_of_beds"`
	BedType      string          `json:"bed_type"`
	MaxOccupancy int             `json:"max_occupancy"`
	BasePrice    decimal.Decimal `json:"base_price"`
}

type UpdateRoomTypeRequest struct {
	Name         *string          `json:"name"`
	Description  *string          `json:"description"`
	NumberOfBeds *int             `json:"number_of_beds"`
	BedType      *string          `json:"bed_type"`
	MaxOccupancy *int             `json:"max_occupancy"`
	BasePrice    *decimal.Decimal `json:"base_price"`
}

type RoomTypeResponse struct {
	ID           string          `json:"id"`
	CreatedAt    string          `json:"created_at"`
	UpdatedAt    string          `json:"updated_at"`
	HotelID      string          `json:"hotel_id"`
	Name         string          `json:"name"`
	Description  string          `json:"description"`
	NumberOfBeds int             `json:"number_of_beds"`
	BedType      string          `json:"bed_type"`
	MaxOccupancy int             `json:"max_occupancy"`
	BasePrice    decimal.Decimal `json:"base_price"`
}

type ListRoomTypeResponse struct {
	Results []RoomTypeResponse `json:"results"`
}

func NewRoomTypeResponse(rt *RoomType) RoomTypeResponse {
	return RoomTypeResponse{
		ID:           rt.ID.String(),
		CreatedAt:    rt.CreatedAt.Format(time.RFC3339),
		UpdatedAt:    rt.UpdatedAt.Format(time.RFC3339),
		HotelID:      rt.HotelID.String(),
		Name:         rt.Name,
		Description:  rt.Description,
		NumberOfBeds: rt.NumberOfBeds,
		BedType:      rt.BedType,
		MaxOccupancy: rt.MaxOccupancy,
		BasePrice:    rt.BasePrice,
	}
}

func NewListRoomTypesResponse(rts RoomTypes) ListRoomTypeResponse {
	responses := make([]RoomTypeResponse, len(rts))
	for i, rt := range rts {
		responses[i] = NewRoomTypeResponse(&rt)
	}
	return ListRoomTypeResponse{Results: responses}
}
