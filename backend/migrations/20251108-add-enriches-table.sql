-- +migrate Up
CREATE TABLE enriches
(
    id                         serial primary key,
    enriched_at                 timestamp default NOW()
);

INSERT INTO enriches (id, enriched_at)
VALUES (1, '2024-01-26 15:30:30');

-- +migrate Down

DROP TABLE enriches;