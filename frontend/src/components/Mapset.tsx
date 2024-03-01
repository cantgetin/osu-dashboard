import {useEffect, useState} from "react";
import {getRemainingPendingTime} from "../utils/utils";
import aveta from 'aveta';
import Button from "./Button.tsx";
import StatsDifference from "./StatsDifference.tsx";

interface MapCardProps {
    map: Mapset
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

    const mapsetExternalLinkOnClick = (mapsetId: number) => window.open(`https://osu.ppy.sh/beatmapsets/${mapsetId}`)

    return (
        <>
            {lastStats != null ?
                <div className="flex bg-zinc-900 text-white w-full rounded-lg overflow-hidden">
                    <div>
                        <img src={props.map.covers.card} className='h-full w-64 min-w-64' alt="map bg"
                             style={{objectFit: 'cover'}}/>
                    </div>
                    <div className="flex flex-col p-4 w-full">
                        <div className="flex gap-2 items-center">
                            <a className="text-xl hover:text-amber-200"
                               href={`/beatmapset/${props.map.id}`}>{props.map.artist} - {props.map.title}</a>
                            <Button className="rounded-md w-12 h-6 text-sm bg-zinc-800"
                                    onClick={() => mapsetExternalLinkOnClick(props.map.id)}
                                    content="osu!"
                                    keyNumber={1}
                            />
                        </div>
                        <div className="flex gap-2 justify-left items-baseline">
                            <h1 className="text-xl text-green-200">{aveta(lastStats.play_count)} plays</h1>
                            <h1 className="text-sm h-full text-pink-200">{aveta(lastStats.favourite_count)} favorites</h1>
                            <h1 className="text-sm h-full text-red-400">{aveta(lastStats.comments_count)} comments</h1>
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