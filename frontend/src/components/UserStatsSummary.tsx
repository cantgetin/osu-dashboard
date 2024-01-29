import {formatDateDiff} from "../utils/utils.ts";
import aveta from "aveta";

interface UserStatsSummaryProps {
    data: UserStatsDataset[]
}

const UserStatsSummary = (props: UserStatsSummaryProps) => {
    return (
        <>
            <h1 className="text-2xl"> {props.data[props.data.length - 1].map_count} Maps fetched</h1>
            {/*<div>*/}
            {/*    <div className="text-xl text-yellow-200">{props.data[props.data.length - 1].play_count} Plays</div>*/}
            {/*    <div className="text-l text-pink-100">{props.data[props.data.length - 1].favourite_count} Favorites</div>*/}
            {/*</div>*/}
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
                            <h1 className="text-xs text-pink-400">▲</h1>
                            <h1 className="text-2xl text-pink-400">
                                {aveta(props.data[props.data.length - 1].favourite_count -
                                    props.data[props.data.length - 2].favourite_count)}
                            </h1>
                        </div>
                        <div className="flex gap-2 justify-center items-center ml-auto px-2">
                            <div className="text-xl text-green-200">
                                {aveta(props.data[props.data.length - 1].play_count)} Plays
                            </div>
                            <h1 className="text-xs text-green-300">▲</h1>
                            <h1 className="text-2xl text-green-300">
                                {aveta(props.data[props.data.length - 1].play_count -
                                    props.data[props.data.length - 2].play_count)}
                            </h1>
                        </div>
                    </>
                    : null}
            </div>
        </>
    );
};

export default UserStatsSummary;