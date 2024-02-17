import * as ReactDOM from "react-dom/client";
import {router} from "./router.tsx";
import {RouterProvider} from "react-router-dom";
import "./index.css";
import '@fontsource/roboto';
import {Provider} from "react-redux";
import {store} from "./store/store.ts";

ReactDOM.createRoot(document.getElementById("root")!).render(
    <Provider store={store}>
        <RouterProvider router={router}/>
    </Provider>
);