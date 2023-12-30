
interface MapCardProps {
    map: Mapset
}

const MapsetSummary = (props: MapCardProps) => {
    return (
        <div className="flex bg-zinc-900 text-white w-full rounded-lg overflow-hidden">
            <div>
                <img src={props.map.covers.card} className='h-full w-64' alt="map bg"
                     style={{objectFit: 'cover'}}/>
            </div>
            <div className="px-2 py-1 mr-auto">
                <a className="text-xl underline" href={`/beatmapset/${props.map.id}`}>{props.map.artist} - {props.map.title}</a>
                {/*<div className="text-xl">{props.map.artist} - {props.map.title}</div>*/}
                <div className="flex gap-2 justify-left items-baseline">
                    <h1 className="text-xl text-yellow-200">{Object.values(props.map.mapset_stats)[1].play_count} plays</h1>
                    <h1 className="text-sm h-full text-orange-200">were {Object.values(props.map.mapset_stats)[0].play_count} plays</h1>
                </div>
                <div className='text-xs text-zinc-500'>
                    {/*{getMapRemainingPendingTime(props.map.last_updated) === '' ? props.map.status : getMapRemainingPendingTime(props.map.last_updated)}*/}
                    pending for 27 days 15 hours 59 minutes
                </div>
            </div>
            <div className="p-4 flex gap-2 justify-center items-center">
                {
                    Object.values(props.map.mapset_stats)[1].play_count - Object.values(props.map.mapset_stats)[0].play_count > 0 ?
                        <>
                            <h1 className="text-xs text-green-300">▲</h1>
                            <h1 className="text-2xl text-green-300">
                                {Object.values(props.map.mapset_stats)[1].play_count - Object.values(props.map.mapset_stats)[0].play_count}
                            </h1>
                        </>
                        : null
                }
            </div>
        </div>
    );
};

export default MapsetSummary;