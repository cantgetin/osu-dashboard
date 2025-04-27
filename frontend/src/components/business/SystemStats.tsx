import { useEffect } from 'react';
import { useAppDispatch, useAppSelector } from "../../store/hooks.ts";
import {
    fetchSystemStats,
    selectSystemStatsLoading,
    selectSystemStatsState,
    SystemStatsState
} from "../../store/systemStats.ts";
import { LoadingState } from "../../interfaces/LoadingState.ts";

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
                <>
                    <section id="stats">
                        <h1 className="text-5xl leading-tight">
                            Stats
                        </h1>
                        <div className="flex justify-between bg-zinc-900 p-8 rounded-lg">
                            <div className="flex flex-col items-center">
                                <span className="text-5xl font-bold">{systemStats.stats!.users}</span>
                                <span className="text-gray-400">Users</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-5xl font-bold">{systemStats.stats!.mapsets}</span>
                                <span className="text-gray-400">Mapsets</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-5xl font-bold">{systemStats.stats!.beatmaps}</span>
                                <span className="text-gray-400">Beatmaps</span>
                            </div>
                            <div className="flex flex-col items-center">
                                <span className="text-5xl font-bold">{systemStats.stats!.tracks}</span>
                                <span className="text-gray-400">Tracks</span>
                            </div>
                        </div>
                    </section>
                </>
            )}
        </>
    );
};

export default SystemStats;