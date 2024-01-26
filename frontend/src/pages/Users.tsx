import {useEffect, useState} from "react";
import Header from "../components/Header.tsx";
import User from "../components/User.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import UserChartsSummary from "../components/UserChartsSummary.tsx";

const Users = () => {

    const [users, setUsers] = useState<User[]>();

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`/api/users/list`);
                const userData = await response.json();

                setUsers(JSON.parse(JSON.stringify(userData)) as User[])
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData()
    }, [])

    return (
        <>
            <Header/>
                <div className="p-10 grid xl:grid-cols-2 gap-4">
                    {users && users.length > 0 && users.map(user => (
                        <User user={user} key={user.id} nameOnClick={() => {window.open(`/user/${user.id}`)}}>
                            <UserChartsSummary data={mapUserStatsToArray(user.user_stats)} onlyPlaycount={true} />
                            <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                        </User>
                    ))}
                </div>
        </>
    );
};

export default Users;