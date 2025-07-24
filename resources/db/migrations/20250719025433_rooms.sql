-- migrate:up
CREATE TABLE public.rooms (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    hotel_id UUID NOT NULL REFERENCES hotels(id),
    room_type_id UUID NOT NULL REFERENCES room_types(id),
    floor INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL,
    status VARCHAR(32) NOT NULL
)

-- migrate:down
DROP TABLE public.rooms
