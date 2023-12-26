export interface OsuMap {
    id: number
    artist: string
    title: string
    created: string
    covers: MapCovers
    beatmaps: Beatmap[]
    status: string
    last_updated: string
}

export interface MapCovers {
    card: string
    cover: string
}

export interface Beatmap {
    difficulty_rating: number
    version: string
    playcount: number
}