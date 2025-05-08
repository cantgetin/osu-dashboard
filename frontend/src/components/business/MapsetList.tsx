import Mapset from "./Mapset.tsx";
import {useAppDispatch, useAppSelector} from "../../store/hooks.ts";
import List from "../logic/List.tsx";
import {fetchMapsets, fetchMapsetsProps, MapsetsState, selectMapsetsState} from "../../store/mapsetsSlice.ts";
import {useEffect} from "react";
import {LoadingState} from "../../interfaces/LoadingState.ts";
import LoadingSpinner from "../ui/LoadingSpinner.tsx";
import Pagination from "./Pagination.tsx";
import MapsetSearch from "./MapsetSearch.tsx";

interface MapsetListProps extends fetchMapsetsProps {
    showMapper?: boolean;
    // Define any additional props here
}

const MapsetList = ({showMapper, ...mapsetProps}: MapsetListProps) => {
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
        <div className="flex flex-col gap-3 md:gap-5 bg-zinc-900 p-2 md:p-4 rounded-lg w-full">
            <MapsetSearch update={onSearch}/>
            {
                mapsetsState.loading == LoadingState.Succeeded ?
                    <>
                        <List className="flex flex-col gap-2 md:gap-4 rounded-lg w-full"
                              items={mapsetsState.mapsets!}
                              renderItem={(mapset: Mapset) =>
                                  <Mapset
                                      key={mapset.id}
                                      map={mapset}
                                      showMapper={showMapper}
                                  />
                              }
                        />
                        <Pagination
                            pages={mapsetsState.pages}
                            currentPage={mapsetsState.currentPage}
                            onPageChange={onPageChange}
                            className="flex gap-1 md:gap-2 justify-end text-sm md:text-md"
                        />
                    </>
                    : <LoadingSpinner/>
            }
        </div>
    );
};

export default MapsetList;