import {formatDateDiff} from "../utils/utils.ts";
import aveta from "aveta";
import StatsDifference from "./StatsDifference.tsx";

interface UserStatsSummaryProps {
    data: UserStatsDataset[]
}

// todo: refactor
const UserStatsSummary = (props: UserStatsSummaryProps) => {
    return (
        <>
            <h1 className="text-2xl"> {props.data[props.data.length - 1].map_count} Maps fetched</h1>
            <div className="flex flex-col mt-auto ml-auto">
                {props.data.length > 1 ?
                    <>
                        <div className="text-xs text-zinc-400 px-2">
                            stats for last {formatDateDiff(props.data[props.data.length - 1].timestamp,
                            props.data[props.data.length - 2].timestamp)}
                        </div>
                        <div className="flex gap-2 justify-center items-center ml-auto px-2">
                            <div className="text-xl text-pink-300">
                                {aveta(props.data[props.data.length - 1].favourite_count)} Favorites
                            </div>
                            <StatsDifference
                                difference={props.data[props.data.length - 1].favourite_count - props.data[props.data.length - 2].favourite_count}
                                className="text-pink-300"
                                forceShowDiff={true}
                            />
                        </div>
                        <div className="flex gap-2 justify-center items-center ml-auto px-2">
                            <div className="text-xl text-green-300">
                                {aveta(props.data[props.data.length - 1].play_count)} Plays
                            </div>
                            <StatsDifference
                                difference={props.data[props.data.length - 1].play_count - props.data[props.data.length - 2].play_count}
                                className="text-green-300"
                                forceShowDiff={true}
                            />
                        </div>
                        <div className="flex gap-2 justify-center items-center ml-auto px-2">
                            <div className="text-xl text-red-300">
                                {aveta(props.data[props.data.length - 1].comments_count)} Comments
                            </div>
                            <StatsDifference
                                difference={props.data[props.data.length - 1].comments_count - props.data[props.data.length - 2].comments_count}
                                className="text-red-300"
                                forceShowDiff={true}
                            />
                        </div>
                    </>
                    : null}
            </div>
        </>
    );
};

export default UserStatsSummary;