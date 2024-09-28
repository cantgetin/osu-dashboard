import Layout from "../components/ui/Layout.tsx";
import Charts from "../images/charts.png"
import User from "../images/user.png"
import Mapset from "../images/mapset.png"
import Diagrams from "../images/diagrams.png"
import Button from "../components/ui/Button.tsx";
import Summarized from "../images/summarized.png"
import Filtering from "../images/filtering.png"

const Main = () => {
    // @ts-ignore
    return (
        <Layout className="py-10 flex flex-col gap-10">
            <div className="flex perspective-600 overflow-hidden">
                <div className="w-3/5 h-[calc(87vh)] flex flex-col gap-10 justify-center">
                    <h1 className="text-7xl leading-tight drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                        Your personal osu! map dashboard
                    </h1>
                    <h1 className="text-3xl text-gray-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                        See all your map statistics in dynamic, track your daily map plays, favourites, comments
                    </h1>
                    <div className="flex gap-10">
                        <Button onClick={() => {
                        }}
                                className="text-xl rounded-md p-4 bg-green-800 w-1/4 hover:bg-green-900"
                                content="Start for free"
                        />
                        <Button onClick={() => window.location.href = '#features'}
                                className="text-xl rounded-md p-4 bg-zinc-700 w-1/4 hover:bg-zinc-800"
                                content="Learn more"
                        />
                    </div>
                </div>
                <img
                    className="absolute -z-10 opacity-15 rounded-lg overflow-hidden rotate transform rotate-x-12 -mt-10"
                    src={Summarized}/>
            </div>

            <section id="features">
            </section>
            <div className="py-10 flex flex-col gap-10">

                <h1 className="text-5xl leading-tight">
                    Features
                </h1>
                <h1 className="text-3xl text-gray-400">
                    Track your total daily map plays, favourites, comments on charts
                </h1>
                <div className="px-4">
                    <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Charts}/>
                </div>
                <h1 className="text-3xl text-gray-400">
                    Track summarized user statistics for last 24 hours and 7 days
                </h1>
                <div className="px-4">
                    <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={User}/>
                </div>
                <h1 className="text-3xl text-gray-400">
                    Filter and sort your mapsets with multiple options
                </h1>
                <div className="px-4">
                    <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Filtering}/>
                </div>
                <h1 className="text-3xl text-gray-400">
                    Track specific mapset and beatmap statistics
                </h1>
                <div className="px-4">
                    <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Mapset}/>
                </div>
                <h1 className="text-3xl text-gray-400">
                    See your summarized genre, bpm, tag, starrate diagrams
                </h1>
                <div className="px-4">
                    <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Diagrams}/>
                </div>
            </div>

        </Layout>
    );
};

export default Main;