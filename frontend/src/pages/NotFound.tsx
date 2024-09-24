import Layout from "../components/ui/Layout.tsx";
import Button from "../components/ui/Button.tsx";
import { useNavigate } from "react-router-dom";

const NotFound = () => {
    const navigate = useNavigate();

    return (
        <Layout className="flex md:justify-center sm:justify-start">
            <div className="flex flex-col w-screen justify-center items-center gap-2">
                <h1 className="text-8xl leading-tight drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">Error 404</h1>
                <h1 className="text-5xl text-gray-400 drop-shadow-[0_1.2px_1.2px_rgba(0,0,0,1)]">
                    The page you requested could not be found.
                </h1>
                <Button onClick={() => navigate("/")}
                        className="text-xl rounded-md px-6 py-3 my-6 bg-green-800 hover:bg-green-900"
                        content="Home"
                />
            </div>
        </Layout>
    );
};

export default NotFound;