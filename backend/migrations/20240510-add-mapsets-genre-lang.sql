-- +migrate Up
ALTER TABLE mapsets ADD COLUMN genre text;
ALTER TABLE mapsets ADD COLUMN language text;

-- +migrate Down

ALTER TABLE mapsets DROP COLUMN genre;
ALTER TABLE mapsets DROP COLUMN language;