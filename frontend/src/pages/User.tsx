import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";

const User = () => {
    const { userId } = useParams();
    const [user, setUser] = useState(null);


    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`http://localhost:8080/user_card/${userId}`);
                const userData = await response.json();

                setUser(userData);
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData();
    }, [userId]);

    return (
        <div>
            <p>User ID: {userId}</p>
            {user && (
                <div>
                    <p>JSON Data: {JSON.stringify(user, null, 2)}</p>
                </div>
            )}
        </div>
    );
};

export default User;