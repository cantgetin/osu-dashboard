import { Link, useLocation } from 'react-router-dom';

interface HeaderLinkProps {
    to: string;
    children: React.ReactNode;
}

const HeaderLink = ({ to, children }: HeaderLinkProps) => {
    const location = useLocation();
    const isActive = location.pathname === to;

    return (
        <Link
            to={to}
            className={`cursor-pointer p-3 sm:p-4 bg-opacity-65 whitespace-nowrap ${
                isActive ? 'bg-zinc-800' : 'hover:bg-zinc-800'
            }`}
        >
            {children}
        </Link>
    );
};

export default HeaderLink;