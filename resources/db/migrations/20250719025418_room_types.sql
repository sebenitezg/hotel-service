-- migrate:up
CREATE TABLE public.room_types (
    id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    hotel_id UUID NOT NULL,
    name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    number_of_beds INTEGER NOT NULL,
    bed_type VARCHAR(64) NOT NULL,
    max_occupancy INTEGER NOT NULL,
    base_price DECIMAL(10, 4) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
)

-- migrate:down
DROP TABLE public.room_types
