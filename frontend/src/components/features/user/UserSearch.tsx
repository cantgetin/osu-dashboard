import {useEffect, useRef, useState} from "react";
import debounce from "lodash.debounce";
import {fetchMapsetsProps} from "../../../store/mapsetsSlice.ts";

interface UserSearchProps {
    update: (props: fetchMapsetsProps) => void
}

const UserSearch = (props: UserSearchProps) => {
    const [search, setSearch] = useState<string>('');
    const [sort, setSort] = useState<string>('');

    const debouncedUpdate = debounce(() => {
        props.update({
            search: search,
            status: status,
            sort: sort.split(' ')[0],
            direction: sort.split(' ')[1],
        });
    }, 500);

    const isFirstRun = useRef(true);

    useEffect(() => {
        if (isFirstRun.current) {
            isFirstRun.current = false;
            return;
        }

        debouncedUpdate();

        return () => {
            debouncedUpdate.cancel();
        };
    }, [search, status, sort]);

    return (
        <div className="flex flex-col md:flex-row gap-2 items-center text-lg rounded-lg w-full z-10">
            <input
                onChange={(e) => setSearch(e.target.value)}
                className="px-4 py-2 bg-zinc-800 bg-opacity-80 rounded-lg w-full md:min-w-[400px] border border-zinc-900"
                placeholder="Search users"
            />
            <div className="flex flex-row gap-2 w-full md:w-auto items-center">
                <h1 className="text-md hidden md:block">Sort:</h1>
                <select
                    className="px-2 py-2 rounded-lg bg-zinc-800 bg-opacity-80 text-white border border-zinc-900 w-full md:w-auto"
                    value={sort}
                    onChange={(e) => setSort(e.target.value)}
                >
                    <option value="">default</option>
                    <option value="last_playcount desc">more playcount</option>
                    <option value="last_playcount asc">less playcount</option>
                    <option value="created_at asc">oldest</option>
                    <option value="created_at desc">newest</option>
                </select>
            </div>
        </div>
    );
};

export default UserSearch;