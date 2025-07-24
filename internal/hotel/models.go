package hotel

import (
	"time"

	"github.com/gofrs/uuid/v5"
	"github.com/uptrace/bun"
)

// --------------------
// Hotel DB models
// --------------------
type Hotel struct {
	bun.BaseModel `bun:"table:hotels"`
	ID            uuid.UUID `bun:"id"`
	CreatedAt     time.Time `bun:"created_at"`
	UpdatedAt     time.Time `bun:"updated_at"`
	Name          string    `bun:"name"`
	Address       string    `bun:"address"`
	Country       string    `bun:"country"`
	State         string    `bun:"state"`
	Status        string    `bun:"status"`
	Description   string    `bun:"description"`
}

type Hotels []Hotel

func NewHotel(
	name string,
	address string,
	country string,
	state string,
	status string,
	description string,
) (*Hotel, error) {
	now := time.Now().UTC()

	id, err := uuid.NewV6()
	if err != nil {
		return nil, err
	}

	return &Hotel{
		ID:          id,
		CreatedAt:   now,
		UpdatedAt:   now,
		Name:        name,
		Address:     address,
		Country:     country,
		State:       state,
		Status:      status,
		Description: description,
	}, nil
}
