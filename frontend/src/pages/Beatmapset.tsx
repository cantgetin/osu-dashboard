import {useParams} from "react-router-dom";
import {useEffect, useState} from "react";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import aveta from "aveta";
import MapsetCharts from "../components/MapsetCharts.tsx";
import {convertDateFormat, mapMapsetStatsToArray} from "../utils/utils.ts";
import Layout from "../components/ui/Layout.tsx";
import {FaExternalLinkAlt} from "react-icons/fa";
import Button from "../components/ui/Button.tsx";

const Beatmapset = () => {
    const {mapId} = useParams();
    const [beatmapset, setBeatmapset] = useState<Mapset>();

    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/beatmapset/${mapId}`);
            const mapsetData = await response.json();

            setBeatmapset(JSON.parse(JSON.stringify(mapsetData)) as Mapset)
        })();
    }, [mapId]);

    const mapsetStats = beatmapset ? mapMapsetStatsToArray(beatmapset.mapset_stats) : [];

    const lastStats = mapsetStats[mapsetStats.length - 1]

    const penultimateStats = mapsetStats[mapsetStats.length - 2] != undefined
        ? mapsetStats[mapsetStats.length - 2]
        : mapsetStats[mapsetStats.length - 1]

    const externalLinkOnClick = () => window.open(`https://osu.ppy.sh/s/${beatmapset?.id}`, "_blank");

    return (
        <Layout>
            {beatmapset ?
                <div className="pt-15 flex flex-col flex-wrap gap-2 relative w-full">
                    <img src={beatmapset.covers['cover@2x']} alt="map bg" className="h-[550px] object-cover rounded-md"/>
                    <div className="p-5 absolute inset-0 flex justify-center h-[550px]">
                        <div className="w-1/2 justify-end flex flex-col gap-6">
                            {/*<div className="flex gap-2">*/}
                            {/*    {beatmapset.beatmaps.map((_, index) =>*/}
                            {/*    <div key={index} className="bg-zinc-900 h-6 w-6 rounded-lg text-center">{index}</div>*/}
                            {/*        )}*/}
                            {/*</div>*/}
                            <div>
                                <div className="flex gap-4 items-center">
                                    <h1 className="text-5xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{beatmapset.title}</h1>
                                    <Button
                                        onClick={() => externalLinkOnClick()}
                                        className="bg-zinc-800 rounded-md p-1 h-6"
                                        content={<FaExternalLinkAlt className="h-3"/>}
                                    />
                                </div>
                                <h1 className="text-4xl text-zinc-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">by {beatmapset.artist}</h1>
                            </div>
                            <div className="flex gap-1 flex-wrap max-w-[600px]">
                                {
                                    beatmapset.tags != '' &&
                                    beatmapset.tags.split(' ').map((tag, index) =>
                                        <div
                                            key={index}
                                            className="bg-zinc-800 px-2 py-1 rounded-lg text-sm cursor-pointer h-7">
                                            {tag}
                                        </div>
                                    )
                                }
                            </div>
                            <div className="text-xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                                last updated {convertDateFormat(beatmapset.last_updated)}
                            </div>
                            <div className="flex gap-2 text-4xl">
                                <p className="text-2xl text-red-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{beatmapset.status}</p>
                                <h1 className="text-2xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">mapped by</h1>
                                <a
                                    href={`/user/${beatmapset.user_id}`}
                                    className="text-2xl text-blue-300 hover:text-yellow-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]"
                                >
                                    {beatmapset.creator}
                                </a>
                            </div>
                        </div>
                        <div className="w-1/2 justify-end flex flex-col ml-auto whitespace-nowrap">
                            <div className="flex flex-col justify-center items-center">
                                <div
                                    className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-pink-300'>
                                    <h1>{aveta(lastStats.favourite_count)} Favorites</h1>
                                    <h1 className="text-xl">▲</h1>
                                    <h1>{aveta(lastStats.favourite_count - penultimateStats.favourite_count)}</h1>
                                </div>
                                <div
                                    className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-green-300'>
                                    <h1>{aveta(lastStats.play_count)} Plays</h1>
                                    <h1 className="text-xl">▲</h1>
                                    <h1>{aveta(lastStats.play_count - penultimateStats.play_count)}</h1>
                                </div>
                                <div
                                    className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-red-300'>
                                    <h1>{aveta(lastStats.comments_count)} Comments</h1>
                                    <h1 className="text-xl">▲</h1>
                                    <h1>{aveta(lastStats.comments_count - penultimateStats.comments_count)}</h1>
                                </div>
                            </div>
                        </div>
                    </div>
                    <div className="justify-center flex w-full">
                        <MapsetCharts
                            data={mapMapsetStatsToArray(beatmapset.mapset_stats)}
                            className="p-4 rounded-md"
                        />
                    </div>
                </div>
                : <LoadingSpinner/>}
        </Layout>
    );
};

export default Beatmapset;