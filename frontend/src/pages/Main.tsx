import Layout from "../components/ui/Layout.tsx";
import HeroSection from "../components/features/common/HeroSection.tsx";
import SystemStats from "../components/features/stats/SystemStats.tsx";
import Features from "../components/features/common/Features.tsx";

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