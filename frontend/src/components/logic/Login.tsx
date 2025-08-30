// LoginComponent.tsx
import { useState, useRef, useEffect } from "react";

interface LoginComponentProps {
    isLoggedIn: boolean;
    onLogin: (username: string, password: string) => void;
    onLogout: () => void;
    user?: {
        nickname: string;
        avatar: string;
    };
}

const Login = ({ isLoggedIn, onLogin, onLogout, user }: LoginComponentProps) => {
    const [showDropdown, setShowDropdown] = useState(false);
    const [username, setUsername] = useState("");
    const [password, setPassword] = useState("");
    const dropdownRef = useRef<HTMLDivElement>(null);
    const buttonRef = useRef<HTMLButtonElement>(null);

    const handleLogin = (e: React.FormEvent) => {
        e.preventDefault();
        onLogin(username, password);
        setUsername("");
        setPassword("");
    };

    // Calculate dropdown position based on button position
    const getDropdownPosition = () => {
        if (!buttonRef.current) return {};

        const rect = buttonRef.current.getBoundingClientRect();
        return {
            top: rect.bottom + window.scrollY + 4, // 4px margin
            right: window.innerWidth - rect.right - window.scrollX,
        };
    };

    // Close dropdown when clicking outside
    useEffect(() => {
        const handleClickOutside = (event: MouseEvent) => {
            if (dropdownRef.current && !dropdownRef.current.contains(event.target as Node)) {
                setShowDropdown(false);
            }
        };

        document.addEventListener("mousedown", handleClickOutside);
        return () => {
            document.removeEventListener("mousedown", handleClickOutside);
        };
    }, []);

    return (
        <div className="flex items-center">
            {isLoggedIn && user ? (
                <div className="relative">
                    <button
                        ref={buttonRef}
                        className="flex items-center space-x-2 focus:outline-none"
                        onClick={() => setShowDropdown(!showDropdown)}
                    >
                        <img
                            src={user.avatar}
                            alt={user.nickname}
                            className="w-8 h-8 rounded-full border border-zinc-600"
                        />
                        <span className="text-white hidden sm:block">{user.nickname}</span>
                        <svg
                            className={`w-4 h-4 text-white transition-transform ${showDropdown ? 'rotate-180' : ''}`}
                            fill="none"
                            stroke="currentColor"
                            viewBox="0 0 24 24"
                        >
                            <path strokeLinecap="round" strokeLinejoin="round" strokeWidth={2} d="M19 9l-7 7-7-7" />
                        </svg>
                    </button>

                    {/* Dropdown with calculated positioning */}
                    {showDropdown && (
                        <div
                            ref={dropdownRef}
                            className="fixed bg-zinc-800 rounded-md shadow-lg py-1 z-50 border border-zinc-700 w-48"
                            style={getDropdownPosition()}
                        >
                            <a href="#" className="block px-4 py-2 text-sm text-white hover:bg-zinc-700 transition-colors">My Profile</a>
                            <a href="#" className="block px-4 py-2 text-sm text-white hover:bg-zinc-700 transition-colors">Feedback</a>
                            <button
                                onClick={() => {
                                    onLogout();
                                    setShowDropdown(false);
                                }}
                                className="block w-full text-left px-4 py-2 text-sm text-white hover:bg-zinc-700 transition-colors"
                            >
                                Logout
                            </button>
                        </div>
                    )}
                </div>
            ) : (
                <div className="flex items-center space-x-3">
                    <form onSubmit={handleLogin} className="flex items-center space-x-2">
                        <input
                            type="text"
                            placeholder="Username"
                            value={username}
                            onChange={(e) => setUsername(e.target.value)}
                            className="bg-zinc-800 text-white px-3 py-1.5 rounded text-sm w-24 sm:w-32 border border-zinc-600 focus:border-blue-500 focus:outline-none transition-colors"
                            required
                        />
                        <input
                            type="password"
                            placeholder="Password"
                            value={password}
                            onChange={(e) => setPassword(e.target.value)}
                            className="bg-zinc-800 text-white px-3 py-1.5 rounded text-sm w-24 sm:w-32 border border-zinc-600 focus:border-blue-500 focus:outline-none transition-colors"
                            required
                        />
                        <button
                            type="submit"
                            className="bg-blue-600 hover:bg-blue-700 text-white px-4 py-1.5 rounded text-sm transition-colors"
                        >
                            Login
                        </button>
                    </form>
                </div>
            )}
        </div>
    );
};

export default Login;