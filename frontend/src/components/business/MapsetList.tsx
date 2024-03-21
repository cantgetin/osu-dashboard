import Mapset from "./Mapset.tsx";
import {useAppDispatch, useAppSelector} from "../../store/hooks.ts";
import List from "../logic/List.tsx";
import {fetchMapsets, fetchMapsetsProps, MapsetsState, selectMapsetsState} from "../../store/mapsetsSlice.ts";
import {useEffect} from "react";
import {LoadingState} from "../../interfaces/LoadingState.ts";
import LoadingSpinner from "../ui/LoadingSpinner.tsx";
import Pagination from "./Pagination.tsx";
import MapsetSearch from "./MapsetSearch.tsx";

const MapsetList = (mapsetProps: fetchMapsetsProps) => {
    const dispatch = useAppDispatch();
    const mapsetsState = useAppSelector<MapsetsState>(selectMapsetsState);

    useEffect(() => {
        dispatch(fetchMapsets(mapsetProps))
    }, [dispatch]);

    const onPageChange = (page: number) => {
        dispatch(fetchMapsets({...mapsetProps, page: page} as fetchMapsetsProps))
    }

    const onSearch = (props: fetchMapsetsProps) => {
        dispatch(fetchMapsets({
            ...mapsetProps,
            search: props.search,
            status: props.status,
            sort: props.sort,
            direction: props.direction,
        } as fetchMapsetsProps))
    }

    return (
        <>
            <div className="flex flex-col gap-2">
                <MapsetSearch update={onSearch}/>
                {
                    mapsetsState.loading == LoadingState.Succeeded ?
                        <>
                            <Pagination
                                pages={mapsetsState.pages}
                                currentPage={mapsetsState.currentPage}
                                onPageChange={onPageChange}
                                className="flex gap-2 justify-end text-md"
                            />
                            <List className="flex flex-col gap-2"
                                  items={mapsetsState.mapsets!}
                                  renderItem={(mapset: Mapset) =>
                                      <Mapset
                                          key={mapset.id}
                                          map={mapset}
                                      />
                                  }
                            />
                        </>
                        : <LoadingSpinner/>
                }
            </div>
        </>
    );
};

export default MapsetList;