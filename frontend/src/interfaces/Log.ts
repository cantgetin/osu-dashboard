interface Log {
    id?: number;
    name?: string;
    message?: string;
    service?: string;
    app_version?: string;
    platform?: string;
    type?: string;
    api_requests?: number;
    success_rate_percent?: number;
    tracked_at: string; // ISO 8601 string from time.Time
    avg_response_time?: number; // nanos from time.Duration
    elapsed_time?: number; // nanos from time.Duration
    time_since_last_track?: number; // nanos from time.Duration
}