import Header from "./Header.tsx";
import Footer from "./Footer.tsx";

interface LayoutProps {
    children: React.ReactNode
    className?: string
}

const Layout = (props: LayoutProps) => {
    return (
        <>
            <Header/>
            <div className="pt-14 min-h-screen w-full flex md:justify-center sm:justify-start">
                <div className={`p-10 container w-[1152px] max-w-[1152px] min-w-[400px] ${props.className}`}>
                    {props.children}
                </div>
            </div>
            <Footer/>
        </>
    );
};

export default Layout;