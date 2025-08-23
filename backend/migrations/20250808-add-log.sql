-- +migrate Up
CREATE TABLE IF NOT EXISTS log (
        id SERIAL PRIMARY KEY,
        name TEXT NOT NULL,
        message TEXT NOT NULL,
        service TEXT NOT NULL,
        app_version TEXT NOT NULL,
        platform TEXT NOT NULL,
        type TEXT NOT NULL CHECK (type IN ('single', 'regular')),
        api_requests INTEGER NOT NULL,
        success_rate_percent INTEGER NOT NULL,
        tracked_at TIMESTAMP,
        avg_response_time BIGINT NOT NULL, -- stored in nanoseconds
        elapsed_time BIGINT NOT NULL, -- stored in nanoseconds
        time_since_last_track BIGINT NOT NULL -- stored in nanoseconds
    );

-- +migrate Down
DROP TABLE log;