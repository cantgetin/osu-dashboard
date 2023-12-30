import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import User from "../components/User.tsx";
import LineChart from "../components/LineChart.tsx";
import PlaysSummary from "../components/PlaysSummary.tsx";
import MapsetSummary from "../components/MapsetSummary.tsx";
import Header from "../components/Header.tsx";

const UserPage = () => {
    const {userId} = useParams();

    const [userCard, setUser] = useState<UserCard>();

    const UserData = [
        {
            year: "2023-12-25",
            plays: 1000,
        },
        {
            year: "2023-12-26",
            plays: 1100,
        },
        {
            year: "2023-12-27",
            plays: 1250,
        },
        {
            year: "2023-12-28",
            plays: 1350,
        },
        {
            year: "2023-12-29",
            plays: 2100,
        },
    ];

    const [chartData] = useState({
        labels: UserData.map((data) => data.year),
        datasets: [
            {
                data: UserData.map((data) => data.plays),
                backgroundColor: [
                    "#FEF07B",
                ],
                borderColor: "#101010",
                borderWidth: 2,
                pointStyle: 'circle',
                pointRadius: 6,
                pointHoverRadius: 10,
            },
        ],
    });


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

    return (
        <>
            <Header/>
            <div className="flex justify-center items-center">
                <div className="p-10 flex flex-col gap-2">
                    {userCard && (
                        <>
                            <User user={userCard.User}>
                                <LineChart chartData={chartData}/>
                                <PlaysSummary/>
                            </User>
                            <MapsetSummary map={userCard.Mapsets[0]}/>
                        </>
                    )}
                </div>
            </div>
        </>
    );
};

export default UserPage;