import Mapset from "./Mapset.tsx";
import {useAppDispatch, useAppSelector} from "../../store/hooks.ts";
import Button from "../ui/Button.tsx";
import List from "./List.tsx";
import {fetchMapsets, selectMapsets} from "../../store/mapsetsSlice.ts";

interface MapsetListProps {
    userId?: string;
}

const MapsetList = (props: MapsetListProps) => {
    const dispatch = useAppDispatch();

    const mapsets = useAppSelector<Mapset[]>(selectMapsets);

    // todo add mapset count prop to list mapset backend handler
    const pages = 10
    const currentPage = 1

    const handlePageChange = (page: number) => {
        dispatch(fetchMapsets({userId: Number(props.userId), page: page}))
    };

    // @ts-ignore
    const buttons = pages === 1 ? [] : Array.apply(null, Array(pages)).map(function (_, i) {
        return i + 1;
    });

    return (
        <div className="flex flex-col gap-2">
            <List className="flex flex-col gap-2"
                  items={mapsets}
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
                                  onClick={() => handlePageChange(num)}
                                  className={"rounded-md w-12 " + (num === currentPage ? "bg-white text-black" : "bg-zinc-800")}
                                  content={num.toString()}
                          />
                      }
                />
            </div>
        </div>
    );
};

export default MapsetList;