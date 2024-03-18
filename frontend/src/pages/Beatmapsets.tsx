import {useEffect, useState} from "react";
import List from "../components/business/List.tsx";
import Mapset from "../components/business/Mapset.tsx";
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import Layout from "../components/ui/Layout.tsx";
import MapsetSearch from "../components/business/MapsetSearch.tsx";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import {fetchMapsets, selectMapsets, selectMapsetsLoading} from "../store/mapsetsSlice.ts";
import {LoadingState} from "../interfaces/LoadingState.ts";

const Beatmapsets = () => {

    const dispatch = useAppDispatch();

    const mapsets = useAppSelector<Mapset[]>(selectMapsets);
    const loaded = useAppSelector<LoadingState>(selectMapsetsLoading)

    useEffect(() => {
        dispatch(fetchMapsets({}))
    }, [dispatch])

    return (
        <Layout className="flex flex-col gap-2 min-w-[800px]">
            <MapsetSearch/>
            {loaded == LoadingState.Succeeded ?
                <List
                    items={mapsets}
                    renderItem={(m: Mapset) =>
                        <Mapset
                            key={m.id}
                            map={m}
                            showMapper={true}
                            className="min-w-[800px]"
                        />
                    }
                    className="grid xl:grid-cols-1 gap-2 py-5"
                />
                : <LoadingSpinner/>
            }
        </Layout>
    );
};

export default Beatmapsets;