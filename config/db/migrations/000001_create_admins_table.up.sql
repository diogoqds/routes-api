CREATE TABLE admins (
    id              SERIAL PRIMARY KEY,
    email           VARCHAR(100) NOT NULL UNIQUE,
    created_at      TIMESTAMP DEFAULT NOW(),
    updated_at      TIMESTAMP DEFAULT NOW()
);

INSERT into admins (email) VALUES ("admin@email.com");