import React from 'react';
import {FaExternalLinkAlt} from "react-icons/fa";
import Button from "../../ui/Button.tsx";
import {convertDateFormat} from "../../../utils/utils.ts";

interface UserProps {
    user: User
    children: React.ReactNode[]
    nameOnClick?: Function
    externalLinkOnClick: Function
}

const User = (props: UserProps) => {
    return (
        <div className="md:w-full flex bg-zinc-900 rounded-lg overflow-hidden max-h-64">
            <img
                src={props.user.avatar_url}
                className="w-16 h-16 md:w-64 md:h-64"
                alt="user avatar"
            />
            <div className="py-4 pl-4 sm:px-4 md:flex flex-col h-full md:max-w-[400px] justify-between whitespace-nowrap">
                <div>
                    <div className="flex gap-2 items-center">
                        <h1 className={`text-xl md:text-3xl ${props.nameOnClick != undefined ?
                            "hover:text-amber-200 cursor-pointer text-wrap" : null}`}
                            onClick={() => {
                                if (props.nameOnClick != undefined) props.nameOnClick!()
                            }}>
                            {props.user.username}
                        </h1>
                        <Button
                            onClick={() => props.externalLinkOnClick()}
                            className="bg-zinc-800 rounded-md p-1 h-6 hidden md:block"
                            content={<FaExternalLinkAlt className="h-3"/>}
                        />
                    </div>
                    <span className="text-sm hidden md:block text-zinc-400 px-1">
                        tracking since {convertDateFormat(props.user.tracking_since)}
                    </span>
                </div>
                <div className="hidden md:block">
                    {props.children[0]}
                </div>
            </div>
            <div className="py-4 pr-4 sm:px-4 ml-auto flex flex-col gap-2 text-right whitespace-nowrap md:flex">
                {props.children[1]}
            </div>
        </div>
    );
};

export default User;