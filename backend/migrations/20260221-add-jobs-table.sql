-- +migrate Up
CREATE TABLE jobs
(
    id         serial primary key,
    name       TEXT      NOT NULL,
    started_at timestamp NOT NULL DEFAULT NOW(),
    ended_at   timestamp NOT NULL DEFAULT NOW(),
    error      boolean   NOT NULL DEFAULT false,
    error_text text      NOT NULL DEFAULT ''
);

-- +migrate Down
DROP TABLE IF EXISTS jobs;