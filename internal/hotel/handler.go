package hotel

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/sebenitezg/hotel-service/pkg/logger"
	"github.com/sebenitezg/hotel-service/pkg/server/rest"

	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"github.com/gofrs/uuid/v5"
	"go.uber.org/zap"
)

type HotelController struct {
	validator    *validator.Validate
	hotelService *HotelService
	log          *zap.SugaredLogger
}

func NewController(
	server *rest.HTTPServer,
	validator *validator.Validate,
	hotelService *HotelService,
) *HotelController {
	c := &HotelController{
		validator:    validator,
		hotelService: hotelService,
		log:          logger.GetLogger(),
	}

	server.Router.Group(func(r chi.Router) {
		r.Get("/v1/hotels/{hotel_id}", c.handleGetHotel)
		r.Get("/v1/hotels/", c.handleListHotels)
		r.Post("/v1/hotels/", c.handleCreateHotel)
		r.Patch("/v1/hotels/{hotel_id}", c.handlePartialUpdateHotel)
		r.Delete("/v1/hotels/{hotel_id}", c.handleDeleteHotel)
	})

	return c
}

func (c *HotelController) handleListHotels(w http.ResponseWriter, r *http.Request) {
	hotels, err := c.hotelService.ListHotels()
	if err != nil {
		rest.RenderError(r.Context(), w, err)
	}
	resp := NewListHotelsResponse(hotels)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *HotelController) handleGetHotel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "hotel_id")
	uuidID, err := uuid.FromString(id)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", id, "error", err)
		rest.RenderError(r.Context(), w, errors.New("requested hotel does not exist"))
		return
	}

	hotel, err := c.hotelService.GetHotelByID(uuidID)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
	}

	resp := NewHotelResponse(hotel)

	rest.RenderJSON(r.Context(), w, http.StatusOK, resp)
}

func (c *HotelController) handleCreateHotel(w http.ResponseWriter, r *http.Request) {
	var payload CreateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.log.Errorw("failed to decode request body", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}
	if err := c.validator.Struct(payload); err != nil {
		c.log.Errorw("validation failure", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}

	hotel, err := NewHotel(
		payload.Name,
		payload.Address,
		payload.Country,
		payload.State,
		payload.Status,
		payload.Description,
	)
	if err != nil {
		c.log.Errorw("failure creating hotel model instance", "error", err)
		rest.RenderError(r.Context(), w, err)
		return
	}

	hotel, err = c.hotelService.CreateHotel(hotel)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
		return
	}

	resp := NewHotelResponse(hotel)

	rest.RenderJSON(r.Context(), w, http.StatusCreated, resp)
}

func (c *HotelController) handlePartialUpdateHotel(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "hotel_id")
	uuidID, err := uuid.FromString(id)
	if err != nil {
		c.log.Errorw("invalid hotel id", "hotelID", id, "error", err)
		rest.RenderError(r.Context(), w, errors.New("requested hotel does not exist"))
		return
	}

	var payload UpdateHotelRequest
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		c.log.Errorw("failed to decode request body", "error", err)
		rest.RenderError(r.Context(), w, err)
	}

	hotel, err := c.hotelService.UpdatePartiallyHotel(
		uuidID,
		payload.Name,
		payload.Address,
		payload.Status,
		payload.Description,
	)
	if err != nil {
		rest.RenderError(r.Context(), w, err)
	}

	resp := NewHotelResponse(hotel)

	rest.RenderJSON(r.Context(), w, http.StatusAccepted, resp)

}

func (c *HotelController) handleDeleteHotel(w http.ResponseWriter, r *http.Request) {

}
