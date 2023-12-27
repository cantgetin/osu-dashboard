interface UserCard {
    User: User;
    Mapsets: Mapset[];
}

interface User {
    id: number;
    avatar_url: string;
    username: string;
    unranked_beatmapset_count: number;
    graveyard_beatmapset_count: number;
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