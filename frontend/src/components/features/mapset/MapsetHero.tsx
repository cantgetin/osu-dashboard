import MapsetHeader from "./MapsetHeader.tsx";
import Tags from "../common/Tags.tsx";
import MapsetMeta from "./MapsetMeta.tsx";


interface MapsetHeroProps {
    beatmapset: Mapset;
    children: React.ReactNode;
    background?: string; // Add background prop
}

const MapsetHero = ({beatmapset, children, background}: MapsetHeroProps) => (
    <div className="p-5 absolute inset-0 md:flex h-[550px]">
        {background && (
            <img
                src={background}
                alt="map bg"
                className="absolute inset-0 h-full w-full object-cover rounded-md opacity-60 -z-10"
            />
        )}
        <div className="md:w-1/2 justify-end flex flex-col gap-6">
            <MapsetHeader
                title={beatmapset.title}
                artist={beatmapset.artist}
                id={beatmapset.id}
            />
            <div className="flex gap-1 flex-wrap-reverse">
                <Tags tags={beatmapset.tags}/>
            </div>
            <MapsetMeta
                status={beatmapset.status}
                creator={beatmapset.creator}
                userId={beatmapset.user_id}
                lastUpdated={beatmapset.last_updated}
            />
        </div>
        <div className="md:w-1/2 md:justify-end mt-6 flex flex-col ml-auto whitespace-nowrap">
            {children}
        </div>
    </div>
);

export default MapsetHero;