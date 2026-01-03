package repository

import (
	"context"
	"database/sql"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"

	"github.com/sebenitezg/hotel-service/internal/room/models"
)

type RoomRepository struct {
	db *bun.DB
}

func NewRepository(db *bun.DB) *RoomRepository {
	return &RoomRepository{
		db: db,
	}
}

func (r *RoomRepository) Save(room *models.Room) error {
	_, err := r.db.NewInsert().Model(room).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) Update(room *models.Room) error {
	_, err := r.db.NewUpdate().Model(room).Where("id = ?", room.ID).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) Delete(id int64) error {
	_, err := r.db.NewDelete().Model((*models.Room)(nil)).Where("id = ?", id).Exec(context.Background())
	if err != nil {
		return err
	}
	return nil
}

func (r *RoomRepository) GetAll() ([]models.Room, error) {
	var rooms []models.Room
	err := r.db.NewSelect().Model(&rooms).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return rooms, nil
}

func (r *RoomRepository) GetByID(id uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := r.db.NewSelect().Model(&room).Where("id = ?", id).Scan(context.Background())
	if err == sql.ErrNoRows {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) GetByHotelRoomID(hotelID, roomID uuid.UUID) (*models.Room, error) {
	var room models.Room
	err := r.db.NewSelect().Model(&room).Where("hotel_id = ? and id = ?", hotelID, roomID).Scan(context.Background())
	if err != nil {
		return nil, err
	}
	return &room, nil
}

func (r *RoomRepository) GetByHotelID(hotelID uuid.UUID) (models.Rooms, error) {
	var rooms models.Rooms
	err := r.db.NewSelect().
		Model(&rooms).
		Where("hotel_id = ?", hotelID).
		Scan(context.Background())
	if err != nil {
		return nil, err
	}

	return rooms, nil
}
