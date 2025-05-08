interface UserStatsSummaryProps {
    user: User
}

const MapStatsSummary = (props: UserStatsSummaryProps) => {
    const types = Object.keys(props.user.user_map_counts) as Array<keyof UserMapCounts>

    return (
        <div className="flex flex-wrap gap-2">
            {types.map((type) => {
                const count = props.user.user_map_counts[type];
                if (count === 0) return null;
                return (
                    <div key={type} className="bg-zinc-800 px-2 py-1 rounded-lg text-sm cursor-pointer h-7">
                        {type} {count}
                    </div>
                );
            })}
        </div>
    );
};

export default MapStatsSummary;