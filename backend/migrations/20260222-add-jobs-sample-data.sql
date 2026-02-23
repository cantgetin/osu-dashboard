-- +migrate Up
INSERT INTO jobs (name, started_at, ended_at, error, error_text)
VALUES ('clean_stats', NOW(), NOW(), false, ''),
       ('enrich_data', NOW(), NOW(), false, ''),
       ('clean_users', NOW(), NOW(), false, ''),
       ('track_users', NOW(), NOW(), false, '');

-- +migrate Down
DELETE FROM jobs WHERE name IN ('clean_stats', 'enrich_data', 'track_users', 'clean_users');
