import SearchBar from "../components/ui/SearchBar.tsx";
import {useEffect, useState} from "react";
import {convertDataToDayMonth} from "../utils/utils.ts";
import Layout from "../components/ui/Layout.tsx";

const AddUser = () => {
    const [users, setUsers] = useState<User[]>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/users/list`);
            const userData = await response.json();

            setUsers(JSON.parse(JSON.stringify(userData)) as User[])
        })()
    }, [])

    return (
        <Layout>
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
        </Layout>
    );
};

export default AddUser;