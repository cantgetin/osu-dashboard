import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import User from "../components/User.tsx";

const UserPage = () => {
    const {userId} = useParams();

    const [user, setUser] = useState<UserCard>();


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
        <div>
            {user && (
                <User user={user.User}>
                    <div></div>
                    <div></div>
                </User>
            )}
        </div>
    );
};

export default UserPage;