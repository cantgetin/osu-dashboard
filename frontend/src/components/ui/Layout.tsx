import Header from "./Header.tsx";
import Footer from "./Footer.tsx";
import {useEffect} from "react";

interface LayoutProps {
    children: React.ReactNode
    className?: string
    title: string
}

const Layout = (props: LayoutProps) => {
    useEffect(() => {
        if (props.title) {
            document.title = "Dashboard | " + props.title;
        }
    }, [props.title]);

    return (
        <>
            <Header/>
            <div className="pt-14 min-h-screen w-full flex justify-center">
                <div className={`p-4 sm:p-6 md:p-10 w-full max-w-[1152px] min-w-0 ${props.className}`}>
                    {props.children}
                </div>
            </div>
            <Footer/>
        </>
    );
};

export default Layout;