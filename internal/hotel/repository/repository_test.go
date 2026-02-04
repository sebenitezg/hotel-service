package repository

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"

	"github.com/sebenitezg/hotel-service/config"
	"github.com/sebenitezg/hotel-service/pkg/db"
	"github.com/sebenitezg/hotel-service/internal/hotel/models"
)

func TestHotelSave(t *testing.T) {

	database := db.NewConnection(config.DatabaseConfigurations{
		Host:       "localhost",
		Port:       5432,
		DbName:     "postgres",
		Username:   "postgres",
		Password:   "postgres",
		PoolSize:   2,
		LogQueries: false,
	})
	repository := NewRepository(database)

	now := time.Now().UTC()
	id, _ := uuid.NewV6()
	hotel := &models.Hotel{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Name:        "test-hotel",
		Address:     "test-address",
		Country:     "test-country",
		State:       "test-state",
		Status:      "enabled",
		Description: "this is a great hotel",
	}

	repository.Save(hotel)

	hotels := make([]models.Hotel, 0, 6)
	err := database.
		NewRaw("select * from ? where id=?", bun.Ident("hotels"), hotel.ID).
		Scan(context.Background(), &hotels)
	if err != nil {
		t.Errorf("failure scaning the database: %v", err)
	}

	assert.Equal(t, len(hotels), 1)
	assert.Equal(t, hotel.ID, hotels[0].ID)
	assert.ObjectsAreEqual(hotel.CreatedAt, hotels[0].CreatedAt)
	assert.ObjectsAreEqual(hotel.UpdatedAt, hotels[0].UpdatedAt)
	assert.Equal(t, hotel.Name, hotels[0].Name)
	assert.Equal(t, hotel.Address, hotels[0].Address)
	assert.Equal(t, hotel.Country, hotels[0].Country)
	assert.Equal(t, hotel.State, hotels[0].State)
	assert.Equal(t, hotel.Status, hotels[0].Status)
	assert.Equal(t, hotel.Description, hotels[0].Description)
}
