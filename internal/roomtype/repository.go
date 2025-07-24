package roomtype

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"
)

type RoomTypeRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *RoomTypeRepository {
	return &RoomTypeRepository{
		db: db,
	}
}

func (r *RoomTypeRepository) Save(roomType *RoomType) error {
	_, err := r.db.NewInsert().
		Model(roomType).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomTypeRepository) Update(roomType *RoomType) error {
	_, err := r.db.NewUpdate().
		Model(roomType).
		Where("id = ?", roomType.ID).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomTypeRepository) Delete(id int64) error {
	_, err := r.db.NewDelete().
		Model((*RoomType)(nil)).
		Where("id = ?", id).
		Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomTypeRepository) GetAll() ([]RoomType, error) {
	var rooms []RoomType
	err := r.db.NewSelect().
		Model(&rooms).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomTypeRepository) GetByID(id uuid.UUID) (*RoomType, error) {
	var roomType RoomType
	err := r.db.NewSelect().
		Model(&roomType).
		Where("id = ?", id).
		Scan(context.Background())
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (r *RoomTypeRepository) GetByHotelRoomID(hotelID, roomTypeID uuid.UUID) (*RoomType, error) {
	var roomType RoomType
	err := r.db.NewSelect().
		Model(&roomType).
		Where("hotel_id = ? and id = ?", hotelID, roomTypeID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &roomType, nil
}

func (r *RoomTypeRepository) GetByHotelID(hotelID uuid.UUID) (RoomTypes, error) {
	var rooms RoomTypes
	err := r.db.NewSelect().
		Model(&rooms).
		Where("hotel_id = ?", hotelID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
