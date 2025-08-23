import { convertDateFormat } from "../../../utils/time.ts";

interface MapsetMetadataProps {
    status: string;
    creator: string;
    userId: number;
    lastUpdated: string;
}

const MapsetMeta = ({ status, creator, userId, lastUpdated }: MapsetMetadataProps) => (
    <div className="text-lg flex flex-col gap-2 md:text-4xl">
        <p className="md:text-2xl text-red-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{status}</p>
        <div className="flex gap-2">
            <h1 className="md:text-2xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">mapped by</h1>
            <a
                href={`/user/${userId}`}
                className="md:text-2xl text-blue-300 hover:text-yellow-200 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]"
            >
                {creator}
            </a>
        </div>
        <div className="md:text-xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
            last updated {convertDateFormat(lastUpdated)}
        </div>
    </div>
);

export default MapsetMeta;