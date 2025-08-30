import Tags from "../common/Tags.tsx";
import {useAppSelector} from "@/store/hooks.ts";
import {selectUserStats, selectUserStatsLoading,} from "@/store/userStatsSlice.ts";
import {LoadingState} from "@/interfaces/LoadingState.ts";
import {useFetchUserStatsOnce} from "@/hooks/useFetchUserStats.ts";

interface UserTagsProps {
    userID: string;
}

const UserTags = (props: UserTagsProps) => {
    const userStats = useAppSelector<UserStatistics | null>(selectUserStats)
    const userStatsLoaded = useAppSelector<LoadingState>(selectUserStatsLoading)

    useFetchUserStatsOnce(props.userID.toString())

    return (
        <>
            {userStatsLoaded == LoadingState.Succeeded ?
                <Tags tags={userStats!.combined.filter(item => item !== "")} colorized={true}/>
                :   <Tags tags={["lol","tag1"]} colorized={true}/>
            }
        </>
    );
};

export default UserTags;