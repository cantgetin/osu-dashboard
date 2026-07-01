-- +migrate Up
ALTER TABLE users ADD COLUMN country_code text;
UPDATE users SET country_code = 'AU' WHERE country_code IS NULL;

-- +migrate Down

ALTER TABLE users DROP COLUMN country_code;