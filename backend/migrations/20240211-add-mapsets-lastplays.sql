-- +migrate Up
ALTER TABLE mapsets ADD COLUMN last_playcount integer not null default 0;

-- +migrate Down

ALTER TABLE mapsets DROP COLUMN last_playcount;