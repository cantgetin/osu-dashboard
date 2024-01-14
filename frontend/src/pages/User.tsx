import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import User from "../components/User.tsx";
import PlaysSummary from "../components/PlaysSummary.tsx";
import MapsetSummary from "../components/MapsetSummary.tsx";
import Header from "../components/Header.tsx";
import {mapUserStatsToArray} from "../utils/utils.ts";
import ChartsSummary from "../components/ChartsSummary.tsx";

const UserPage = () => {
    const {userId} = useParams();

    const [userCard, setUser] = useState<UserCard>();
    const [userData, setUserData] = useState<UserStatsDataset[]>([]);

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user_card/${userId}`);
                const userData = await response.json();

                setUser(JSON.parse(JSON.stringify(userData)) as UserCard)
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData()
    }, [userId]);

    useEffect(() => {
        if (userCard) {
            setUserData(mapUserStatsToArray(userCard!.User.user_stats))
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
                                <div></div>
                                <PlaysSummary data={userData}/>
                            </User>
                            <ChartsSummary data={userData}/>
                            {userCard.Mapsets.map(mapset => <MapsetSummary key={mapset.id} map={mapset}/>)}
                        </>
                    )}
                </div>
            </div>
        </>
    );
};

export default UserPage;