interface User {
    id: number;
    avatar_url: string;
    username: string;
    tracking_since: string;
    user_stats: UserStats;
    user_map_counts: UserMapCounts;
}

interface UserMapCounts {
    graveyard: number;
    wip: number;
    pending: number;
    ranked: number;
    approved: number;
    qualified: number;
    loved: number;
}

interface UserStatistics {
    [key: string]: UserStatisticUnit;
}

interface UserStatisticUnit {
    [key: string]: number;
}

interface UserStats {
    [key: string]: UserStatsModel;
}

interface UserStatsModel {
    play_count: number;
    favourite_count: number;
    map_count: number;
    comments_count: number;
}

interface UserStatsDataset {
    timestamp: string;
    play_count: number;
    favourite_count: number;
    map_count: number;
    comments_count: number;
}