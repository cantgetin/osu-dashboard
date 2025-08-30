import MyButton from "../../ui/MyButton.tsx";
import {FaExternalLinkAlt} from "react-icons/fa";

interface MapsetHeaderProps {
    title: string;
    artist: string;
    id: number;
}

const MapsetHeader = ({ title, artist, id }: MapsetHeaderProps) => {
    const externalLinkOnClick = () => window.open(`https://osu.ppy.sh/s/${id}`, "_blank");

    return (
        <div>
            <div className="flex gap-4 items-center">
                <h1 className="text-2xl md:text-5xl drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">{title}</h1>
                <MyButton
                    onClick={externalLinkOnClick}
                    className="bg-zinc-800 rounded-md p-1 h-6"
                    content={<FaExternalLinkAlt className="h-3"/>}
                />
            </div>
            <h1 className="text-xl md:text-4xl text-zinc-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">by {artist}</h1>
        </div>
    );
};

export default MapsetHeader;