import {useParams} from "react-router-dom";
import Header from "../components/Header.tsx";
import Content from "../components/Content.tsx";
import {useEffect, useState} from "react";
import LoadingSpinner from "../components/LoadingSpinner.tsx";

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

    return (
        <>
            <Header/>
            <Content>
                {beatmapset ?
                    <div className="-mt-6 w-full pt-0 flex flex-col gap-2 text-4xl">
                        <img src={beatmapset.covers['cover@2x']} alt="map bg"/>
                        <h1>{beatmapset.title}</h1>
                        <h1>{beatmapset.artist}</h1>
                        <p>{beatmapset.status}</p>
                        <p>{beatmapset.creator}</p>
                        <p>{beatmapset.last_updated}</p>

                    </div>
                    : <LoadingSpinner/>}
            </Content>
        </>
    );
};

export default Beatmapset;