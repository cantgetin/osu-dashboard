const Footer = () => {
    return (
        <div
            className="z-20 bg-zinc-900 w-full bg-opacity-85 text-md backdrop-blur-sm flex justify-center items-center">
            <div className="max-w-[1152px] w-[1152px] h-14 flex items-center justify-center">
                <a href={"https://osu.ppy.sh/users/7192129"} className="mx-1.5 cursor-pointer hover:text-zinc-300">Creator</a>
                <a href={"https://github.com/cantgetin/playcount-monitor"} className="mx-10 cursor-pointer hover:text-zinc-300">Source code</a>
                <a href={"/"} className="cursor-pointer hover:text-zinc-300">Donate</a>
            </div>
        </div>
    );
};

export default Footer;