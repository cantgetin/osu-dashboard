import React, {useEffect, useState} from "react";
import Header from "../components/Header.tsx";
import User from "../components/User.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import UserCharts from "../components/UserCharts.tsx";
import LoadingSpinner from "../components/LoadingSpinner.tsx";

const Users = () => {
    const [users, setUsers] = useState<User[]>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/users/list`);
            const userData = await response.json();

            setUsers(JSON.parse(JSON.stringify(userData)) as User[])
        })()
    }, [])

    return (
        <>
            <Header/>
            <div className="p-10 grid xl:grid-cols-2 gap-4">
                {users && users.length > 0 ?
                    users.map(user => (
                        <User user={user} key={user.id} nameOnClick={() => {
                            window.open(`/user/${user.id}`)
                        }}>
                            <UserCharts data={mapUserStatsToArray(user.user_stats)} onlyPlaycount={true}/>
                            <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                        </User>))
                    :
                    <LoadingSpinner/>}
            </div>
        </>
    );
};

export default Users;