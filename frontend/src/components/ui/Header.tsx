import HeaderLink from "./HeaderLink.tsx";
import Search from "./Search.tsx";

const Header = () => {
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
                    {/* Search */}
                    <Search />
                </div>
            </div>
        </div>
    );
};

export default Header;