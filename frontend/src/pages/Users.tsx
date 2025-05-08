import {useEffect} from "react";
import User from "../components/business/User.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import UserStatsSummary from "../components/business/UserStatsSummary.tsx";
import UserCharts from "../components/business/UserCharts.tsx";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import List from "../components/logic/List.tsx";
import Layout from "../components/ui/Layout.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";
import {fetchUsers, selectUsers, selectUsersLoadingState} from "../store/usersSlice.ts";

const Users = () => {
    const dispatch = useAppDispatch();
    const users = useAppSelector<User[]>(selectUsers)
    const usersLoaded = useAppSelector<LoadingState>(selectUsersLoadingState)

    useEffect(() => {
        dispatch(fetchUsers())
    }, [dispatch])

    const userNameOnClick = (userId: number) => window.open(`/user/${userId}`, "_self")
    const userExtLinkOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    return (
        <Layout className="flex justify-center" title="Users">
            {usersLoaded == LoadingState.Succeeded ? (
                <List
                    className="w-full px-2 sm:px-0 sm:w-[1152px] grid grid-cols-1 gap-4"
                    items={users}
                    title={`Total users: ${users.length}`}
                    renderItem={(user: User) => (
                        <User
                            user={user}
                            key={user.id}
                            nameOnClick={() => userNameOnClick(user.id)}
                            externalLinkOnClick={() => userExtLinkOnClick(user.id)}
                        >
                            <UserCharts
                                className="w-full sm:w-[400px] sm:min-w-[400px] sm:max-w-[400px]"
                                data={mapUserStatsToArray(user.user_stats)}
                                asSlideshow={true}
                            />
                            <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                        </User>
                    )}
                />
            ) : (
                <LoadingSpinner/>
            )}
        </Layout>
    );
};

export default Users;