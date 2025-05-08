import {useEffect} from 'react';
import {useAppDispatch, useAppSelector} from "../../store/hooks.ts";
import {
    fetchSystemStats,
    selectSystemStatsLoading,
    selectSystemStatsState,
    SystemStatsState
} from "../../store/systemStats.ts";
import {LoadingState} from "../../interfaces/LoadingState.ts";

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
                            className="grid grid-cols-2 md:flex justify-between bg-zinc-900
                            py-4 md:py-8 px-4 md:px-16 rounded-lg gap-4 md:gap-0">
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{systemStats.stats!.users}</span>
                                <span className="text-gray-400 text-sm md:text-base">Users</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{systemStats.stats!.mapsets}</span>
                                <span className="text-gray-400 text-sm md:text-base">Mapsets</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{systemStats.stats!.beatmaps}</span>
                                <span className="text-gray-400 text-sm md:text-base">Beatmaps</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-3xl md:text-5xl font-bold">{systemStats.stats!.tracks}</span>
                                <span className="text-gray-400 text-sm md:text-base">Tracks</span>
                            </div>
                        </div>
                    </div>
                </section>
            )}
        </>
    );
};

export default SystemStats;