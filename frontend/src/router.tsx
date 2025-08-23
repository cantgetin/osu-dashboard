import {createBrowserRouter, Outlet, ScrollRestoration} from "react-router-dom";
import Main from "./pages/Main.tsx";
import User from "./pages/User.tsx";
import Beatmapset from "./pages/Beatmapset.tsx";
import Users from "./pages/Users.tsx";
import Beatmapsets from "./pages/Beatmapsets.tsx";
import Authorize from "./pages/Authorize.tsx";
import NotFound from "./pages/NotFound.tsx";
import Tracks from "./pages/Tracks.tsx";

export const router = createBrowserRouter([
    {
        element: (
            <>
                <ScrollRestoration/>
                <Outlet/>
            </>
        ),
        children: [
            {
                path: "/",
                element: <Main/>,
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
            },
            {
                path: "/authorize",
                element: <Authorize/>,
            },
            {
                path: "/tracks",
                element: <Tracks/>,
            },
            {
                path: "*",
                element: <NotFound/>,
            },
        ]
    }


]);