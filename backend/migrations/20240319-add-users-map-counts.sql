-- +migrate Up
ALTER TABLE users ADD COLUMN map_counts jsonb;

-- +migrate Down

ALTER TABLE users DROP COLUMN map_counts;