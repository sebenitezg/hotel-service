package hotel

import (
	"context"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"
)

type HotelRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *HotelRepository {
	return &HotelRepository{
		db: db,
	}
}

func (r *HotelRepository) Save(hotel *Hotel) error {
	_, err := r.db.NewInsert().Model(hotel).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) UpdateHotel(hotel *Hotel) error {
	_, err := r.db.NewUpdate().Model(hotel).Where("id = ?", hotel.ID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) DeleteHotel(id string) error {
	_, err := r.db.NewDelete().Model((*Hotel)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *HotelRepository) GetAllHotels() (Hotels, error) {
	var hotels Hotels
	err := r.db.NewSelect().Model(&hotels).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return hotels, nil
}

func (r *HotelRepository) GetHotelByID(id uuid.UUID) (*Hotel, error) {
	var hotel Hotel
	err := r.db.NewSelect().Model(&hotel).Where("id = ?", id).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &hotel, nil
}
