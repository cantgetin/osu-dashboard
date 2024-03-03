import Header from "./Header.tsx";
import Content from "./Content.tsx";
import Footer from "./Footer.tsx";

interface LayoutProps {
    children: React.ReactNode
    className?: string
}

const Layout = (props: LayoutProps) => {
    return (
        <>
            <Header/>
            <Content className="min-h-screen w-full flex md:justify-center sm:justify-start">
                <div className={`container w-[1152px] max-w-[1152px] ${props.className}`}>
                    {props.children}
                </div>
            </Content>
            <Footer/>
        </>
    );
};

export default Layout;