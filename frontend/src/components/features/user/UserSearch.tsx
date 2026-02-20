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
        <div className="flex flex-col md:flex-row gap-2 items-center text-lg rounded-lg w-full z-1">
            <input
                onChange={(e) => setSearch(e.target.value)}
                className="px-4 py-2 bg-zinc-800/80 placeholder:text-zinc-500 rounded-lg w-full md:min-w-[400px] border border-zinc-900"
                placeholder="Search users..."
            />
            <div className="flex flex-row gap-2 w-full md:w-auto items-center">
                <h1 className="text-md hidden md:block text-zinc-200">Sort:</h1>
                <select
                    className="px-2 py-2 rounded-lg bg-zinc-800/80 text-zinc-200 text-white border border-zinc-900 w-full md:w-auto"
                    value={sort}
                    onChange={(e) => setSort(e.target.value)}
                >
                    <option value="playcount desc">more plays</option>
                    <option value="playcount asc">less plays</option>
                    <option value="map_count desc">more maps</option>
                    <option value="map_count asc">less maps</option>
                    {/*uncomment when it works*/}
                    {/*<option value="favourites desc">more favourites</option>*/}
                    {/*<option value="favourites asdc">less favourites</option>*/}
                    {/*<option value="comments desc">more comments</option>*/}
                    {/*<option value="comments asc">less comments</option>*/}
                </select>
            </div>
        </div>
    );
};

export default UserSearch;