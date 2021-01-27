CREATE TABLE routes (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) NOT NULL,
    bounds          GEOMETRY(Polygon, 28992),
    seller_id       INTEGER REFERENCES sellers (id),
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    deleted_at      TIMESTAMP DEFAULT NULL
);

