interface MapsetTagsProps {
    tags: string | string[];
    colorized?: boolean;
}

const colorPool = ['#003f5c', '#58508d', '#34ab9c', "#ff6361", '#983a73'];

const getTagColor = (tag: string) => {
    const hash = tag.split('').reduce((acc, char) => char.charCodeAt(0) + acc, 0);
    const baseColor = colorPool[hash % colorPool.length];
    return subtleColorVariant(baseColor, hash);
};

const subtleColorVariant = (hexColor: string, seed: number) => {
    const r = parseInt(hexColor.substr(1, 2), 16);
    const g = parseInt(hexColor.substr(3, 2), 16);
    const b = parseInt(hexColor.substr(5, 2), 16);

    const variation = (seed % 41) - 20;
    const newR = Math.min(255, Math.max(0, r + (variation * 0.8)));
    const newG = Math.min(255, Math.max(0, g + (variation * 1.2)));
    const newB = Math.min(255, Math.max(0, b + (variation * 0.5)));

    return `#${Math.round(newR).toString(16).padStart(2, '0')}${Math.round(newG).toString(16).padStart(2, '0')}${Math.round(newB).toString(16).padStart(2, '0')}`;
};

const Tags = ({ tags, colorized = false }: MapsetTagsProps) => {
    const normalizedTags = Array.isArray(tags)
        ? tags
        : tags.split(/[\s,]+/).filter(tag => tag.trim() !== '');

    return (
        <div className="flex gap-1 flex-wrap max-w-[600px]">
            {normalizedTags.map((tag, index) => (
                <div
                    key={index}
                    className="text-xs px-1 md:px-2 md:py-1 rounded-lg md:text-sm cursor-pointer h-7 flex items-center"
                    style={colorized ? {
                        backgroundColor: getTagColor(tag),
                        color: '#ffffff',
                        textShadow: '2px 2px 3px rgba(0, 0, 0, 0.7)'
                    } : {
                        backgroundColor: '#27272a',
                        color: '#ffffff',
                        textShadow: '2px 2px 3px rgba(0, 0, 0, 0.7)'
                    }}
                >
                    {tag}
                </div>
            ))}
        </div>
    );
};

export default Tags;