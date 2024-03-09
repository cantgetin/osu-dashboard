import {useParams} from "react-router-dom";
import {useEffect} from "react";
import User from "../components/User.tsx";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import UserCharts from "../components/UserCharts.tsx";
import MapsetList from "../components/MapsetList.tsx";
import MapStatsSummary from "../components/MapStatsSummary.tsx";
import {extractUserMapsCountFromStats, mapUserStatsToArray} from "../utils/utils.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {fetchUserCard, selectLoadingState, selectUserCard} from "../store/userCardSlice.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";
import Layout from "../components/ui/Layout.tsx";

const UserPage = () => {
    const {userId} = useParams();

    const dispatch = useAppDispatch();

    const userCard = useAppSelector<UserCard>(selectUserCard);
    const loaded = useAppSelector<LoadingState>(selectLoadingState)

    useEffect(() => {
        dispatch(fetchUserCard({userId: Number(userId), page: 1}))
    }, [dispatch, userId])

    const userNameOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    return (
        <Layout className="flex md:justify-center sm:justify-start">
            {loaded == LoadingState.Succeeded ?
                <div className="w-[1152px] grid 2xl:grid-cols-1 l:grid-cols-1 gap-4">
                    <User user={userCard.User} nameOnClick={() => userNameOnClick(userCard.User.id)}>
                        <MapStatsSummary user={userCard.User}/>
                        <UserStatsSummary data={mapUserStatsToArray(userCard.User.user_stats)}/>
                    </User>
                    <UserCharts
                        className="p-4"
                        data={mapUserStatsToArray(userCard.User.user_stats)}/>
                    <MapsetList
                        Mapsets={userCard.Mapsets}
                        MapsetCount={extractUserMapsCountFromStats(userCard.User.user_stats)}
                        userId={userId!}
                    />
                </div>
                : <LoadingSpinner/>}
        </Layout>
    );
};

export default UserPage;