import HeaderLink from "./HeaderLink.tsx";
import {useState} from "react";
import UserLogin from "../logic/LoginNew.tsx";
import SearchBar from "@/components/ui/SearchBar.tsx";

const Header = () => {

    const [user, setUser] = useState<{
        name: string;
        email: string;
        avatar?: string;
        id: string,
    } | null>(null);

    const handleLogin = () => {
        // Simulate login
        setUser({
            name: "Gasha",
            email: "john.doe@example.com",
            avatar: "https://a.ppy.sh/7192129?1725050876.jpeg",
            id: "7192129",
        });
    };

    return (
        <div
            className="z-20 bg-zinc-900 fixed w-full bg-opacity-85 text-[16px] sm:text-[16px] backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-full sm:px-0 h-14 flex items-center justify-between overflow-x-auto">
                {/* This div maintains the px-10 padding */}
                <div className="flex px-10 items-center justify-between w-full"> {/* Added w-full here */}
                    {/* Navigation Links - Left Side */}
                    <div className="flex items-center flex-nowrap">
                        <HeaderLink to="/">Home</HeaderLink>
                        <HeaderLink to="/users">Users</HeaderLink>
                        <HeaderLink to="/beatmapsets">Beatmaps</HeaderLink>
                        <HeaderLink to="/logs">Logs</HeaderLink>
                    </div>
                    <div className="flex gap-4 items-center">
                        <SearchBar className="bg-zinc-800 bg-opacity-80 my-2 rounded-md w-24 sm:w-40 md:w-64 px-3 sm:px-5 h-8"
                                   placeholder="Search anything"></SearchBar>
                        <UserLogin
                            nicknamePosition={"left"}
                            user={user}
                            onLogin={handleLogin}
                            onLogout={() => setUser(null)}
                            onProfile={() => {
                                console.log("Profile clicked");
                            }}
                            onFeedback={() => {
                                console.log("Feedback clicked");
                            }}
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Header;