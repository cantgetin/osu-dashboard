interface MapsetTagsProps {
    tags: string;
}

const MapsetTags = ({ tags }: MapsetTagsProps) => (
    <div className="flex gap-1 flex-wrap max-w-[600px]">
        {tags && tags.split(' ').map((tag, index) => (
            <div
                key={index}
                className="bg-zinc-800 px-2 py-1 rounded-lg text-sm cursor-pointer h-7"
            >
                {tag}
            </div>
        ))}
    </div>
);

export default MapsetTags;