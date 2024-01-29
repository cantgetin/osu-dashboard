import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import User from "../components/User.tsx";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import Header from "../components/Header.tsx";
import UserCharts from "../components/UserCharts.tsx";
import MapsetList from "../components/MapsetList.tsx";
import MapStatsSummary from "../components/MapStatsSummary.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import LoadingSpinner from "../components/LoadingSpinner.tsx";
import Content from "../components/Content.tsx";

const UserPage = () => {
    const {userId} = useParams();
    const [userCard, setUserCard] = useState<UserCard>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/user_card/${userId}`);
            const userData = await response.json();

            setUserCard(JSON.parse(JSON.stringify(userData)) as UserCard)
        })();
    }, [userId]);

    return (
        <>
            <Header/>
            <Content className="flex justify-center items-center">
                {userCard ?
                    <div className="p-10 pt-0 flex flex-col gap-2 2xl:w-1/2">
                        <>
                            <User user={userCard.User}
                                  nameOnClick={() => window.open(`https://osu.ppy.sh/users/${userCard.User.id}`)}>
                                <MapStatsSummary data={userCard}/>
                                <UserStatsSummary data={mapUserStatsToArray(userCard.User.user_stats)}/>
                            </User>
                            <UserCharts data={mapUserStatsToArray(userCard.User.user_stats)}/>
                            <MapsetList Mapsets={userCard.Mapsets}/>
                        </>

                    </div>
                    :
                    <LoadingSpinner/>}
            </Content>
        </>
    );
};

export default UserPage;