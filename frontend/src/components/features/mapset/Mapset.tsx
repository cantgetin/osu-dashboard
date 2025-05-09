import React, {useEffect, useState} from "react";
import aveta from "aveta";
import StatsDifference from "../stats/StatsDifference.tsx";
import {getRemainingPendingTime} from "../../../utils/utils.ts";
import {FaFileExcel} from "react-icons/fa";

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
                <div
                    className={`flex flex-col sm:flex-row bg-zinc-800 bg-opacity-30 text-white w-full 
                overflow-hidden rounded-lg ${props.className}`}>
                    <div className="sm:w-64 sm:min-w-64 bg-zinc-700 flex items-center justify-center">
                        {props.map.covers.card ? (
                            <img
                                src={props.map.covers.card}
                                className='h-36 sm:h-full w-full sm:w-64 object-cover'
                                alt={`${props.map.title} cover`}
                                onError={(e: React.SyntheticEvent<HTMLImageElement>) => {
                                    const imgElement = e.currentTarget;  // This is properly typed as HTMLImageElement
                                    imgElement.style.display = 'none';

                                    const fallbackElement = imgElement.nextElementSibling as HTMLElement | null;
                                    if (fallbackElement) {
                                        fallbackElement.style.display = 'flex';
                                    }
                                }}
                            />
                        ) : null}
                        <div className="hidden flex-col items-center justify-center text-zinc-400 h-36 sm:h-full w-full sm:w-64">
                            <FaFileExcel className="text-4xl" />
                            <span className="mt-2 text-sm">No cover image</span>
                        </div>
                    </div>
                    <div className="flex flex-col p-4 w-full gap-1">
                        <a className="hover:text-amber-200" href={`/beatmapset/${props.map.id}`}>
                            <div className="text-lg line-clamp-1">{props.map.title}</div>
                            <div className="text-md text-zinc-400 line-clamp-1">by {props.map.artist}</div>
                        </a>
                        <div className="mt-2">
                            <div className="flex flex-wrap gap-x-2 gap-y-1 items-baseline">
                                <h1 className="text-sm sm:text-md text-green-200">
                                    {aveta(lastStats.play_count)} plays
                                </h1>
                                <h1 className="text-xs sm:text-sm text-pink-200">
                                    {aveta(lastStats.favourite_count)} favorites
                                </h1>
                                <h1 className="text-xs sm:text-sm text-red-400">
                                    {aveta(lastStats.comments_count)} comments
                                </h1>
                            </div>
                            <div className="flex flex-wrap gap-x-2 gap-y-1 items-center mt-1">
                                <h1 className="text-xs sm:text-sm text-zinc-400">
                                    {props.map.status == "wip" || props.map.status == "pending" ?
                                        getRemainingPendingTime(props.map.last_updated)
                                        : props.map.status}
                                </h1>
                                {props.showMapper &&
                                    <div className="text-xs sm:text-sm flex gap-1">
                                        <span>mapped by</span>
                                        <a
                                            href={`/user/${props.map.user_id}`}
                                            className="text-blue-300 hover:text-yellow-200"
                                        >
                                            {props.map.creator}
                                        </a>
                                    </div>
                                }
                            </div>
                        </div>
                    </div>
                    {penultimateStats &&
                        <div
                            className="px-4 py-2 sm:py-0 flex flex-row sm:flex-col
                        justify-center items-center gap-2 sm:gap-1 border-t sm:border-t-0 border-zinc-700">
                            <StatsDifference difference={favouriteCountDifference} className="text-pink-300"/>
                            <StatsDifference difference={playCountDifference} className="text-green-300"/>
                            <StatsDifference difference={commentsCountDifference} className="text-red-300"/>
                        </div>
                    }
                </div>
                : null}
        </>
    );
};

export default Mapset;