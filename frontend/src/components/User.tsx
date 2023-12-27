import React from 'react';

interface UserProps {
    user: User
    children: React.ReactNode[]
}

const User = (props: UserProps) => {
    return (
        <div className="flex bg-zinc-900 w-full rounded-lg overflow-hidden">
            <img src={props.user.avatar_url} height={200} width={200}/>
            <div className="p-2 flex flex-col">
                <span className="text-sm text-zinc-400">Logged in as:</span>
                <h1 className="text-3xl">{props.user.username}</h1>
                {props.children[1]}
            </div>
            <div className="p-2 flex flex-col bg-zinc-900 gap-2 w-72 text-right">
                {props.children[0]}
            </div>
        </div>
    );
};

export default User;