package hotel

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type CreateHotelRequest struct {
	Name        string `json:"name"`
	Address     string `json:"address"`
	Country     string `json:"country"`
	State       string `json:"state"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

type UpdateHotelRequest struct {
	Name        *string `json:"name"`
	Address     *string `json:"address"`
	Status      *string `json:"status"`
	Description *string `json:"description"`
}

type HotelResponse struct {
	ID          uuid.UUID `json:"id"`
	CreatedAt   string    `json:"created_at"`
	UpdatedAt   string    `json:"updated_at"`
	Name        string    `json:"name"`
	Address     string    `json:"address"`
	Country     string    `json:"country"`
	State       string    `json:"state"`
	Status      string    `json:"status"`
	Description string    `json:"description"`
}

type ListHotelsResponse struct {
	Results []HotelResponse `json:"results"`
}

func NewHotelResponse(hotel *Hotel) HotelResponse {
	return HotelResponse{
		ID:          hotel.ID,
		CreatedAt:   hotel.CreatedAt.Format(time.RFC3339),
		UpdatedAt:   hotel.UpdatedAt.Format(time.RFC3339),
		Name:        hotel.Name,
		Address:     hotel.Address,
		Country:     hotel.Country,
		State:       hotel.State,
		Status:      hotel.Status,
		Description: hotel.Description,
	}
}

func NewListHotelsResponse(hotels Hotels) ListHotelsResponse {
	hotelsResponse := make([]HotelResponse, len(hotels))
	for i, hotel := range hotels {
		hotelsResponse[i] = NewHotelResponse(&hotel)
	}
	return ListHotelsResponse{
		Results: hotelsResponse,
	}
}
