import Mapset from "./Mapset.tsx";
import {useEffect, useState} from "react";

interface MapsetSummaryProps {
    Mapsets: Mapset[]
}

const MapsetList = (props: MapsetSummaryProps) => {

    const [sorted, setSorted] = useState<boolean>(false);

    useEffect(() => {
        if (props.Mapsets.length === 0) return;
        props.Mapsets.sort((a, b) => {
            const getKey = (mapset: Mapset) : string => Object.keys(mapset.mapset_stats).pop()!;
            return (getKey(b) ? b.mapset_stats[getKey(b)].play_count : 0) -
                (getKey(a) ? a.mapset_stats[getKey(a)].play_count : 0);
        });
        setSorted(true)
    }, [props]);

    return (
        <div className="flex flex-col gap-2">
            {sorted ? props.Mapsets.map(mapset =>
                <Mapset key={mapset.id} map={mapset}/>
            ) : null}
        </div>
    );
};

export default MapsetList;