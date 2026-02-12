-- +migrate Up
-- +migrate StatementBegin

CREATE TABLE IF NOT EXISTS bioskop (
    id SERIAL PRIMARY KEY,
    nama VARCHAR(255) NOT NULL,
    lokasi VARCHAR(255) NOT NULL,
    rating FLOAT CHECK (rating >= 0 AND rating <= 5)
);

-- +migrate StatementEnd
