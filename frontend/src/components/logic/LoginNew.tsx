import React, {useState} from 'react';
import {
    DropdownMenu,
    DropdownMenuContent,
    DropdownMenuItem,
    DropdownMenuLabel,
    DropdownMenuSeparator,
    DropdownMenuTrigger,
} from "@/components/ui/dropdown-menu";
import {Avatar, AvatarFallback, AvatarImage} from "@/components/ui/avatar";
import {Button} from "@/components/ui/button";
import {ChevronDown, LogOut, User} from "lucide-react";
import {cn} from "@/lib/utils";
import {useNavigate} from "react-router-dom";

interface UserData {
    id: string; // Added user ID
    name: string;
    email: string;
    avatar?: string;
    nickname?: string;
}

interface UserLoginProps {
    user?: UserData | null;
    onLogin?: () => void;
    onLogout?: () => void;
    onProfile?: () => void;
    onFeedback?: () => void;
    showNickname?: boolean;
    nicknamePosition?: 'left' | 'right';
}

const UserLogin: React.FC<UserLoginProps> = ({
                                                 user,
                                                 onLogin,
                                                 onLogout,
                                                 showNickname = true,
                                                 nicknamePosition = 'right'
                                             }) => {
    const [isOpen, setIsOpen] = useState(false);
    const navigate = useNavigate();

    const handleProfileClick = () => {
        navigate(`/user/${user!.id!}`);
    };

    if (!user) {
        return (
            <Button onClick={onLogin} variant="default" size="sm"
                    className="px-4 py-4 text-md text-white bg-green-800 hover:bg-green-900">
                Sign up
            </Button>
        );
    }

    const getInitials = (name: string) => {
        return name
            .split(' ')
            .map(part => part[0])
            .join('')
            .toUpperCase()
            .slice(0, 2);
    };

    const displayName = user.nickname || user.name;

    return (
        <div className="flex items-center bg-transparent">
            <DropdownMenu modal={false} open={isOpen} onOpenChange={setIsOpen}>
                <DropdownMenuTrigger asChild>
                    <Button
                        variant="ghost"
                        className={cn(
                            "bg-transparent relative rounded-md flex items-center gap-2 hover:bg-zinc-800 text-zinc-100",
                            showNickname ? "px-3 py-2" : "h-10 w-10 p-0"
                        )}
                        aria-label="User menu"
                    >
                        {showNickname && nicknamePosition === 'left' && (
                            <div
                                className="text-[16px] sm:text-[16px] font-medium text-zinc-200 truncate max-w-[120px]">
                                {displayName}
                            </div>
                        )}

                        <div className="flex items-center gap-2">
                            <ChevronDown className={cn(
                                "h-4 w-4 transition-transform text-zinc-400",
                                isOpen ? "rotate-180" : ""
                            )}/>
                            <Avatar className="h-8 w-8 rounded-md">
                                <AvatarFallback className="bg-zinc-700 text-zinc-200 rounded-md">
                                    {getInitials(user.name)}
                                </AvatarFallback>
                                <AvatarImage
                                    src={user.avatar}
                                    alt={user.name}
                                    className="bg-zinc-800 rounded-md"
                                />
                            </Avatar>
                        </div>

                        {showNickname && nicknamePosition === 'right' && (
                            <div
                                className="text-[16px] sm:text-[16px] font-medium text-zinc-200 truncate max-w-[120px]">
                                {displayName}
                            </div>
                        )}
                    </Button>
                </DropdownMenuTrigger>

                <DropdownMenuContent
                    className="w-56 bg-zinc-900 border-zinc-800 text-zinc-200"
                    align="end"
                    forceMount
                    onCloseAutoFocus={(e) => e.preventDefault()}
                >
                    <DropdownMenuLabel className="font-normal bg-zinc-900">
                        <div className="flex flex-col space-y-1">
                            <p className="text-md font-medium leading-none text-zinc-100">{user.name}</p>
                            <p className="text-xs leading-none text-zinc-400">
                                {user.email}
                            </p>
                        </div>
                    </DropdownMenuLabel>

                    <DropdownMenuSeparator className="bg-zinc-700"/>

                    <DropdownMenuItem
                        onClick={handleProfileClick}
                        className="text-sm focus:bg-zinc-800 focus:text-zinc-100 cursor-pointer"
                    >
                        <User className="mr-2 h-4 w-4 text-zinc-400"/>
                        <span>Profile</span>
                    </DropdownMenuItem>

                    <DropdownMenuItem
                        onClick={onLogout}
                        className="text-sm text-red-400 focus:text-red-300 focus:bg-zinc-800 cursor-pointer"
                    >
                        <LogOut className="mr-2 h-4 w-4"/>
                        <span>Logout</span>
                    </DropdownMenuItem>
                </DropdownMenuContent>
            </DropdownMenu>
        </div>
    );
};

export default UserLogin;