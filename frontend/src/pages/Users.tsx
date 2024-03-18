import {useEffect, useState} from "react";
import User from "../components/business/User.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import UserStatsSummary from "../components/business/UserStatsSummary.tsx";
import UserCharts from "../components/business/UserCharts.tsx";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import List from "../components/business/List.tsx";
import Layout from "../components/ui/Layout.tsx";

const Users = () => {
    const [users, setUsers] = useState<User[]>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/user/list`);
            const userData = await response.json();

            setUsers(JSON.parse(JSON.stringify(userData)) as User[])
        })()
    }, [])

    const userNameOnClick = (userId: number) => window.open(`/user/${userId}`, "_self")

    const userExtLinkOnClick = (userId: number) => window.open(`https://osu.ppy.sh/users/${userId}`)

    return (
        <Layout className="flex md:justify-center sm:justify-start">
            {users != null ?
                <List className="w-[1152px] grid 2xl:grid-cols-1 l:grid-cols-1 gap-4" items={users}
                      renderItem={(user: User) => (
                          <User
                              user={user}
                              key={user.id}
                              nameOnClick={() => userNameOnClick(user.id)}
                              externalLinkOnClick={() => userExtLinkOnClick(user.id)}
                          >
                              <UserCharts
                                  className="w-[400px] min-w-[400px] max-w-[400px]"
                                  data={mapUserStatsToArray(user.user_stats)} asSlideshow={true}
                              />
                              <UserStatsSummary data={mapUserStatsToArray(user.user_stats)}/>
                          </User>
                      )}
                />
                : <LoadingSpinner/>}
        </Layout>
    );
};

export default Users;