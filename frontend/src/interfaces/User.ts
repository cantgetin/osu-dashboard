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
    most_popular_tags: UserStatisticUnit;
    most_popular_languages: UserStatisticUnit;
    most_popular_genres: UserStatisticUnit;
    most_popular_bpms: UserStatisticUnit;
    most_popular_starrates: UserStatisticUnit;
    combined: string[];
}

interface UserStatisticUnit {
    [key: string]: number;
}

interface UserStats {
    [key: string]: UserStatsModel;
}

// idk what is this below

interface UserStatsModel {
    play_count: number;
    favorite_count: number;
    map_count: number;
    comments_count: number;
}

interface UserStatsDataset {
    timestamp: string;
    play_count: number;
    favorite_count: number;
    map_count: number;
    comments_count: number;
}