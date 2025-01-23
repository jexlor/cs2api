-- +goose Up
CREATE TABLE IF NOT EXISTS skins (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    weapon VARCHAR(255),
    rarity VARCHAR(255),
    collection VARCHAR(255),
    price VARCHAR(255),
    stattrack_price VARCHAR(255),
    url VARCHAR(255)
);

-- +goose Down
DROP TABLE IF EXISTS skins;