-- +goose Up
CREATE TABLE IF NOT EXISTS products (
    id  INTEGER NOT NULL PRIMARY KEY,
    name VARCHAR(255)
);

-- +goose Down
DROP TABLE IF EXISTS products;