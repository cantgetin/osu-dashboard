import aveta from "aveta";
import {useEffect, useState} from "react";
import StatsDifference from "../stats/StatsDifference.tsx";

interface UserStatsSummaryProps {
    data: UserStatsDataset[];
}

const UserStatsSummary = (props: UserStatsSummaryProps) => {
    const [forLastNDays, setForLastNDays] = useState(1);
    const [favoritesDiff, setFavoritesDiff] = useState(0);
    const [commentsDiff, setCommentsDiff] = useState(0);
    const [playCountDiff, setPlayCountDiff] = useState(0);
    const lastIndex = props.data.length - 1;
    const lastData = props.data[lastIndex];

    const calculateDiff = (current: number, previous: number) => {
        return current - previous;
    };

    useEffect(() => {
        const previousIndex = lastIndex - forLastNDays;

        if (props.data.length > 1 && previousIndex >= 0) {
            setFavoritesDiff(calculateDiff(lastData.favourite_count, props.data[previousIndex].favourite_count));
            setCommentsDiff(calculateDiff(lastData.comments_count, props.data[previousIndex].comments_count));
            setPlayCountDiff(calculateDiff(lastData.play_count, props.data[previousIndex].play_count));
        } else {
            // Reset Diffs when there's no previous data to compare with
            setFavoritesDiff(0);
            setCommentsDiff(0);
            setPlayCountDiff(0);
        }
    }, [forLastNDays, lastIndex, props.data]);

    const renderStatsDifference = (count: number, difference: number, name: string, color: string) => {
        return (
            <div className={`text-sm flex gap-2 justify-center items-center ml-auto px-2 md:text-xl text-${color}-300`}>
                {aveta(count)} {name}
                <StatsDifference difference={difference} className={`text-${color}-300`} forceShowDiff={true}/>
            </div>
        );
    };

    return (
        <>
            <h1 className="text-lg sm:text-xl"> {lastData.map_count} Maps fetched</h1>
            <div className="text-xs sm:text-sm text-zinc-400 flex gap-2 justify-end">
                <h1>Stats for last </h1>
                <select
                    className="bg-zinc-800 text-zinc-400 rounded-md text-xs sm:text-sm"
                    value={forLastNDays}
                    onChange={(e) => setForLastNDays(Number(e.target.value))}
                >
                    <option value="1">24 hours</option>
                    <option value="6">7 days</option>
                </select>
            </div>
            <div className="flex flex-col mt-auto ml-auto">
                {renderStatsDifference(lastData.favourite_count, favoritesDiff, "Favorites", "pink")}
                {renderStatsDifference(lastData.play_count, playCountDiff, "Plays", "green")}
                {renderStatsDifference(lastData.comments_count, commentsDiff, "Comments", "red")}
            </div>
        </>
    );
};

export default UserStatsSummary;