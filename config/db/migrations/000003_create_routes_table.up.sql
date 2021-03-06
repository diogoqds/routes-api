CREATE TABLE routes (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) UNIQUE NOT NULL,
    bounds          GEOMETRY(Polygon, 4326),
    seller_id       INTEGER REFERENCES sellers (id),
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    deleted_at      TIMESTAMP DEFAULT NULL
);

