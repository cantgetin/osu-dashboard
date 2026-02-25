import HeaderLink from "./HeaderLink.tsx";
import GlobalSearch from "./GlobalSearch .tsx";
import UserLogin from "../logic/LoginNew.tsx";
import {useAppDispatch, useAppSelector} from "../../store/hooks";
import {clearUser, selectAuthUser, setUser} from "../../store/authSlice";
import {useEffect} from "react";
import {useNavigate} from "react-router-dom";

const Header = () => {
    const dispatch = useAppDispatch();
    const navigate = useNavigate();
    const user = useAppSelector(selectAuthUser);

    useEffect(() => {
        try {
            const stored = localStorage.getItem("osu_dashboard_user");
            if (stored) {
                const parsed = JSON.parse(stored);
                if (parsed && parsed.id && parsed.username) {
                    dispatch(setUser(parsed));
                }
            }
        } catch {
            // ignore malformed storage
        }
    }, [dispatch]);

    const handleLogin = () => {
        navigate("/authorize");
    };

    const handleLogout = () => {
        localStorage.removeItem("osu_dashboard_user");
        dispatch(clearUser());
    };

    return (
        <div
            className="z-20 bg-zinc-900 fixed w-full bg-opacity-85 text-[16px] sm:text-[16px] backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-full sm:px-0 h-14 flex items-center justify-between overflow-x-auto">
                <div className="flex px-10 items-center justify-between w-full">
                    {/* Navigation Links - Left Side */}
                    <div className="flex items-center flex-nowrap">
                        <HeaderLink to="/">Home</HeaderLink>
                        <HeaderLink to="/users">Users</HeaderLink>
                        <HeaderLink to="/beatmapsets">Beatmaps</HeaderLink>
                        <HeaderLink to="/logs">Logs</HeaderLink>
                    </div>
                    {/* Search + Auth */}
                    <div className="flex items-center gap-4">
                        <GlobalSearch/>
                        <UserLogin
                            user={user ? {
                                id: String(user.id),
                                name: user.username,
                                email: "",
                                avatar: user.avatar_url,
                            } : null}
                            onLogin={handleLogin}
                            onLogout={handleLogout}
                            showNickname={true}
                            nicknamePosition="left"
                        />
                    </div>
                </div>
            </div>
        </div>
    );
};

export default Header;