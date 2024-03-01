
interface SearchBarProps {
    className?: string
    placeholder?: string
}

const SearchBar = (props: SearchBarProps) => {
    return (
        <input type="text" className={props.className} placeholder={props.placeholder}/>
    );
};

export default SearchBar;