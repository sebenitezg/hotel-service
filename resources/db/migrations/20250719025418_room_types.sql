-- migrate:up
CREATE TABLE public.room_types (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    hotel_id UUID NOT NULL REFERENCES hotels(id),
    name VARCHAR(128) NOT NULL,
    description TEXT NOT NULL,
    number_of_beds INTEGER NOT NULL,
    bed_type VARCHAR(64) NOT NULL,
    max_occupancy INTEGER NOT NULL,
    base_price DECIMAL(10, 4) NOT NULL
)

-- migrate:down
DROP TABLE public.room_types
