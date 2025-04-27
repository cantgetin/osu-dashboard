import Layout from "../components/ui/Layout.tsx";
import Button from "../components/ui/Button.tsx";
import Summarized from "../images/summarized.png"
import {useNavigate} from "react-router-dom";
import SystemStats from "../components/business/SystemStats.tsx";
import Features from "../components/business/Features.tsx";

const Main = () => {
    const navigate = useNavigate();

    // @ts-ignore
    return (
        <Layout className="py-10 flex flex-col gap-10" title="Home">
            <div className="flex perspective-600 overflow-hidden">
                <div className="w-3/5 h-[calc(87vh)] flex flex-col gap-10 justify-center">
                    <h1 className="text-7xl leading-tight drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                        Your personal osu! map dashboard
                    </h1>
                    <h1 className="text-3xl text-gray-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                        See all your map statistics in dynamic, track your daily map plays, favourites, comments
                    </h1>
                    <div className="flex gap-10">
                        <Button onClick={() => navigate('/authorize')}
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
            <SystemStats/>
            <Features/>
        </Layout>
    );
};

export default Main;