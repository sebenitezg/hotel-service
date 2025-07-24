-- migrate:up
CREATE TABLE public.hotels (
    id UUID NOT NULL PRIMARY KEY,
    created_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT NOW(),
    name VARCHAR(128) NOT NULL,
    address VARCHAR(256) NOT NULL,
    country VARCHAR(64) NOT NULL,
    state VARCHAR(64) NOT NULL,
    status VARCHAR(64) NOT NULL,
    description TEXT NOT NULL
)

-- migrate:down
DROP TABLE public.hotels
