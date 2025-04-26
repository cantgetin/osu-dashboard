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
import {LoadingState} from "../interfaces/LoadingState.ts";
import Layout from "../components/ui/Layout.tsx";
import {fetchUser, selectUser, selectUserLoadingState} from "../store/userSlice.ts";
import UserDiagrams from "../components/business/UserDiagrams.tsx";

const UserPage = () => {
    const {userId} = useParams();

    const dispatch = useAppDispatch();
    const user = useAppSelector<User>(selectUser)
    const userLoaded = useAppSelector<LoadingState>(selectUserLoadingState)

    useEffect(() => {
        dispatch(fetchUser(Number(userId)))
    }, [dispatch, userId])

    const extLinkOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    return (
        <Layout className="flex md:justify-center sm:justify-start" title={user ? user.username : "Loading..."}>
            {userLoaded == LoadingState.Succeeded ?
                <div className="w-[1152px] grid 2xl:grid-cols-1 l:grid-cols-1 gap-4">
                    <User
                        user={user}
                        externalLinkOnClick={() => extLinkOnClick(user.id)}
                    >
                        <MapStatsSummary user={user}/>
                        <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                    </User>
                    <div className="flex gap-4">
                        <UserCharts
                            className="p-4 w-1/2"
                            data={mapUserStatsToArray(user.user_stats)}
                        />
                        <UserDiagrams
                            className="p-4 w-1/2"
                            userId={user.id}
                        />
                    </div>
                    <MapsetList
                        userId={user.id}
                        forUser={true}
                        page={1}
                        sort="last_playcount"
                        direction="desc"
                    />
                </div>
                : <LoadingSpinner/>}
        </Layout>
    );
};

export default UserPage;