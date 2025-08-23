import {useEffect} from 'react';
import {LoadingState} from "../../../interfaces/LoadingState.ts";
import List from "../../logic/List.tsx";
import Pagination from "../../ui/Pagination.tsx";
import LoadingSpinner from "../../ui/LoadingSpinner.tsx";
import {useAppDispatch, useAppSelector} from "../../../store/hooks.ts";
import UserSearch from "./UserSearch.tsx";
import User from "./User.tsx";
import UserCharts from "./UserCharts.tsx";
import {mapUserStatsToArray} from "../../../utils/utils.ts";
import UserStatsSummary from "./UserStatsSummary.tsx";
import {
    fetchUsers,
    fetchUsersProps,
    selectUsers,
    UsersState
} from "../../../store/usersSlice.ts";
import Container from "../../ui/Container.tsx";


const UserList = ({...props}: fetchUsersProps) => {
    const dispatch = useAppDispatch();

    const usersState = useAppSelector<UsersState>(selectUsers)

    useEffect(() => {
        dispatch(fetchUsers(props))
    }, [dispatch])

    const userNameOnClick = (userId: number) => window.open(`/user/${userId}`, "_self")
    const userExtLinkOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    const onPageChange = (page: number) => {
        dispatch(fetchUsers({...props, page: page} as fetchUsersProps))
    }

    const onSearch = (props: fetchUsersProps) => {
        dispatch(fetchUsers({
            ...props,
            search: props.search,
            sort: props.sort,
            direction: props.direction,
        } as fetchUsersProps))
    }

    return (
        <Container>
            <UserSearch update={onSearch}/>
            {
                usersState.loading == LoadingState.Succeeded ?
                    <>
                        <List
                            className="w-full px-2 sm:px-0 grid grid-cols-1 gap-4"
                            items={usersState.users!}
                            renderItem={(user: User) => (
                                <User
                                    user={user}
                                    key={user.id}
                                    nameOnClick={() => userNameOnClick(user.id)}
                                    externalLinkOnClick={() => userExtLinkOnClick(user.id)}
                                    className="bg-zinc-800 bg-opacity-30"
                                >
                                    <UserCharts
                                        className="w-full sm:w-[400px] sm:min-w-[400px] sm:max-w-[400px]"
                                        data={mapUserStatsToArray(user.user_stats)}
                                        asSlideshow={true}/>
                                    <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                                </User>
                            )}/>
                        <Pagination
                            pages={usersState.pages}
                            currentPage={usersState.currentPage}
                            onPageChange={onPageChange}
                            className="flex gap-1 md:gap-2 justify-end text-sm md:text-md"
                        />
                    </>
                    : <LoadingSpinner/>
            }
        </Container>
    )
}

export default UserList;