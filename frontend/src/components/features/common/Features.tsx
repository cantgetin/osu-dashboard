import Charts from "../../../images/charts.png";
import User from "../../../images/user.png";
import Filtering from "../../../images/filtering.png";
import Mapset from "../../../images/mapset.png";
import Diagrams from "../../../images/diagrams.png";

const Features = () => {
    return (
        <section id="features">
            <div className="flex flex-col gap-5 md:gap-10">
                <h1 className="text-3xl md:text-5xl leading-tight">
                    Features
                </h1>
                <h1 className="text-xl md:text-3xl text-gray-400">
                    Track your total daily map plays, favourites, comments on charts
                </h1>
                <div className="px-2 md:px-4">
                    <img className="z-0 rounded-lg overflow-hidden w-full" src={Charts}/>
                </div>
                <h1 className="text-xl md:text-3xl text-gray-400">
                    Track summarized user statistics for last 24 hours and 7 days
                </h1>
                <div className="px-2 md:px-4">
                    <img className="z-0 rounded-lg overflow-hidden w-full" src={User}/>
                </div>
                <h1 className="text-xl md:text-3xl text-gray-400">
                    Filter and sort your mapsets with multiple options
                </h1>
                <div className="px-2 md:px-4">
                    <img className="z-0 rounded-lg overflow-hidden w-full" src={Filtering}/>
                </div>
                <h1 className="text-xl md:text-3xl text-gray-400">
                    Track specific mapset and beatmap statistics
                </h1>
                <div className="px-2 md:px-4">
                    <img className="z-0 rounded-lg overflow-hidden w-full" src={Mapset}/>
                </div>
                <h1 className="text-xl md:text-3xl text-gray-400">
                    See your summarized genre, bpm, tag, starrate diagrams
                </h1>
                <div className="px-2 md:px-4">
                    <img className="z-0 rounded-lg overflow-hidden w-full" src={Diagrams}/>
                </div>
            </div>
        </section>
    );
};

export default Features;