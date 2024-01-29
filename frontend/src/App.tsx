import Header from "./components/Header.tsx";
import Image from "../public/screen.png"
import Content from "./components/Content.tsx";

function App() {

    return (
        <>
            <Header/>
            <Content className="p-4 flex z-10">
                <div className="w-1/2 h-[calc(100vh-100px)] overflow-auto flex flex-col items-center justify-center">
                    <h1 className="text-8xl leading-tight">
                        Your personal osu! map dashboard
                    </h1>
                    <h1 className="text-4xl text-gray-500">
                        See your map stats in dynamic, track your map plays, favourites, comments
                    </h1>
                </div>
                <div>
                    <img className="absolute z-0 opacity-70" src={Image}/>
                </div>
            </Content>
        </>
    )
}

export default App
