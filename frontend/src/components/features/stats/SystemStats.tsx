import {useEffect} from 'react';
import {useAppDispatch, useAppSelector} from "../../../store/hooks.ts";
import {LoadingState} from "../../../interfaces/LoadingState.ts";
import {
    fetchSystemStats,
    selectSystemStatsLoading,
    selectSystemStatsState,
    SystemStatsState
} from "../../../store/systemStats.ts";
import aveta from "aveta";

const SystemStats = () => {
    const dispatch = useAppDispatch();
    const systemStats = useAppSelector<SystemStatsState>(selectSystemStatsState);
    const statsLoaded = useAppSelector<LoadingState>(selectSystemStatsLoading);

    useEffect(() => {
        dispatch(fetchSystemStats());
    }, [dispatch]);

    return (
        <>
            {statsLoaded === LoadingState.Succeeded && (
                <section id="stats">
                    <div className="flex flex-col gap-5 md:gap-10">
                        <h1 className="text-3xl md:text-5xl leading-tight">
                            Stats
                        </h1>
                        <div
                            className="rounded-lg grid grid-cols-3 gap-y-7 grid-rows-2 justify-center items-center p-5 gap-5 bg-zinc-900"
                        >
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.users)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Users</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.mapsets)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Mapsets</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.beatmaps)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Beatmaps</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.plays)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Plays</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.favourites)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Favourites</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{aveta(systemStats.stats!.comments)}</span>
                                <span className="text-gray-400 text-sm md:text-base">Comments</span>
                            </div>
                        </div>
                    </div>
                </section>
            )}
        </>
    );
};

export default SystemStats;