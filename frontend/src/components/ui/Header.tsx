import HeaderLink from "./HeaderLink.tsx";

const Header = () => {
    return (
        <div
            className="z-20 bg-zinc-900 fixed w-full bg-opacity-85 text-[16px] sm:text-[16px]
            backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-full px-2 sm:px-0 h-14 flex items-center justify-between overflow-x-auto">
                <div className="flex items-center flex-nowrap">
                    <HeaderLink to="/">Home</HeaderLink>
                    <HeaderLink to="/users">Users</HeaderLink>
                    <HeaderLink to="/beatmapsets">Beatmaps</HeaderLink>
                </div>
                {/* TODO: Uncomment when search is ready */}
                {/*<SearchBar className="my-2 rounded-md w-24 sm:w-40 md:w-64 px-3 sm:px-5 h-8"*/}
                {/*           placeholder="Search..."></SearchBar>*/}
            </div>
        </div>
    );
};

export default Header;