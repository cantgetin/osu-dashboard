import Header from "../components/Header.tsx";
import Content from "../components/Content.tsx";
import {useEffect, useState} from "react";
import List from "../components/List.tsx";
import Mapset from "../components/Mapset.tsx";

const Beatmapsets = () => {
    const [mapsets, setMapsets] = useState<Mapset[]>([]);


    useEffect(() => {
        (async () => {
            const response = await fetch(`/api/beatmapset/list`);
            const userData = await response.json();

            setMapsets(JSON.parse(JSON.stringify(userData)) as Mapset[])
        })()
    }, [])

    return (
        <>
            <Header/>
            <Content className="p-5">
                <List
                    items={mapsets}
                    renderItem={(m: Mapset) =>
                        <Mapset
                            map={m}
                            showMapper={true}
                            className="min-w-[800px]"
                        />
                    }
                    className="grid xl:grid-cols-2 gap-2"/>
            </Content>
        </>
    );
};

export default Beatmapsets;