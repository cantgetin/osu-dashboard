import { useState, useRef, useEffect } from "react";
import { useNavigate } from "react-router-dom";
import { createPortal } from "react-dom";

interface SearchResult {
    id: number;
    type: "user" | "mapset";
    title: string;
    picture_url: string;
}

const Search = () => {
    const [query, setQuery] = useState("");
    const [results, setResults] = useState<SearchResult[]>([]);
    const [isOpen, setIsOpen] = useState(false);
    const [isLoading, setIsLoading] = useState(false);
    const [selectedIndex, setSelectedIndex] = useState(-1);
    const [dropdownPosition, setDropdownPosition] = useState({ top: 0, left: 0 });
    const containerRef = useRef<HTMLDivElement>(null);
    const dropdownRef = useRef<HTMLDivElement>(null);
    const inputRef = useRef<HTMLInputElement>(null);
    const navigate = useNavigate();

    // Update dropdown position
    useEffect(() => {
        if (isOpen && inputRef.current) {
            const rect = inputRef.current.getBoundingClientRect();
            setDropdownPosition({
                top: rect.bottom + 8,
                left: rect.left,
            });
        }
    }, [isOpen, query]);

    // Debounced search
    useEffect(() => {
        if (query.trim().length < 2) {
            setResults([]);
            setIsOpen(false);
            return;
        }

        const timeoutId = setTimeout(async () => {
            setIsLoading(true);
            try {
                const response = await fetch(`/api/search/${encodeURIComponent(query.trim())}`, { method: 'POST' });
                if (response.ok) {
                    const data = await response.json();
                    setResults(data || []);
                    setIsOpen(true);
                }
            } catch (error) {
                console.error("Search failed:", error);
                setResults([]);
            } finally {
                setIsLoading(false);
            }
        }, 300);

        return () => clearTimeout(timeoutId);
    }, [query]);

    // Close dropdown when clicking outside
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            const target = event.target as Node;
            const isInsideContainer = containerRef.current?.contains(target);
            const isInsideDropdown = dropdownRef.current?.contains(target);
            
            if (!isInsideContainer && !isInsideDropdown) {
                setIsOpen(false);
                setSelectedIndex(-1);
            }
        };

        document.addEventListener("mousedown", handleClickOutside);
        return () => document.removeEventListener("mousedown", handleClickOutside);
    }, []);

    const handleSelect = (result: SearchResult) => {
        if (result.type === "user") {
            navigate(`/user/${result.id}`);
        } else if (result.type === "mapset") {
            navigate(`/beatmapset/${result.id}`);
        }
        setQuery("");
        setIsOpen(false);
        setSelectedIndex(-1);
    };

    const handleKeyDown = (e: React.KeyboardEvent) => {
        if (!isOpen || results.length === 0) return;

        switch (e.key) {
            case "ArrowDown":
                e.preventDefault();
                setSelectedIndex((prev) => (prev < results.length - 1 ? prev + 1 : prev));
                break;
            case "ArrowUp":
                e.preventDefault();
                setSelectedIndex((prev) => (prev > 0 ? prev - 1 : prev));
                break;
            case "Enter":
                e.preventDefault();
                if (selectedIndex >= 0 && selectedIndex < results.length) {
                    handleSelect(results[selectedIndex]);
                }
                break;
            case "Escape":
                setIsOpen(false);
                setSelectedIndex(-1);
                inputRef.current?.blur();
                break;
        }
    };

    const dropdown = (
        <div
            ref={dropdownRef}
            style={{ top: dropdownPosition.top, left: dropdownPosition.left }}
            className="fixed w-80 bg-zinc-900 rounded-lg shadow-2xl shadow-black overflow-hidden"
        >
            <div className="max-h-96 overflow-y-auto">
                {results.map((result, index) => (
                    <div
                        key={`${result.type}-${result.id}`}
                        onClick={() => handleSelect(result)}
                        onMouseEnter={() => setSelectedIndex(index)}
                        className={`flex items-center gap-3 px-4 py-3 cursor-pointer transition-colors
                            ${selectedIndex === index ? "bg-zinc-800" : "hover:bg-zinc-800/60"}`}
                    >
                        {result.type === "user" ? (
                            <img
                                src={result.picture_url || `https://a.ppy.sh/${result.id}`}
                                alt=""
                                className="w-10 h-10 rounded-full object-cover flex-shrink-0"
                            />
                        ) : (
                            <img
                                src={result.picture_url || ""}
                                alt=""
                                className="w-14 h-10 rounded object-cover flex-shrink-0 bg-zinc-700"
                            />
                        )}
                        <div className="flex-1 min-w-0">
                            <p className="text-base font-medium truncate">{result.title}</p>
                            <p className="text-sm text-zinc-400 capitalize">{result.type}</p>
                        </div>
                        <span className={`text-xs px-2 py-1 rounded font-medium uppercase tracking-wide
                            ${result.type === "user"
                                ? "bg-zinc-700 text-zinc-300"
                                : "bg-zinc-600 text-zinc-200"}`}
                        >
                            {result.type === "user" ? "User" : "Map"}
                        </span>
                    </div>
                ))}
            </div>
        </div>
    );

    const noResultsDropdown = (
        <div
            style={{ top: dropdownPosition.top, left: dropdownPosition.left }}
            className="fixed w-80 bg-zinc-900 rounded-lg shadow-black shadow-2xl p-5"
        >
            <p className="text-base text-zinc-400 text-center">No results found</p>
        </div>
    );

    return (
        <div ref={containerRef}>
            <div className="relative">
                <input
                    ref={inputRef}
                    type="text"
                    value={query}
                    onChange={(e) => setQuery(e.target.value)}
                    onFocus={() => query.trim().length >= 2 && results.length > 0 && setIsOpen(true)}
                    onKeyDown={handleKeyDown}
                    placeholder="Search users or maps..."
                    className="bg-zinc-800/80 rounded-md w-36 sm:w-48 md:w-64 px-4 h-9 text-base
                        placeholder:text-zinc-500 focus:outline-none focus:ring-1 focus:ring-zinc-500
                        transition-all duration-200"
                />
                {isLoading && (
                    <div className="absolute right-3 top-1/2 -translate-y-1/2">
                        <div className="w-4 h-4 border-2 border-zinc-600 border-t-zinc-300 rounded-full animate-spin" />
                    </div>
                )}
            </div>

            {isOpen && results.length > 0 && createPortal(dropdown, document.body)}

            {isOpen && query.trim().length >= 2 && results.length === 0 && !isLoading && createPortal(noResultsDropdown, document.body)}
        </div>
    );
};

export default Search;
