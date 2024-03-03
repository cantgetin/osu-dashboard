import {useParams} from "react-router-dom";
import Header from "../components/Header.tsx";
import Content from "../components/Content.tsx";
import {useEffect, useState} from "react";
import LoadingSpinner from "../components/LoadingSpinner.tsx";
import aveta from "aveta";
import MapsetCharts from "../components/MapsetCharts.tsx";
import {convertDateFormat, mapMapsetStatsToArray} from "../utils/utils.ts";

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

    const penultimateStats = mapsetStats[mapsetStats.length - 2] != undefined
        ? mapsetStats[mapsetStats.length - 2]
        : mapsetStats[mapsetStats.length - 1]


    return (
        <>
            <Header/>
            <Content>
                {beatmapset ?
                    <div className="-mt-6 pt-0 flex flex-col gap-2">
                        <img src={beatmapset.covers['cover@2x']} alt="map bg" className="2xl:mx-72 h-[550px] object-cover"/>
                        <div className="justify-center 2xl:px-72 absolute flex w-full">
                            <div className="p-4 justify-end flex flex-col gap-6 h-[550px]">
                                {/*<div className="flex gap-2">*/}
                                {/*    {beatmapset.beatmaps.map((_, index) =>*/}
                                {/*    <div key={index} className="bg-zinc-900 h-6 w-6 rounded-lg text-center">{index}</div>*/}
                                {/*        )}*/}
                                {/*</div>*/}
                                <div>
                                    <h1 className="text-6xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{beatmapset.title}</h1>
                                    <h1 className="text-5xl text-zinc-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">by {beatmapset.artist}</h1>
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
                                    <p className="text-4xl text-red-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{beatmapset.status}</p>
                                    <h1 className="drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">mapped by</h1>
                                    <a
                                        href={`/user/${beatmapset.user_id}`}
                                        className="text-blue-300 hover:text-yellow-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]"
                                    >
                                        {beatmapset.creator}
                                    </a>
                                </div>
                            </div>
                            <div className="p-4 justify-end flex flex-col ml-auto whitespace-nowrap">
                                <div className=" px-4 flex flex-col justify-center items-center">
                                    <div
                                        className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-pink-300'>
                                        <h1>{aveta(mapsetStats[mapsetStats.length - 1].favourite_count)} Favorites</h1>
                                        <h1 className="text-xl">▲</h1>
                                        <h1>{aveta(penultimateStats.favourite_count)}</h1>
                                    </div>
                                    <div
                                        className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-green-300'>
                                        <h1>{aveta(mapsetStats[mapsetStats.length - 1].play_count)} Plays</h1>
                                        <h1 className="text-xl">▲</h1>
                                        <h1>{aveta(penultimateStats.play_count)}</h1>
                                    </div>
                                    <div
                                        className='text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] text-red-300'>
                                        <h1>{aveta(mapsetStats[mapsetStats.length - 1].comments_count)} Comments</h1>
                                        <h1 className="text-xl">▲</h1>
                                        <h1>{aveta(penultimateStats.comments_count)}</h1>
                                    </div>
                                </div>
                            </div>
                        </div>
                        <div className="justify-center 2xl:px-72 flex w-full">
                            <MapsetCharts
                                data={mapMapsetStatsToArray(beatmapset.mapset_stats)}
                                className="p-4"
                            />
                        </div>
                    </div>
                    : <LoadingSpinner/>}
            </Content>
        </>
    );
};

export default Beatmapset;