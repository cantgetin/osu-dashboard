import React from 'react';

interface FooterLinkProps {
    href: string;
    children: React.ReactNode;
    target?: string;
    rel?: string;
}

const FooterLink: React.FC<FooterLinkProps> = (props) => (
    <a
        href={props.href}
        className="mx-2 sm:mx-3 cursor-pointer hover:text-zinc-300 whitespace-nowrap text-[16px]"
        target={props.target}
        rel={props.rel}
    >
        {props.children}
    </a>
);

FooterLink.defaultProps = {
    target: "_blank",
    rel: "noopener noreferrer"
};

export default FooterLink;