-- +migrate Up
CREATE TABLE tracking
(
    id                         integer not null primary key,
    username                   text    not null,
    created_at                 timestamp default NOW()
);

-- +migrate Down
DROP TABLE tracking;