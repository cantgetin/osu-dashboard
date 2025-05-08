import Layout from "../components/ui/Layout.tsx";
import SystemStats from "../components/business/SystemStats.tsx";
import Features from "../components/business/Features.tsx";
import HeroSection from "../components/business/HeroSection.tsx";

const Main = () => {
    return (
        <Layout className="py-5 md:py-10 flex flex-col gap-5 md:gap-10" title="Home">
            <HeroSection/>
            <SystemStats/>
            <Features/>
        </Layout>
    );
};

export default Main;