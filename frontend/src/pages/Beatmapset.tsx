import Layout from "../components/ui/Layout.tsx";
import MapsetHero from "../components/features/mapset/MapsetHero.tsx";
import StatsComparison from "../components/features/mapset/StatsComparison.tsx";
import MapsetCharts from "../components/features/mapset/MapsetCharts.tsx";
import {mapMapsetStatsToArray} from "../utils/utils.ts";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import {LoadingState} from "../interfaces/LoadingState.ts";
import {fetchMapset, selectMapset, selectMapsetLoading} from "../store/mapsetSlice.ts";
import {useEffect} from "react";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {useParams} from "react-router-dom";

const Beatmapset = () => {
    const {mapId} = useParams();
    const dispatch = useAppDispatch();
    const beatmapset = useAppSelector(selectMapset);
    const beatmapsetLoaded = useAppSelector(selectMapsetLoading);

    useEffect(() => {
        dispatch(fetchMapset(mapId as string));
    }, [dispatch, mapId]);

    if (beatmapsetLoaded !== LoadingState.Succeeded || !beatmapset) {
        return <LoadingSpinner/>;
    }

    const mapsetStats = mapMapsetStatsToArray(beatmapset.mapset_stats);
    const lastStats = mapsetStats[mapsetStats.length - 1];
    const penultimateStats = mapsetStats[mapsetStats.length - 2] || lastStats;

    return (
        <Layout title={beatmapset.title}>
            <div className="pt-15 flex flex-col flex-wrap gap-2 w-full">
                <div className="relative h-[550px] rounded-md overflow-hidden">
                    <MapsetHero
                        beatmapset={beatmapset}
                        background={beatmapset.covers['cover@2x']}
                    >
                        <StatsComparison
                            lastStats={lastStats}
                            penultimateStats={penultimateStats}
                        />
                    </MapsetHero>
                </div>

                {/* Charts Section */}
                <div className="justify-center flex w-full mt-4">
                    <MapsetCharts
                        data={mapsetStats}
                        className="p-4 rounded-md"
                    />
                </div>
            </div>
        </Layout>
    );
};

export default Beatmapset;