package room

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

type HotelController struct {
	validator   *validator.Validate
	roomService *RoomService
	log         *zap.SugaredLogger
}

func NewController(
	server *rest.HTTPServer,
	validator *validator.Validate,
	roomService *RoomService,
) *HotelController {
	c := &HotelController{
		validator:   validator,
		roomService: roomService,
		log:         logger.GetLogger(),
	}

	server.Router.Group(func(r chi.Router) {
		r.Get("/v1/hotels/{hotelID}/rooms", c.handleListHotelRooms)
		r.Get("/v1/hotels/{hotelID}/rooms/{roomID}", c.handleGetHotelRoom)
		r.Post("/v1/hotels/{hotelID}/rooms", c.handleCreateHotelRoom)
		r.Put("/v1/hotels/{hotelID}/rooms/{roomID}", c.handlePartialUpdateHotelRoom)
		r.Delete("/v1/hotels/{hotelID}/rooms/{roomID}", c.handleDeleteHotelRoom)
	})

	return c
}

func (c *HotelController) handleListHotelRooms(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotelID")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("hotel does not exist"))
		return
	}
	hotelRooms, err := c.roomService.ListRoomsByHotelID(uuidHotelID)
	if err != nil {
		c.log.Errorw("error retrieving hotel's rooms", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, err)
	}

	resp := NewListRoomsResponse(hotelRooms)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *HotelController) handleGetHotelRoom(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotelID")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("hotel does not exist"))
		return
	}

	roomID := chi.URLParam(r, "roomID")
	uuidRoomID, err := uuid.FromString(roomID)
	if err != nil {
		c.log.Errorw("invalid room id", "roomID", roomID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("room does not exist"))
		return
	}

	resp, err := c.roomService.RetrieveRoomByHotelRoomID(uuidHotelID, uuidRoomID)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	rest.RenderJSON(r.Context(), w, http.StatusCreated, resp)
}

func (c *HotelController) handleCreateHotelRoom(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotelID")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("hotel does not exist"))
		return
	}

	var payload CreateRoomRequest
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

	room, err := NewRoom(
		uuidHotelID,
		payload.RoomTypeID,
		payload.Floor,
		payload.Number,
		payload.Name,
		payload.Status,
	)
	if err != nil {
		c.log.Errorw("failure creating hotel instance", "error", err)
		rest.RenderError(r.Context(), w, err)
	}

	room, err = c.roomService.CreateRoom(room)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	resp := NewRoomResponse(room)

	rest.RenderJSON(r.Context(), w, http.StatusCreated, resp)
}

func (c *HotelController) handlePartialUpdateHotelRoom(w http.ResponseWriter, r *http.Request) {
	hotelID := chi.URLParam(r, "hotelID")
	uuidHotelID, err := uuid.FromString(hotelID)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", hotelID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("hotel does not exist"))
		return
	}

	roomID := chi.URLParam(r, "id")
	uuidRoomID, err := uuid.FromString(roomID)
	if err != nil {
		c.log.Errorw("invalid room id", "romID", roomID, "error", err)
		rest.RenderError(r.Context(), w, errors.New("room does not exist"))
	}

	var payload UpdateRoomRequest
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

	hotel, err := c.roomService.UpdatePartiallyRoom(
		uuidRoomID,
		uuidHotelID,
		payload.RoomTypeID,
		payload.Floor,
		payload.Number,
		payload.Name,
		payload.Status,
	)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	resp := NewRoomResponse(hotel)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *HotelController) handleDeleteHotelRoom(w http.ResponseWriter, r *http.Request) {

}
