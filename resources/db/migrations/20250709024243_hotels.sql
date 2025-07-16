-- migrate:up
CREATE TABLE public.hotels (
    id UUID NOT NULL,
    created_at TIMESTAMPTZ NOT NULL,
    updated_at TIMESTAMPTZ NOT NULL,
    name VARCHAR(128) NOT NULL,
    address VARCHAR(256) NOT NULL,
    country VARCHAR(64) NOT NULL,
    state VARCHAR(64) NOT NULL,
    status VARCHAR(64) NOT NULL,
    PRIMARY KEY (id)
)

-- migrate:down
DROP TABLE public.hotels
