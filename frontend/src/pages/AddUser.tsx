import SearchBar from "../components/ui/SearchBar.tsx";
import {useEffect} from "react";
import {convertDataToDayMonth} from "../utils/utils.ts";
import Layout from "../components/ui/Layout.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {fetchUsers, selectUsers, selectUsersLoadingState} from "../store/usersSlice.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";

const AddUser = () => {
    const dispatch = useAppDispatch();
    const users = useAppSelector<User[]>(selectUsers)
    const usersLoaded = useAppSelector<LoadingState>(selectUsersLoadingState)

    useEffect(() => {
        dispatch(fetchUsers())
    }, [dispatch])

    return (
        <Layout>
            {usersLoaded === LoadingState.Succeeded ?
                <>
                    <h1 className="text-4xl leading-tight">Add user</h1>
                    <SearchBar className="my-4 rounded-md w-64 px-5 h-8" placeholder="Username..."></SearchBar>
                    <div className="w-64">
                        {
                            <ul>
                                {
                                    users?.map(user =>
                                        <li className="flex justify-between items-center">
                                            <div className="text-xl">{user.username}</div>
                                            <div
                                                className="text-md">added {convertDataToDayMonth(user.tracking_since)}</div>
                                        </li>
                                    )
                                }
                            </ul>
                        }
                    </div>
                </>
                : <LoadingSpinner/>
            }
        </Layout>
    );
};

export default AddUser;