const Footer = () => {
    return (
        <div
            className="z-20 bg-zinc-900 w-full bg-opacity-85 text-md backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-[1152px] h-14 flex items-center">
                <a href={"https://osu.ppy.sh/users/7192129"} className="cursor-pointer hover:bg-zinc-800 p-4">Made by
                    Gasha</a>
                <a href={"https://github.com/cantgetin/playcount-monitor"}
                   className="cursor-pointer hover:bg-zinc-800 p-4">Github</a>
                <a href={"/"} className="cursor-pointer hover:bg-zinc-800 p-4">2024</a>
            </div>
        </div>
    );
};

export default Footer;