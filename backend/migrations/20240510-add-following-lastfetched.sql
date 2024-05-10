-- +migrate Up
ALTER TABLE following ADD COLUMN last_fetched timestamp;
UPDATE following SET last_fetched = NOW() WHERE last_fetched IS NULL;

-- +migrate Down

ALTER TABLE following DROP COLUMN last_fetched;