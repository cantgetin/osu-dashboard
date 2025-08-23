interface ProgressBarProps {
    value: number;
    max: number;
    className?: string;
    colorClass?: string;
}

export const ProgressBar = ({value, max, className = '', colorClass = 'bg-blue-500'}: ProgressBarProps) => {
    const percentage = Math.min(100, (value / max) * 100);

    return (
        <div className={`w-full bg-zinc-700 rounded-full h-2 ${className}`}>
            <div
                className={`h-full rounded-full ${colorClass}`}
                style={{width: `${percentage}%`}}
            />
        </div>
    );
};
export default ProgressBar;