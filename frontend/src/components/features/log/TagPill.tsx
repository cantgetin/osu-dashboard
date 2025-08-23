interface TagPillProps {
    text: string;
    color: 'blue' | 'green' | 'emerald' | 'amber' | 'red' | 'purple';
    className?: string;
}

const colorMap = {
    blue: 'bg-blue-900 bg-opacity-60 text-blue-300',
    green: 'bg-green-900 bg-opacity-60 text-green-300',
    emerald: 'bg-emerald-900 bg-opacity-60 text-emerald-300',
    amber: 'bg-amber-900 bg-opacity-60 text-amber-300',
    red: 'bg-red-900 bg-opacity-60 text-red-300',
    purple: 'bg-purple-900 bg-opacity-60 text-purple-300',
};

export const TagPill = ({text, color, className = ''}: TagPillProps) => {
    return (
        <span
            className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${colorMap[color]} ${className}`}>
            {text}
        </span>
    );
};