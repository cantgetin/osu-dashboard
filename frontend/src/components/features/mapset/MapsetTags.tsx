interface MapsetTagsProps {
    tags: string;
}

const MapsetTags = ({ tags }: MapsetTagsProps) => (
    <div className="flex gap-1 flex-wrap max-w-[600px]">
        {tags && tags.split(' ').map((tag, index) => (
            <div
                key={index}
                className="text:xs px-1 bg-zinc-800 md:px-2 md:py-1 rounded-lg md:text-sm cursor-pointer h-7"
            >
                {tag}
            </div>
        ))}
    </div>
);

export default MapsetTags;