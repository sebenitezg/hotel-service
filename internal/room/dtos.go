package room

import (
	"time"

	"github.com/gofrs/uuid/v5"
)

type CreateRoomRequest struct {
	RoomTypeID uuid.UUID `json:"room_type_id"`
	Floor      int       `json:"floor"`
	Number     int       `json:"number"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
}

type UpdateRoomRequest struct {
	RoomTypeID *uuid.UUID `json:"room_type_id"`
	Floor      *int       `json:"floor"`
	Number     *int       `json:"number"`
	Name       *string    `json:"name"`
	Status     *string    `json:"status"`
}

type RoomResponse struct {
	ID         uuid.UUID `json:"id"`
	CreatedAt  string    `json:"created_at"`
	UpdatedAt  string    `json:"updated_at"`
	HotelID    uuid.UUID `json:"hotel_id"`
	RoomTypeID uuid.UUID `json:"room_type_id"`
	Floor      int       `json:"floor"`
	Number     int       `json:"number"`
	Name       string    `json:"name"`
	Status     string    `json:"status"`
}

type ListRoomsResponse struct {
	Results []RoomResponse `json:"results"`
}

func NewRoomResponse(r *Room) RoomResponse {
	return RoomResponse{
		ID:         r.ID,
		CreatedAt:  r.CreatedAt.Format(time.RFC3339),
		UpdatedAt:  r.UpdatedAt.Format(time.RFC3339),
		HotelID:    r.HotelID,
		RoomTypeID: r.RoomTypeID,
		Floor:      r.Floor,
		Number:     r.Number,
		Name:       r.Name,
		Status:     r.Status,
	}
}

func NewListRoomsResponse(rooms Rooms) ListRoomsResponse {
	roomsResponse := make([]RoomResponse, len(rooms))
	for i, room := range rooms {
		roomsResponse[i] = NewRoomResponse(&room)
	}
	return ListRoomsResponse{
		Results: roomsResponse,
	}
}
