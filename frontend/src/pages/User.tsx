import {useParams} from "react-router-dom";
import {useEffect} from "react";
import {mapUserStatsToArray} from "../utils/utils.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";
import Layout from "../components/ui/Layout.tsx";
import {fetchUser, selectUser, selectUserLoadingState} from "../store/userSlice.ts";
import User from "../components/features/user/User.tsx";
import MapStatsSummary from "../components/features/stats/MapStatsSummary.tsx";
import UserStatsSummary from "../components/features/user/UserStatsSummary.tsx";
import UserCharts from "../components/features/user/UserCharts.tsx";
import UserDiagrams from "../components/features/user/UserDiagrams.tsx";
import MapsetList from "../components/features/mapset/MapsetList.tsx";

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
        <Layout className="flex justify-center" title={user ? user.username : "Loading..."}>
            {userLoaded == LoadingState.Succeeded ?
                <div className="w-full max-w-[1152px] grid gap-4 px-2 md:px-0">
                    <User
                        user={user}
                        externalLinkOnClick={() => extLinkOnClick(user.id)}
                    >
                        <MapStatsSummary user={user}/>
                        <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                    </User>
                    <div className="flex flex-col md:flex-row gap-4">
                        <UserCharts
                            className="p-2 md:p-4 w-full md:w-1/2"
                            data={mapUserStatsToArray(user.user_stats)}
                        />
                        <UserDiagrams
                            className="p-2 md:p-4 w-full md:w-1/2"
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