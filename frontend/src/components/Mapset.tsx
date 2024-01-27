import { useEffect, useState } from "react";
import { getRemainingPendingTime } from "../utils/utils";
import aveta from 'aveta';

interface MapCardProps {
    map: Mapset
}

// todo: refactor this crap
const Mapset = (props: MapCardProps) => {

    const [lastStats, setLastStats] = useState<MapsetStatsModel | null>(null);
    const [penultimateStats, setPenultimateStats] = useState<MapsetStatsModel | null>(null);
    const [playCountDifference, setPlayCountDifference] = useState<number>(0);
    const [favouriteCountDifference, setFavouriteCountDifference] = useState<number>(0);

    useEffect(() => {
        const mapsetStatsValues = Object.values(props.map.mapset_stats);
        const statsCount = mapsetStatsValues.length;

        if (statsCount > 1) {
            const lastStats = mapsetStatsValues[statsCount - 1];
            const penultimateStats = mapsetStatsValues[statsCount - 2];

            const newPlayCountDifference = lastStats.play_count - penultimateStats.play_count;
            const newFavouriteCountDifference = lastStats.favourite_count - penultimateStats.favourite_count;

            setPenultimateStats(penultimateStats);
            setPlayCountDifference(newPlayCountDifference);
            setFavouriteCountDifference(newFavouriteCountDifference);
            setLastStats(lastStats);
        }
    }, [props.map.mapset_stats]);

    return (
        <>
            {lastStats != null ?
                <div className="flex bg-zinc-900 text-white w-full rounded-lg overflow-hidden">
                    <div>
                        <img src={props.map.covers.card} className='h-full w-64 min-w-64' alt="map bg"
                            style={{ objectFit: 'cover' }} />
                    </div>
                    <div className="flex flex-col p-2 w-full">
                        <a className="text-xl"
                            href={`/beatmapset/${props.map.id}`}>{props.map.artist} - {props.map.title}</a>
                        <div className="flex gap-2 justify-left items-baseline">
                            <h1 className="text-xl text-green-200">{aveta(lastStats?.play_count) ?? 0} plays</h1>
                            {penultimateStats ?
                                <h1 className="text-sm h-full text-pink-200">{aveta(lastStats.favourite_count)} favourites</h1>
                                : null
                            }
                        </div>
                        <div className='text-xs text-zinc-400'>
                            {props.map.status == "wip" || props.map.status == "pending" ?
                                getRemainingPendingTime(props.map.last_updated)
                                : props.map.status}
                        </div>
                    </div>
                    <div className="px-4 flex flex-col gap-1 justify-center items-center">
                        {penultimateStats ?
                            <>
                                {favouriteCountDifference != 0 ?
                                    <div className="flex gap-2 items-center w-full justify-end">
                                        <h1 className="text-xs text-pink-400">▲</h1>
                                        <h1 className="text-2xl text-pink-400">{favouriteCountDifference}</h1>
                                    </div>
                                    : null
                                }
                                {playCountDifference != 0 ?
                                    <div className="flex gap-2 items-center w-full justify-end">
                                        <h1 className="text-xs text-green-300">▲</h1>
                                        <h1 className="text-2xl text-green-300">{playCountDifference}</h1>
                                    </div>
                                    : null
                                }
                            </>
                            : null
                        }
                    </div>
                </div>
                : null}
        </>
    );
};

export default Mapset;