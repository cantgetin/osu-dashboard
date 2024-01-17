import MapsetSummary from "./MapsetSummary.tsx";
import {useEffect} from "react";

interface MapsetSummaryProps {
    Mapsets: Mapset[]
}

const MapsetSummaryList = (props: MapsetSummaryProps) => {

    useEffect(() => {
        if (props.Mapsets.length === 0) return;
        props.Mapsets.sort((a, b) => {
            const getKey = (mapset: Mapset) : string => Object.keys(mapset.mapset_stats).pop()!;
            return (getKey(b) ? b.mapset_stats[getKey(b)].play_count : 0) - (getKey(a) ? a.mapset_stats[getKey(a)].play_count : 0);
        });
    }, [props]);

    return (
        <>
            {props.Mapsets.map(mapset =>
                <MapsetSummary key={mapset.id} map={mapset}/>
            )}
        </>
    );
};

export default MapsetSummaryList;