-- +goose Up
ALTER TABLE skins ADD COLUMN skin_coefficient VARCHAR(100);

-- +goose Down
ALTER TABLE skins DROP COLUMN IF EXISTS skin_coefficient;