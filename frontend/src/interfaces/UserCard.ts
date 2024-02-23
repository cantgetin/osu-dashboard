interface UserCard {
    User: User;
    Mapsets: Mapset[];
}

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

interface UserStats {
    [key: string]: UserStatsModel;
}

interface UserStatsModel {
    play_count: number;
    favourite_count: number;
    map_count: number;
    comment_count: number;
}

interface UserStatsDataset {
    timestamp: string;
    play_count: number;
    favourite_count: number;
    map_count: number;
    comments_count: number;
}

interface Mapset {
    id: number;
    artist: string;
    title: string;
    covers: { [key: string]: string };
    status: string;
    last_updated: string; // Assuming time.Time is serialized as string in JSON
    user_id: number;
    preview_url: string;
    tags: string;
    mapset_stats: MapsetStats;
    bpm: number;
    creator: string;
    beatmaps: Beatmap[];
}

interface MapsetStats {
    [key: string]: MapsetStatsModel;
}

interface MapsetStatsModel {
    play_count: number;
    favourite_count: number;
}

interface Beatmap {
    id: number;
    beatmapset_id: number;
    difficulty_rating: number;
    version: string;
    accuracy: number;
    ar: number;
    bpm: number;
    cs: number;
    status: string;
    url: string;
    total_length: number;
    user_id: number;
    beatmap_stats: BeatmapStats;
    last_updated: string; // Assuming time.Time is serialized as string in JSON
}

interface BeatmapStats {
    [key: string]: BeatmapStatsModel;
}

interface BeatmapStatsModel {
    play_count: number;
    pass_count: number;
}