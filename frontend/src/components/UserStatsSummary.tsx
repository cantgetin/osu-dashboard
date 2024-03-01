import aveta from "aveta";
import StatsDifference from "./StatsDifference.tsx";
import {useEffect, useState} from "react";

interface UserStatsSummaryProps {
    data: UserStatsDataset[];
}

const UserStatsSummary = (props: UserStatsSummaryProps) => {
    const [forLastNDays, setForLastNDays] = useState(1);
    const [favoritesDifference, setFavoritesDifference] = useState(0);
    const [commentsDifference, setCommentsDifference] = useState(0);
    const [playCountDifference, setPlayCountDifference] = useState(0);
    const lastIndex = props.data.length - 1;
    const lastData = props.data[lastIndex];

    const calculateDifference = (current: number, previous: number) => {
        return current - previous;
    };

    useEffect(() => {
        const previousIndex = lastIndex - forLastNDays;
        setFavoritesDifference(calculateDifference(lastData.favourite_count, props.data[previousIndex].favourite_count));
        setCommentsDifference(calculateDifference(lastData.comments_count, props.data[previousIndex].comments_count));
        setPlayCountDifference(calculateDifference(lastData.play_count, props.data[previousIndex].play_count));
    }, [forLastNDays, lastIndex, props.data]);

    const renderStatsDifference = (count: number, difference: number, name: string, color: string) => {
        return (
            <div className={`flex gap-2 justify-center items-center ml-auto px-2 text-xl text-${color}-300`}>
                {aveta(count)} {name}
                <StatsDifference difference={difference} className={`text-${color}-300`} forceShowDiff={true}/>
            </div>
        );
    };

    return (
        <>
            <h1 className="text-2xl"> {lastData.map_count} Maps fetched</h1>
            <div className="text-sm text-zinc-400 flex gap-2 justify-end">
                <h1>Stats for last </h1>
                <select
                    className="bg-zinc-800 text-zinc-400 rounded-md"
                    value={forLastNDays}
                    onChange={(e) => setForLastNDays(Number(e.target.value))}
                >
                    <option value="1">24 hours</option>
                    <option value="6">7 days</option>
                </select>
            </div>
            <div className="flex flex-col mt-auto ml-auto">
                {props.data.length > 1 && (
                    <>
                        {renderStatsDifference(lastData.favourite_count, favoritesDifference, "Favorites", "pink")}
                        {renderStatsDifference(lastData.play_count, playCountDifference, "Plays", "green")}
                        {renderStatsDifference(lastData.comments_count, commentsDifference, "Comments", "red")}
                    </>
                )}
            </div>
        </>
    );
};

export default UserStatsSummary;