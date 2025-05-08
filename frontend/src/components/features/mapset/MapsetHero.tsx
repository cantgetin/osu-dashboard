import MapsetHeader from "./MapsetHeader.tsx";
import MapsetTags from "./MapsetTags.tsx";
import MapsetMeta from "./MapsetMeta.tsx";


interface MapsetHeroProps {
    beatmapset: Mapset;
    children: React.ReactNode;
}

const MapsetHero = ({beatmapset, children}: MapsetHeroProps) => (
    <div className="p-5 absolute inset-0 flex h-[550px]">
        <div className="w-1/2 justify-end flex flex-col gap-6">
            <MapsetHeader
                title={beatmapset.title}
                artist={beatmapset.artist}
                id={beatmapset.id}
            />
            <MapsetTags tags={beatmapset.tags}/>
            <MapsetMeta
                status={beatmapset.status}
                creator={beatmapset.creator}
                userId={beatmapset.user_id}
                lastUpdated={beatmapset.last_updated}
            />
        </div>
        {children}
    </div>
);

export default MapsetHero;