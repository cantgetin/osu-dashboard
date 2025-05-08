import Button from "../ui/Button.tsx";
import Summarized from "../images/summarized.png"
import {useNavigate} from "react-router-dom";

const HeroSection = () => {
    const navigate = useNavigate();

    return (
        <div className="flex flex-col md:flex-row perspective-600 overflow-hidden">
            <div className="w-full md:w-3/5 h-auto md:h-[calc(87vh)] flex flex-col gap-5 md:gap-10 justify-center">
                <h1 className="text-4xl md:text-7xl leading-tight drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                    Your personal osu! map dashboard
                </h1>
                <h1 className="text-xl md:text-3xl text-gray-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                    See all your map statistics in dynamic, track your daily map plays, favourites, comments
                </h1>
                <div className="flex flex-col md:flex-row gap-3 md:gap-10">
                    <Button onClick={() => navigate('/authorize')}
                            className="text-lg md:text-xl rounded-md p-3 md:p-4 bg-green-800 w-full md:w-1/4 hover:bg-green-900"
                            content="Start for free"
                    />
                    <Button onClick={() => window.location.href = '#features'}
                            className="text-lg md:text-xl rounded-md p-3 md:p-4 bg-zinc-700 w-full md:w-1/4 hover:bg-zinc-800"
                            content="Learn more"
                    />
                </div>
            </div>
            <img
                className="absolute -z-10 opacity-15 rounded-lg overflow-hidden rotate transform rotate-x-12 mt-5 md:-mt-10 w-full md:w-auto"
                src={Summarized}/>
        </div>
    );
};

export default HeroSection;