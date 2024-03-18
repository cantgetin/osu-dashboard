-- +migrate Up

INSERT INTO following (id, username, created_at)
VALUES (7192129, 'Gasha', NOW());

-- +migrate Down