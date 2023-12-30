import {createBrowserRouter} from "react-router-dom";
import App from "./App.tsx";
import User from "./pages/User.tsx";
import Beatmapset from "./pages/Beatmapset.tsx";
import Users from "./pages/Users.tsx";
import Beatmapsets from "./pages/Beatmapsets.tsx";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <App/>,
    },
    {
        path: "/user/:userId",
        element: <User/>,
    },
    {
        path: "/beatmapset/:mapId",
        element: <Beatmapset/>,
    },
    {
        path: "/users",
        element: <Users/>,
    },
    {
        path: "/beatmapsets",
        element: <Beatmapsets/>,
    }
]);