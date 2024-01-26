import { useParams } from "react-router-dom";
import { useEffect, useState } from "react";
import User from "../components/User.tsx";
import UserStatsSummary from "../components/UserStatsSummary.tsx";
import Header from "../components/Header.tsx";
import UserChartsSummary from "../components/UserChartsSummary.tsx";
import MapsetSummaryList from "../components/MapsetSummaryList.tsx";
import MapStatsSummary from "../components/MapStatsSummary.tsx";
import { mapUserStatsToArray } from "../utils/utils.ts";

const UserPage = () => {
    const { userId } = useParams();
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
            <Header />
            <div className="flex justify-center items-center">
                <div className="p-10 flex flex-col gap-2 2xl:w-1/2">
                    {userCard && (
                        <>
                            <User user={userCard.User} nameOnClick={() => window.open(`https://osu.ppy.sh/users/${userCard.User.id}`)}>
                                <MapStatsSummary data={userCard} />
                                <UserStatsSummary data={mapUserStatsToArray(userCard.User.user_stats)} />
                            </User>
                            <UserChartsSummary data={mapUserStatsToArray(userCard.User.user_stats)} />
                            <MapsetSummaryList Mapsets={userCard.Mapsets} />
                        </>
                    )}
                </div>
            </div>
        </>
    );
};

export default UserPage;