import { Link, useLocation } from 'react-router-dom';
import SearchBar from "./SearchBar.tsx";

const Header = () => {
    const location = useLocation();

    return (
        <div className="z-20 bg-zinc-900 fixed w-full bg-opacity-85 text-md backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-[1152px] h-14 flex items-center">
                <Link
                    to="/"
                    className={`cursor-pointer p-4 bg-opacity-65 ${location.pathname === '/' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                >
                    Home
                </Link>
                <Link
                    to="/users"
                    className={`cursor-pointer p-4 bg-opacity-65 ${location.pathname === '/users' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                >
                    Users
                </Link>
                <Link
                    to="/beatmapsets"
                    className={`cursor-pointer p-4 bg-opacity-65 ${location.pathname === '/beatmapsets' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                >
                    Beatmaps
                </Link>
                <SearchBar className="my-4 rounded-md w-64 px-5 h-8 ml-auto" placeholder="Search..."></SearchBar>
            </div>
        </div>
    );
};

export default Header;