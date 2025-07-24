package roomtype

import (
	"encoding/json"
	"errors"
	"hotel-service/pkg/logger"
	"hotel-service/pkg/server/rest"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type RoomTypeController struct {
	validator       *validator.Validate
	roomTypeService *RoomTypeService
	log             *zap.SugaredLogger
}

func NewController(
	server *rest.HTTPServer,
	validator *validator.Validate,
	roomTypeService *RoomTypeService,
) *RoomTypeController {
	c := &RoomTypeController{
		validator:       validator,
		roomTypeService: roomTypeService,
		log:             logger.GetLogger(),
	}

	server.Router.Group(func(r chi.Router) {
		r.Get("/v1/hotels/{hotel_id}/roomtypes", c.handleListHotelRoomTypes)
		r.Get("/v1/hotels/{hotel_id}/roomtypes/{room_type_id}", c.handleGetHotelRoomType)
		r.Post("/v1/hotels/{hotel_id}/roomtypes", c.handleCreateHotelRoomType)
		r.Patch("/v1/hotels/{hotel_id}/roomtypes/{room_type_id}", c.handlePartialUpdateHotelRoomType)
		r.Delete("/v1/hotels/{hotel_id}/roomtypes/{room_type_id}", c.handleDeleteHotelRoomType)
	})

	return c
}

func (c *RoomTypeController) handleListHotelRoomTypes(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotel_id")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid hotel id"))
		return
	}
	hotelRooms, err := c.roomTypeService.ListRoomTypesByHotelID(uuidHotelID)
	if err != nil {
		c.log.Errorw("error retrieving hotel's room types", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, err)
	}

	resp := NewListRoomTypesResponse(hotelRooms)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *RoomTypeController) handleGetHotelRoomType(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotel_id")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid hotel id"))
		return
	}

	roomTypeID := chi.URLParam(r, "room_type_id")
	uuidRoomTypeID, err := uuid.FromString(roomTypeID)
	if err != nil {
		c.log.Errorw("invalid room type id", "roomTypeID", roomTypeID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid room type id"))
		return
	}

	resp, err := c.roomTypeService.RetrieveRoomTypeByHotelRoomTypeID(uuidHotelID, uuidRoomTypeID)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	rest.RenderJSON(r.Context(), w, http.StatusCreated, resp)
}

func (c *RoomTypeController) handleCreateHotelRoomType(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotel_id")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid hotel id"))
		return
	}

	var payload CreateRoomTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.log.Errorw("failed to decode request body", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}
	if err := c.validator.Struct(payload); err != nil {
		c.log.Errorw("validation error", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}

	roomType, err := NewRoomType(
		uuidHotelID,
		payload.Name,
		payload.Description,
		payload.NumberOfBeds,
		payload.BedType,
		payload.MaxOccupancy,
		payload.BasePrice,
	)
	if err != nil {
		c.log.Errorw("failure creating room type", "error", err)
		rest.RenderError(r.Context(), w, err)
	}

	roomType, err = c.roomTypeService.CreateRoomType(roomType)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	resp := NewRoomTypeResponse(roomType)

	rest.RenderJSON(r.Context(), w, http.StatusCreated, resp)
}

func (c *RoomTypeController) handlePartialUpdateHotelRoomType(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotel_id")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid hotel id"))
		return
	}

	roomID := chi.URLParam(r, "room_type_id")
	uuidRoomTypeID, err := uuid.FromString(roomID)
	if err != nil {
		c.log.Errorw("invalid room type id", "roomTypeID", roomID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("invalid room type id"))
	}

	var payload UpdateRoomTypeRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.log.Errorw("failed to decode request body", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}
	if err := c.validator.Struct(payload); err != nil {
		c.log.Errorw("validation error", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}

	hotel, err := c.roomTypeService.UpdatePartiallyRoomType(
		uuidRoomTypeID,
		uuidHotelID,
		payload.Name,
		payload.Description,
		payload.NumberOfBeds,
		payload.BedType,
		payload.MaxOccupancy,
		payload.BasePrice,
	)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	resp := NewRoomTypeResponse(hotel)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *RoomTypeController) handleDeleteHotelRoomType(w http.ResponseWriter, r *http.Request) {

}
