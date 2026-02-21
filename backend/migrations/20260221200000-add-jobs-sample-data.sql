-- +migrate Up
INSERT INTO jobs (name, started_at, ended_at, error, error_text)
VALUES ('clean stats', NOW(), NOW(), false, ''),
       ('enrich data', NOW(), NOW(), false, ''),
       ('track users', NOW(), NOW(), false, '');

-- +migrate Down
DELETE FROM jobs WHERE name IN ('clean stats', 'enrich data', 'track users');
