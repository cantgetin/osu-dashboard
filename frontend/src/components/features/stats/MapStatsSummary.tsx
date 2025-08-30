import Tags from "@/components/features/common/Tags.tsx";

interface UserStatsSummaryProps {
    user: User
}

const MapStatsSummary = (props: UserStatsSummaryProps) => {
    const types = Object.keys(props.user.user_map_counts) as Array<keyof UserMapCounts>

    // Create tags array with {type} {count} format
    const tags = types
        .map(type => {
            const count = props.user.user_map_counts[type];
            return count > 0 ? `${type} ${count}` : null;
        })
        .filter(Boolean) as string[]; // Filter out null values and type as string array

    return (
        <Tags tags={tags} colorized={true}/>
    );
};

export default MapStatsSummary;