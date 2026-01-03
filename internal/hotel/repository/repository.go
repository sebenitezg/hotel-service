package repository

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"

	"github.com/sebenitezg/hotel-service/internal/hotel/models"
)

type HotelRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) Save(hotel *models.Hotel) error {
	_, err := r.db.NewInsert().Model(hotel).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) Update(hotel *models.Hotel) error {
	_, err := r.db.NewUpdate().Model(hotel).Where("id = ?", hotel.ID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) Delete(id string) error {
	_, err := r.db.NewDelete().Model((*models.Hotel)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) GetAll() (models.Hotels, error) {
	var hotels models.Hotels
	err := r.db.NewSelect().Model(&hotels).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *HotelRepository) GetByID(id uuid.UUID) (*models.Hotel, error) {
	var hotel models.Hotel
	err := r.db.NewSelect().Model(&hotel).Where("id = ?", id).Scan(context.Background())
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}
