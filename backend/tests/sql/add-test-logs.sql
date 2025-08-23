INSERT INTO log (
    name,
    message,
    service,
    app_version,
    platform,
    type,
    api_requests,
    success_rate_percent,
    tracked_at,
    avg_response_time,
    elapsed_time,
    time_since_last_track
)
SELECT
    'LogEntry_' || (ROW_NUMBER() OVER ()) as name,
    'Log message number ' || (ROW_NUMBER() OVER ()) || ' with random content' as message,
    (ARRAY['auth-service', 'payment-service', 'user-service', 'notification-service', 'analytics-service'])[floor(random() * 5) + 1] as service,
    (ARRAY['1.0.0', '1.1.0', '2.0.0', '2.1.0', '3.0.0'])[floor(random() * 5) + 1] as app_version,
    (ARRAY['web', 'ios', 'android', 'desktop', 'api'])[floor(random() * 5) + 1] as platform,
    (ARRAY['single', 'regular'])[floor(random() * 2) + 1] as type,
    floor(random() * 1000) + 1 as api_requests,
    floor(random() * 100) + 1 as success_rate_percent,
    NOW() - (random() * INTERVAL '30 days') as tracked_at,
    floor(random() * 1000000000) + 1000000 as avg_response_time, -- 1ms to 1s in nanoseconds
    floor(random() * 5000000000) + 1000000 as elapsed_time, -- 1ms to 5s in nanoseconds
    floor(random() * 60000000000) + 1000000000 as time_since_last_track -- 1s to 60s in nanoseconds
FROM generate_series(1, 100);