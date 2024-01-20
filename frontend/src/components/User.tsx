import React from 'react';
import {convertDateFormat} from "../utils/utils.ts";

interface UserProps {
    user: User
    children: React.ReactNode[]
}

const User = (props: UserProps) => {
    return (
        <div className="flex bg-zinc-900 w-full rounded-lg overflow-hidden max-h-64">
            <img src={props.user.avatar_url} className="w-64 h-64" alt="user avatar"/>
            <div className="p-2 flex flex-col w-96 gap-2 h-64">
                <div>
                    <h1 className="text-3xl">{props.user.username}</h1>
                    <span className="text-xs text-zinc-400 px-1">
                        tracking since {convertDateFormat(props.user.tracking_since)}
                    </span>
                </div>
                <div>
                    {props.children[0]}
                </div>
            </div>
            <div className="p-2 flex flex-col bg-zinc-900 gap-2 text-right w-full">
                {props.children[1]}
            </div>
        </div>
    );
};

export default User;