import {useEffect, useState} from "react";
import StatsDifference from "./StatsDifference.tsx";
import aveta from "aveta";
import {getRemainingPendingTime} from "../../utils/utils.ts";

interface MapCardProps {
    map: Mapset
    showMapper?: boolean
    className?: string
}

// todo: refactor this crap
const Mapset = (props: MapCardProps) => {

    const [lastStats, setLastStats] = useState<MapsetStatsModel | null>(null);
    const [penultimateStats, setPenultimateStats] = useState<MapsetStatsModel | null>(null);
    const [playCountDifference, setPlayCountDifference] = useState<number>(0);
    const [favouriteCountDifference, setFavouriteCountDifference] = useState<number>(0);
    const [commentsCountDifference, setCommentsCountDifference] = useState<number>(0);

    useEffect(() => {
        const mapsetStatsValues = Object.values(props.map.mapset_stats);
        const statsCount = mapsetStatsValues.length;
        setLastStats(mapsetStatsValues[statsCount - 1])

        if (statsCount > 1) {
            const penultimateStats = mapsetStatsValues[statsCount - 2];

            const newPlayCountDifference = mapsetStatsValues[statsCount - 1].play_count - penultimateStats.play_count;
            const newFavouriteCountDifference = mapsetStatsValues[statsCount - 1].favourite_count - penultimateStats.favourite_count;
            const newCommentsCountDifference = mapsetStatsValues[statsCount - 1].comments_count - penultimateStats.comments_count;

            setPenultimateStats(penultimateStats);
            setPlayCountDifference(newPlayCountDifference);
            setFavouriteCountDifference(newFavouriteCountDifference);
            setCommentsCountDifference(newCommentsCountDifference);
        }

    }, [props.map.mapset_stats]);

    return (
        <>
            {lastStats != null ?
                <div className={`flex bg-zinc-900 text-white w-full rounded-lg overflow-hidden ${props.className}`}>
                    <div>
                        <img src={props.map.covers.card} className='h-full w-64 min-w-64' alt="map bg"
                             style={{objectFit: 'cover'}}/>
                    </div>
                    <div className="flex flex-col p-4 w-full gap-1">
                        <a className="h-12 hover:text-amber-200" href={`/beatmapset/${props.map.id}`}>
                            <div className="h-6 text-lg">{props.map.title}</div>
                            <div className="h-6 text-md text-zinc-400">by {props.map.artist}</div>
                        </a>
                        <div>
                            <div className="flex gap-2 justify-left items-baseline">
                                <h1 className="text-md text-green-200">{aveta(lastStats.play_count)} plays</h1>
                                <h1 className="text-sm h-full text-pink-200">{aveta(lastStats.favourite_count)} favorites</h1>
                                <h1 className="text-sm h-full text-red-400">{aveta(lastStats.comments_count)} comments</h1>
                            </div>
                            <div className="flex gap-2 items-center">
                                <h1 className="text-sm text-zinc-400">
                                    {props.map.status == "wip" || props.map.status == "pending" ?
                                        getRemainingPendingTime(props.map.last_updated)
                                        : props.map.status}
                                </h1>
                                <div className="text-sm flex gap-1">
                                    {props.showMapper ?
                                        <>
                                            <h1>mapped by</h1>
                                            <a
                                                href={`/user/${props.map.user_id}`}
                                                className="text-blue-300 hover:text-yellow-200"
                                            >
                                                {props.map.creator}
                                            </a>
                                        </>

                                        : null}
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="px-4 flex flex-col justify-center items-center">
                        {penultimateStats ?
                            <>
                                <StatsDifference difference={favouriteCountDifference} className="text-pink-300"/>
                                <StatsDifference difference={playCountDifference} className="text-green-300"/>
                                <StatsDifference difference={commentsCountDifference} className="text-red-300"/>
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