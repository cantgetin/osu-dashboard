import {Link, useLocation} from 'react-router-dom';

const Header = () => {
    const location = useLocation();

    return (
        <div
            className="z-20 bg-zinc-900 fixed w-full bg-opacity-85 text-[16px] sm:text-[16px] backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-full px-2 sm:px-0 h-14 flex items-center justify-between overflow-x-auto">
                <div className="flex items-center flex-nowrap">
                    <Link
                        to="/"
                        className={`cursor-pointer p-3 sm:p-4 bg-opacity-65 whitespace-nowrap ${location.pathname === '/' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                    >
                        Home
                    </Link>
                    <Link
                        to="/users"
                        className={`cursor-pointer p-3 sm:p-4 bg-opacity-65 whitespace-nowrap ${location.pathname === '/users' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                    >
                        Users
                    </Link>
                    <Link
                        to="/beatmapsets"
                        className={`cursor-pointer p-3 sm:p-4 bg-opacity-65 whitespace-nowrap ${location.pathname === '/beatmapsets' ? 'bg-zinc-800' : 'hover:bg-zinc-800'}`}
                    >
                        Beatmaps
                    </Link>
                </div>
                {/* TODO: Uncomment when search is ready */}
                {/*<SearchBar className="my-2 rounded-md w-24 sm:w-40 md:w-64 px-3 sm:px-5 h-8"*/}
                {/*           placeholder="Search..."></SearchBar>*/}
            </div>
        </div>
    );
};

export default Header;