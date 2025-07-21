-- migrate:up
CREATE TABLE public.rooms (
    id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    hotel_id UUID NOT NULL,
    room_type_id UUID NOT NULL,
    floor INTEGER NOT NULL,
    number INTEGER NOT NULL,
    name VARCHAR(128) NOT NULL,
    status VARCHAR(32) NOT NULL,
    PRIMARY KEY (id),
    FOREIGN KEY (hotel_id) REFERENCES hotels(id)
    FOREIGN KEY (room_type_id) REFERENCES room_types(id)
)

-- migrate:down
DROP TABLE public.rooms
