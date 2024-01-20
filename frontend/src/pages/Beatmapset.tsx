import {useParams} from "react-router-dom";
import Header from "../components/Header.tsx";

const Beatmapset = () => {
    const {mapId} = useParams();

    return (
        <>
            <Header/>
            <div>
                {mapId}
            </div>
        </>
    );
};

export default Beatmapset;