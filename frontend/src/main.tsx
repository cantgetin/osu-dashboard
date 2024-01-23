import * as ReactDOM from "react-dom/client";
import { router } from "./router.tsx";
import { RouterProvider } from "react-router-dom";
import "./index.css";
import '@fontsource/roboto';

ReactDOM.createRoot(document.getElementById("root")!).render(
    <RouterProvider router={router} />
);