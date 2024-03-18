import Mapset from "./Mapset.tsx";
import {useAppDispatch, useAppSelector} from "../../store/hooks.ts";
import Button from "../ui/Button.tsx";
import List from "./List.tsx";
import {fetchMapsets, MapsetsState, selectMapsetsState} from "../../store/mapsetsSlice.ts";
import {useEffect, useState} from "react";
import {LoadingState} from "../../interfaces/LoadingState.ts";
import LoadingSpinner from "../ui/LoadingSpinner.tsx";

interface MapsetListProps {
    userId?: number;
}

const MapsetList = (props: MapsetListProps) => {
    const dispatch = useAppDispatch();
    const mapsetsState = useAppSelector<MapsetsState>(selectMapsetsState);
    const [page, setPage] = useState<number>(1);

    useEffect(() => {
        dispatch(fetchMapsets({
            userId: props.userId,
            forUser: true,
            page: page,
            sort: "last_playcount",
            direction: "desc",
        }))
    }, [dispatch, page]);

    const buttons = mapsetsState.pages === 1 ? [] : Array.apply(null, Array(mapsetsState.pages)).map(function (_, i) {
        return i + 1;
    });

    return (
        <>
            {
                mapsetsState.loading == LoadingState.Succeeded ?
                    <div className="flex flex-col gap-2">
                        <List className="flex flex-col gap-2"
                              items={mapsetsState.mapsets!}
                              renderItem={(mapset: Mapset) =>
                                  <Mapset
                                      key={mapset.id}
                                      map={mapset}
                                  />
                              }
                        />
                        <div className="flex gap-2 justify-end">
                            <List className="flex gap-2"
                                  title="Page:"
                                  items={buttons}
                                  renderItem={(num) =>
                                      <Button keyNumber={num}
                                              key={num}
                                              onClick={(key: number) => {
                                                  setPage(key)
                                              }}
                                              className={"rounded-md w-12 " + (num === page ? "bg-white text-black" : "bg-zinc-800")}
                                              content={num.toString()}
                                      />
                                  }
                            />
                        </div>
                    </div>
                    : <LoadingSpinner/>

            }
        </>
    );
};

export default MapsetList;