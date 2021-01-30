CREATE TABLE clients (
    id              SERIAL PRIMARY KEY,
    name            VARCHAR(100) UNIQUE NOT NULL,
    geolocation     GEOMETRY(Point, 4326),
    route_id        INTEGER REFERENCES routes (id),
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW(),
    deleted_at      TIMESTAMP DEFAULT NULL
);

