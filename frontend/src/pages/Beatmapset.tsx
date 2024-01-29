import {useParams} from "react-router-dom";
import Header from "../components/Header.tsx";
import Content from "../components/Content.tsx";

const Beatmapset = () => {
    const {mapId} = useParams();

    return (
        <>
            <Header/>
            <Content>
                {mapId}
            </Content>
        </>
    );
};

export default Beatmapset;