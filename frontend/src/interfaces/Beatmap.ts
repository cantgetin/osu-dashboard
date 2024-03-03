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