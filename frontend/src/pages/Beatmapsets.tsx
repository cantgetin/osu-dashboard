import Layout from "../components/ui/Layout.tsx";
import MapsetList from "../components/business/MapsetList.tsx";

const Beatmapsets = () => {
    return (
        <Layout className="flex flex-col gap-2 min-w-[800px]" title="Beatmaps">
            <MapsetList
                forUser={false}
                page={1}
                sort="created_at"
                direction="desc"
                showMapper={true}
            />
        </Layout>
    );
};

export default Beatmapsets;