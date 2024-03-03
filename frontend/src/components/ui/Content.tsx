import React from "react";

interface UserProps {
    children: React.ReactNode
    className?: string
}

const Content = (props: UserProps) => {
    return (
        <div className={`pt-10 pb-10 ${props.className}`}>
            {props.children}
        </div>
    );
};

export default Content;