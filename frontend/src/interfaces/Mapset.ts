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
    favorite_count: number;
    comments_count: number;
}

interface MapsetStatsDataset {
    timestamp: string;
    play_count: number;
    favorite_count: number;
    comments_count: number;
}