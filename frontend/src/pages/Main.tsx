import Layout from "../components/ui/Layout.tsx";
import Image1 from "../../public/screen1.png"
import Image2 from "../../public/screen2.png"

const Main = () => {
    return (
        <Layout className="p-10 py-20 flex flex-col gap-10">
            <h1 className="text-7xl leading-tight">
                Your personal osu! map dashboard
            </h1>
            <h1 className="text-4xl text-gray-500">
                See all your map stats in dynamic, track your map plays, favourites, comments
            </h1>
            <div className="px-4">
                <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Image1}/>
            </div>
            <h1 className="text-4xl text-gray-500">
                See other users stats, mapset stats, beatmap stats
            </h1>
            <div className="px-4">
                <img className="z-0 opacity-70 rounded-lg overflow-hidden" src={Image2}/>
            </div>
        </Layout>
    );
};

export default Main;