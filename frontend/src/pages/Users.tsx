import {useEffect, useState} from "react";
import Header from "../components/Header.tsx";

interface Following {
    id: number
    username: string
    tracking_since: string
}

const Users = () => {

    const [follows, setFollows] = useState<Following[]>();

    useEffect(() => {
        const fetchUserData = async () => {
            try {
                const response = await fetch(`/api/following/list`);
                const userData = await response.json();

                setFollows(JSON.parse(JSON.stringify(userData)  ) as Following[])
            } catch (error) {
                console.error('Error fetching user data:', error);
            }
        };

        fetchUserData()
    }, [])

    return (
        <>
            <Header/>
            {/*{JSON.stringify(follows)}*/}
            <div className="p-4 flex flex-col gap-2">
            {follows && follows.length > 0 && follows.map(follow => (
                <a className="text-xl underline" href={`/user/${follow.id}`}>
                    {follow.username}
                </a>
            ))}
            </div>
        </>
    );
};

export default Users;