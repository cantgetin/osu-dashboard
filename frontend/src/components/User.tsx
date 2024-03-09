import React from 'react';
import {convertDateFormat} from "../utils/utils.ts";
import Button from "./ui/Button.tsx";
import {FaExternalLinkAlt} from "react-icons/fa";

interface UserProps {
    user: User
    children: React.ReactNode[]
    nameOnClick?: Function
    externalLinkOnClick: Function
}

const User = (props: UserProps) => {
    return (
        <div className="min-w-[800px] flex bg-zinc-900 rounded-lg overflow-hidden max-h-64">
            <img src={props.user.avatar_url} className="w-64 h-64" alt="user avatar"/>
            <div className="p-4 flex flex-col gap-2 h-64 max-w-[400px] justify-between whitespace-nowrap">
                <div>
                    <div className="flex gap-2 items-center">
                        <h1 className={`text-3xl ${props.nameOnClick != undefined ? "hover:text-amber-200 cursor-pointer" : null}`}
                            onClick={() => {
                                if (props.nameOnClick != undefined) props.nameOnClick!()
                            }}>
                            {props.user.username}
                        </h1>
                        <Button
                            onClick={() => props.externalLinkOnClick()}
                            className="bg-zinc-800 rounded-md p-1 h-6"
                            content={<FaExternalLinkAlt className="h-3"/>}
                        />
                    </div>
                    <span className="text-xs text-zinc-400 px-1">
                        tracking since {convertDateFormat(props.user.tracking_since)}
                    </span>
                </div>
                <div>
                    {props.children[0]}
                </div>
            </div>
            <div className="ml-auto p-4 flex flex-col gap-2 text-right whitespace-nowrap">
                {props.children[1]}
            </div>
        </div>
    );
};

export default User;