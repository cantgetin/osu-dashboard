import {useParams} from "react-router-dom";
import {useEffect} from "react";
import User from "../components/business/User.tsx";
import UserStatsSummary from "../components/business/UserStatsSummary.tsx";
import UserCharts from "../components/business/UserCharts.tsx";
import MapsetList from "../components/business/MapsetList.tsx";
import MapStatsSummary from "../components/business/MapStatsSummary.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {selectLoadingState, selectUserCard} from "../store/userCardSlice.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";
import Layout from "../components/ui/Layout.tsx";
import {fetchUser} from "../store/userSlice.ts";

const UserPage = () => {
    const {userId} = useParams();

    const dispatch = useAppDispatch();

    const userCard = useAppSelector<UserCard>(selectUserCard);
    const loaded = useAppSelector<LoadingState>(selectLoadingState)

    useEffect(() => {
        dispatch(fetchUser(Number(userId)))
    }, [dispatch, userId])

    const extLinkOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    return (
        <Layout className="flex md:justify-center sm:justify-start">
            {loaded == LoadingState.Succeeded ?
                <div className="w-[1152px] grid 2xl:grid-cols-1 l:grid-cols-1 gap-4">
                    <User
                        user={userCard.User}
                        externalLinkOnClick={() => extLinkOnClick(userCard.User.id)}
                    >
                        <MapStatsSummary user={userCard.User}/>
                        <UserStatsSummary data={mapUserStatsToArray(userCard.User.user_stats)}/>
                    </User>
                    <UserCharts
                        className="p-4"
                        data={mapUserStatsToArray(userCard.User.user_stats)}/>
                    <MapsetList userId={userId!}/>
                </div>
                : <LoadingSpinner/>}
        </Layout>
    );
};

export default UserPage;