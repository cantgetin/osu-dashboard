import aveta from "aveta";

interface StatCardProps {
    value: number;
    previousValue: number;
    label: string;
    color: string;
}

const StatCard = ({ value, previousValue, label, color }: StatCardProps) => {
    const difference = value - previousValue;

    return (
        <div className={`text-4xl flex gap-2 items-center w-full justify-end drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)] ${color}`}>
            <h1>{aveta(value)} {label}</h1>
            <h1 className="text-xl">▲</h1>
            <h1>{aveta(difference)}</h1>
        </div>
    );
};

interface StatsComparisonProps {
    lastStats: any;
    penultimateStats: any;
}

const StatsComparison = ({ lastStats, penultimateStats }: StatsComparisonProps) => (
    <div className="flex flex-col justify-center items-center">
        <StatCard
            value={lastStats.favourite_count}
            previousValue={penultimateStats.favourite_count}
            label="Favorites"
            color="text-pink-300"
        />
        <StatCard
            value={lastStats.play_count}
            previousValue={penultimateStats.play_count}
            label="Plays"
            color="text-green-300"
        />
        <StatCard
            value={lastStats.comments_count}
            previousValue={penultimateStats.comments_count}
            label="Comments"
            color="text-red-300"
        />
    </div>
);

export default StatsComparison;