const Header = () => {
    return (
        <div className="z-20 bg-zinc-900 fixed h-14 w-full flex bg-opacity-85 text-md backdrop-blur-sm">
            <a href={"/"} className="cursor-pointer hover:bg-zinc-800 p-4">Main page</a>
            <a href={"/users"} className="cursor-pointer hover:bg-zinc-800 p-4">Users</a>
            <a href={"/users/add"} className="cursor-pointer hover:bg-zinc-800 p-4">Add user</a>
            <a href={"/beatmapsets"} className="cursor-pointer hover:bg-zinc-800 p-4">Beatmaps</a>
        </div>
    );
};

export default Header;