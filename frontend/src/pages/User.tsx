import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import User from "../components/User.tsx";
import PlaysSummary from "../components/PlaysSummary.tsx";
import Header from "../components/Header.tsx";
import UserChartsSummary from "../components/UserChartsSummary.tsx";
import MapsetSummaryList from "../components/MapsetSummaryList.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";

const UserPage = () => {
    const {userId} = useParams();

    const [userCard, setUserCard] = useState<UserCard>();
    const [userData, setUserData] = useState<UserStatsDataset[]>([]);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`/api/user_card/${userId}`);
                const userData = await response.json();

                setUserCard(JSON.parse(JSON.stringify(userData)) as UserCard)
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData()
    }, [userId]);

    useEffect(() => {
        if (userCard) {
            setUserData(mapUserStatsToArray(userCard!.User.user_stats));
        }
    }, [userCard]);

    return (
        <>
            <Header/>
            <div className="flex justify-center items-center">
                <div className="p-10 flex flex-col gap-2">
                    {userCard && userData.length > 0 && (
                        <>
                            <User user={userCard.User}>
                                <></>
                                <PlaysSummary data={userData}/>
                            </User>
                            <UserChartsSummary data={userData}/>
                            <MapsetSummaryList Mapsets={userCard.Mapsets}/>
                        </>
                    )}
                </div>
            </div>
        </>
    );
};

export default UserPage;