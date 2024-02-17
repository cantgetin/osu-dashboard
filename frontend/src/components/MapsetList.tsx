import Mapset from "./Mapset";
import {fetchUserCard, selectUserCardPage} from "../store/userCardSlice.ts";
import {useAppDispatch, useAppSelector} from "../store/hooks.ts";

interface MapsetSummaryProps {
    Mapsets: Mapset[];
    MapsetCount: number;
    userId: string;
}

const MapsetList = (props: MapsetSummaryProps) => {
    const numPages = Math.ceil(props.MapsetCount / 50);

    const dispatch = useAppDispatch();

    const currentPage = useAppSelector<number>(selectUserCardPage);

    const handlePageChange = (page: number) => {
        dispatch(fetchUserCard({userId: Number(props.userId), page: page}))
    };

    const renderPaginationButtons = () => {
        const buttons = [];
        for (let i = 1; i <= numPages; i++) {
            buttons.push(
                <button
                    key={i}
                    onClick={() => handlePageChange(i)}
                    className={i === currentPage ? "w-12 bg-blue-500 text-white" : "w-12"}
                >
                    {i}
                </button>
            );
        }
        return buttons;
    };

    return (
        <div className="flex flex-col gap-2">
            {numPages > 1 && (
                <div className="flex space-x-2">
                    <div>Page:</div>
                    {renderPaginationButtons()}
                </div>
            )}
            {props.Mapsets.map(mapset => (
                <Mapset key={mapset.id} map={mapset} />
            ))}
        </div>
    );
};

export default MapsetList;