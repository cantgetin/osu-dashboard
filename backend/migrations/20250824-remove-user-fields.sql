-- +migrate Up
ALTER TABLE users
    DROP COLUMN graveyard_beatmapset_count,
    DROP COLUMN unranked_beatmapset_count;

-- +migrate Down
ALTER TABLE users
    ADD COLUMN graveyard_beatmapset_count INTEGER NOT NULL DEFAULT 0,
    ADD COLUMN unranked_beatmapset_count INTEGER NOT NULL DEFAULT 0;