import Mapset from "./Mapset";
import {fetchUserCard, selectUserCardPage} from "../store/userCardSlice.ts";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";
import Button from "./Button.tsx";
import List from "./List.tsx";

interface MapsetSummaryProps {
    Mapsets: Mapset[];
    MapsetCount: number;
    userId: string;
}

const MapsetList = (props: MapsetSummaryProps) => {
    const dispatch = useAppDispatch();

    const pages = Math.ceil(props.MapsetCount / 50);
    const currentPage = useAppSelector<number>(selectUserCardPage);

    const handlePageChange = (page: number) => {
        dispatch(fetchUserCard({userId: Number(props.userId), page: page}))
    };

    const buttons = pages === 1 ? [] : Array.apply(null, Array(pages)).map(function (_, i) {
        return i + 1;
    });

    return (
        <div className="flex flex-col gap-2">
            <div className="flex gap-2">
                <List className="flex gap-2"
                      title="Page:"
                      items={buttons}
                      renderItem={(num) =>
                          <Button keyNumber={num}
                                  key={num}
                                  onClick={() => handlePageChange(num)}
                                  className={"rounded-md w-12 " + (num === currentPage ? "bg-blue-500" : "bg-zinc-800")}
                                  content={num.toString()}
                          />
                      }
                />
            </div>
            <List className="flex flex-col gap-2"
                  items={props.Mapsets}
                  renderItem={(mapset: Mapset) =>
                      <Mapset
                          key={mapset.id}
                          map={mapset}
                      />
                  }
            />
        </div>
    );
};

export default MapsetList;