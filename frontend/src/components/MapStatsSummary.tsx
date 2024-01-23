interface UserStatsSummaryProps {
    data: UserCard
}


const types = ['graveyard', 'loved', 'nominated', 'pending', 'ranked', 'wip']

const MapStatsSummary = (props: UserStatsSummaryProps) => {
    return (
        <div className="flex flex-wrap gap-2">
            {
                types.map((type) => {
                    if (props.data.Mapsets.filter(o => o.status == type).length == 0) return null
                    else return (
                        <div key={type} className="bg-zinc-800 px-2 py-1 rounded-lg text-sm cursor-pointer">
                            {type} {props.data.Mapsets.filter(o => o.status == type).length}
                        </div>
                    )
                }
                )
            }
        </div>
    );
};

export default MapStatsSummary;