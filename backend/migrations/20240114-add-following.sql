-- +migrate Up

INSERT INTO following (id, username, created_at)
VALUES (7192129, 'Gasha', NOW()),
       (8985360, 'TheOnlyNEET', NOW()),
       (5371497, 'Alumetri', NOW()),
       (4452992, 'Sotarks', NOW()),
       (1722835, 'elchxyrlia', NOW());

-- +migrate Down