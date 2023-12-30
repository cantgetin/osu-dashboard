const Header = () => {
    return (
        <div className="z-10 bg-zinc-900 h-10 w-full flex">
            <a href={"/"} className="cursor-pointer hover:bg-zinc-800 p-2">Main page</a>
            <a href={"/users"} className="cursor-pointer hover:bg-zinc-800 p-2">Users</a>
            <a href={"/beatmapsets"} className="cursor-pointer hover:bg-zinc-800 p-2">Beatmaps</a>
        </div>
    );
};

export default Header;