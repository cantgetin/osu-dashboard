import {useEffect, useState} from "react";
import Header from "../components/Header.tsx";
import User from "../components/User.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import UserCharts from "../components/UserCharts.tsx";
import LoadingSpinner from "../components/LoadingSpinner.tsx";
import Content from "../components/Content.tsx";
import List from "../components/List.tsx";

const Users = () => {
    const [users, setUsers] = useState<User[]>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/users/list`);
            const userData = await response.json();

            setUsers(JSON.parse(JSON.stringify(userData)) as User[])
        })()
    }, [])

    const userNameOnClick = (userId: number) => window.open(`/user/${userId}`, "_self")

    return (
        <>
            <Header/>
            <Content>
                {users != null ?
                    <List className="min-w-[1000px] p-10 grid 2xl:grid-cols-2 l:grid-cols-1 gap-4" items={users}
                          renderItem={(user: User) => (
                              <User user={user} key={user.id} nameOnClick={() => userNameOnClick(user.id)}>
                                  <UserCharts
                                      className="w-[400px] min-w-[400px] max-w-[400px]"
                                      data={mapUserStatsToArray(user.user_stats)} asSlideshow={true}
                                  />
                                  <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                              </User>
                          )}
                    />
                    : <LoadingSpinner/>}
            </Content>
        </>
    );
};

export default Users;