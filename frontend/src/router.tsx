import {createBrowserRouter} from "react-router-dom";
import User from "./pages/User.tsx";
import Beatmapset from "./pages/Beatmapset.tsx";
import Users from "./pages/Users.tsx";
import Beatmapsets from "./pages/Beatmapsets.tsx";
import AddUser from "./pages/AddUser.tsx";
import Main from "./pages/Main.tsx";
import NotFound from "./pages/NotFound.tsx";

export const router = createBrowserRouter([
    {
        path: "/",
        element: <Main/>,
    },
    {
        path: "/user/:userId",
        element: <User/>,
    },
    {
        path: "/users/add",
        element: <AddUser/>,
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
    },
    {
        path: "*",
        element: <NotFound />,
    },
]);