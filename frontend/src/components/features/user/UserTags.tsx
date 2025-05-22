import Tags from "../common/Tags.tsx";
import {useAppDispatch, useAppSelector} from "../../../store/hooks.ts";
import {
    fetchUserStats,
    selectUserStats,
    selectUserStatsLoading,
} from "../../../store/userStatsSlice.ts";
import {LoadingState} from "../../../interfaces/LoadingState.ts";
import {useEffect} from "react";
import LoadingSpinner from "../../ui/LoadingSpinner.tsx";

interface UserTagsProps {
    userID: string;
}

const UserTags = (props: UserTagsProps) => {
    const dispatch = useAppDispatch();
    const userStats = useAppSelector<UserStatistics | null>(selectUserStats)
    const userStatsLoaded = useAppSelector<LoadingState>(selectUserStatsLoading)

    useEffect(() => {
        dispatch(fetchUserStats(props.userID.toString()))
    }, [dispatch])

    return (
        <>
            {userStatsLoaded == LoadingState.Succeeded ?
                <Tags tags={userStats!.combined.filter(item => item !== "")} colorized={true}/>
                : <LoadingSpinner/>
            }
        </>
    );
};

export default UserTags;