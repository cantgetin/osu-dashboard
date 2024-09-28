-- +migrate Up
CREATE TABLE cleans
(
    id                         serial primary key,
    cleaned_at                 timestamp default NOW()
);

INSERT INTO cleans (id, cleaned_at)
VALUES (1, '2024-01-26 15:30:30');

-- +migrate Down

DROP TABLE tracks;