const PlaysSummary = () => {
    return (
        <>
            <h1 className="text-2xl">123 Maps fetched</h1>
            <div>
                <div className="text-xl text-yellow-200">2100 Plays now</div>
                <div className="text-sm text-orange-200">1450 plays last time</div>
            </div>
            <div className="flex flex-col mt-auto ml-auto">
                <div className="flex gap-2 justify-center items-center ml-auto px-2">
                    <h1 className="text-xs text-green-300">▲</h1>
                    <h1 className="text-2xl text-green-300">432</h1>
                </div>
                <div className="text-xs text-zinc-400">total plays for
                    last 24 hours
                </div>
            </div>

        </>
    );
};

export default PlaysSummary;