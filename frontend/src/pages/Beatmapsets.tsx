import {useEffect, useState} from "react";
import List from "../components/List.tsx";
import Mapset from "../components/Mapset.tsx";
import debounce from 'lodash.debounce';
import LoadingSpinner from "../components/ui/LoadingSpinner.tsx";
import Layout from "../components/ui/Layout.tsx";

const Beatmapsets = () => {
    const [mapsets, setMapsets] = useState<Mapset[]>([]);

    const [search, setSearch] = useState<string>('');
    const [sort, setSort] = useState<string>('');
    const [status, setStatus] = useState<string>('');

    useEffect(() => {
        const debouncedSearchSetter = debounce(setSearch, 500);

        return () => {
            debouncedSearchSetter.cancel();
        };
    }, []);

    useEffect(() => {
        const fetchData = async () => {
            let trueSort = sort.split(' ')[0]
            let direction = sort.split(' ')[1]


            const response =
                await fetch(`/api/beatmapset/list?search=${search}&status=${status}&sort=${trueSort}&direction=${direction}`);
            const userData = await response.json();
            setMapsets(JSON.parse(JSON.stringify(userData)) as Mapset[]);
        };

        const timerId = setTimeout(fetchData, 500);
        return () => clearTimeout(timerId);
    }, [search, status, sort]);

    return (
        <Layout className="p-10 flex flex-col gap-2 min-w-[800px]">
            <div className="flex gap-2 items-center text-lg rounded-lg min-w-[800px] z-10">
                <input
                    onChange={(e) => setSearch(e.target.value)}
                    className="px-4 py-2 bg-zinc-800 rounded-lg min-w-[400px] w-full border border-zinc-900"
                    placeholder="Search beatmapsets"
                />
                <h1 className="text-md">Status:</h1>
                <select
                    className="px-2 py-2  rounded-lg bg-zinc-800 text-white border border-zinc-900"
                    value={status}
                    onChange={(e) => setStatus(e.target.value)}
                >
                    <option value="">any</option>
                    <option value="graveyard">graveyard</option>
                    <option value="wip">wip</option>
                    <option value="pending">pending</option>
                    <option value="ranked">ranked</option>
                    <option value="approved">approved</option>
                    <option value="qualified">qualified</option>
                    <option value="loved">loved</option>
                </select>
                <h1 className="text-md">Sort:</h1>
                <select
                    className="px-2 py-2 rounded-lg bg-zinc-800 text-white border border-zinc-900"
                    value={sort}
                    onChange={(e) => setSort(e.target.value)}
                >
                    <option value="">default</option>
                    <option value="last_playcount desc">more playcount</option>
                    <option value="last_playcount asc">less playcount </option>
                    <option value="created_at asc">oldest</option>
                    <option value="created_at desc">newest</option>
                </select>
            </div>
            {mapsets != null && mapsets.length > 0 ?
                <List
                    items={mapsets}
                    renderItem={(m: Mapset) =>
                        <Mapset
                            key={m.id}
                            map={m}
                            showMapper={true}
                            className="min-w-[800px]"
                        />
                    }
                    className="grid xl:grid-cols-1 gap-2 py-5"
                />
                : <LoadingSpinner/>
            }
        </Layout>
    );
};

export default Beatmapsets;