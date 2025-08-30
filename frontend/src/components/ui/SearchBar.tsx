interface SearchBarProps {
    className?: string
    placeholder?: string
}

const SearchBar = (props: SearchBarProps) => {
    return (
        // <input
        //     onChange={(e) => setSearch(e.target.value)}
        //     className="px-4 py-2 bg-zinc-800 bg-opacity-80 rounded-lg w-full md:min-w-[400px] border border-zinc-900"
        //     placeholder="Search users"
        // />

    <input type="text" className={props.className} placeholder={props.placeholder}/>
)
    ;
};

export default SearchBar;