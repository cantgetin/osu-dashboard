import Layout from "../components/ui/Layout.tsx";
import Pagination from "../components/ui/Pagination.tsx";
import List from "../components/logic/List.tsx";
import UserSearch from "../components/features/user/UserSearch.tsx";
import {FaCalendarAlt, FaClock, FaServer} from 'react-icons/fa';

interface ProgressBarProps {
    value: number;
    max: number;
    className?: string;
    colorClass?: string;
}

export const ProgressBar = ({value, max, className = '', colorClass = 'bg-blue-500'}: ProgressBarProps) => {
    const percentage = Math.min(100, (value / max) * 100);

    return (
        <div className={`w-full bg-zinc-700 rounded-full h-2 ${className}`}>
            <div
                className={`h-full rounded-full ${colorClass}`}
                style={{width: `${percentage}%`}}
            />
        </div>
    );
};

interface TagPillProps {
    text: string;
    color: 'blue' | 'green' | 'emerald' | 'amber' | 'red' | 'purple';
    className?: string;
}

const colorMap = {
    blue: 'bg-blue-900 bg-opacity-60 text-blue-300',
    green: 'bg-green-900 bg-opacity-60 text-green-300',
    emerald: 'bg-emerald-900 bg-opacity-60 text-emerald-300',
    amber: 'bg-amber-900 bg-opacity-60 text-amber-300',
    red: 'bg-red-900 bg-opacity-60 text-red-300',
    purple: 'bg-purple-900 bg-opacity-60 text-purple-300',
};

export const TagPill = ({text, color, className = ''}: TagPillProps) => {
    return (
        <span
            className={`inline-flex items-center px-2.5 py-0.5 rounded-full text-xs font-medium ${colorMap[color]} ${className}`}>
            {text}
        </span>
    );
};

// Clock Icon
export const ClockIcon = ({className = 'w-4 h-4'}) => (
    <FaClock className={`text-blue-400 ${className}`}/>
);

// Server Icon
export const ServerIcon = ({className = 'w-4 h-4'}) => (
    <FaServer className={`text-purple-400 ${className}`}/>
);

// Calendar Icon
export const CalendarIcon = ({className = 'w-4 h-4'}) => (
    <FaCalendarAlt className={`text-green-400 ${className}`}/>
);

export const formatDuration = (seconds: number): string => {
    if (seconds < 60) return `${seconds}s`;

    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = seconds % 60;

    if (minutes < 60) return `${minutes}m ${remainingSeconds}s`;

    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;

    if (hours < 24) return `${hours}h ${remainingMinutes}m`;

    const days = Math.floor(hours / 24);
    const remainingHours = hours % 24;

    return `${days}d ${remainingHours}h`;
};

export const formatDate = (dateString: string | Date): string => {
    const date = new Date(dateString);
    return new Intl.DateTimeFormat('en-US', {
        month: 'short',
        day: 'numeric',
        year: 'numeric',
        hour: '2-digit',
        minute: '2-digit',
    }).format(date);
};

const tracks = [
    {
        id: "track_001",
        name: "User Authentication",
        trackingTime: 125400, // in seconds
        apiRequests: 12450,
        lastUpdated: "2023-05-15T08:30:45Z",
        platform: "Web",
        version: "2.1",
        isActive: true,
        isDeprecated: false,
        successRate: 98.7,
        avgResponseTime: 124
    },
    {
        id: "track_002",
        name: "Payment Processing",
        trackingTime: 89200,
        apiRequests: 8560,
        lastUpdated: "2023-05-18T14:22:10Z",
        platform: "iOS",
        version: "1.8",
        isActive: true,
        isDeprecated: false,
        successRate: 99.2,
        avgResponseTime: 210
    },
    {
        id: "track_003",
        name: "Legacy Data Sync",
        trackingTime: 356800,
        apiRequests: 24500,
        lastUpdated: "2023-04-30T11:05:33Z",
        platform: "Android",
        version: "1.2",
        isActive: false,
        isDeprecated: true,
        successRate: 85.4,
        avgResponseTime: 450
    },
    {
        id: "track_004",
        name: "Notification Service",
        trackingTime: 67800,
        apiRequests: 34560,
        lastUpdated: "2023-05-20T09:15:22Z",
        platform: "Web",
        version: "3.0",
        isActive: true,
        isDeprecated: false,
        successRate: 97.1,
        avgResponseTime: 89
    },
    {
        id: "track_005",
        name: "Analytics Collector",
        trackingTime: 215600,
        apiRequests: 102400,
        lastUpdated: "2023-05-17T16:45:18Z",
        platform: "Backend",
        version: "2.5",
        isActive: true,
        isDeprecated: false,
        successRate: 99.9,
        avgResponseTime: 55
    },
    {
        id: "track_006",
        name: "Image Processing",
        trackingTime: 45600,
        apiRequests: 7800,
        lastUpdated: "2023-05-19T10:30:00Z",
        platform: "Web",
        version: "1.5",
        isActive: true,
        isDeprecated: false,
        successRate: 92.3,
        avgResponseTime: 320
    },
    {
        id: "track_007",
        name: "Experimental API",
        trackingTime: 12000,
        apiRequests: 1200,
        lastUpdated: "2023-05-10T13:20:15Z",
        platform: "Beta",
        version: "0.9",
        isActive: true,
        isDeprecated: false,
        successRate: 76.8,
        avgResponseTime: 680
    }
];

export interface Track {
    id: string;
    name: string;
    trackingTime: number; // in seconds
    apiRequests: number;
    lastUpdated: string; // ISO date string
    platform: string;
    version: string;
    isActive: boolean;
    isDeprecated: boolean;
    successRate: number; // percentage (0-100)
    avgResponseTime: number; // in milliseconds
}

const Tracks = () => {
    return (
        <Layout className="flex justify-center" title="Tracks">
            <div className="flex flex-col gap-5 p-5 bg-zinc-900 rounded-lg">
                <UserSearch update={() => {
                }}/>
                <List
                    className="w-full px-2 sm:px-0 sm:w-[1152px] grid grid-cols-1 gap-4"
                    items={tracks}
                    renderItem={(track: Track) => (
                        <div
                            key={track.id}
                            className="flex flex-col sm:flex-row gap-4 p-4 bg-zinc-800 bg-opacity-30 rounded-lg hover:bg-zinc-700 transition-colors"
                        >
                            <div className="flex-1 flex flex-col gap-2">
                                <div className="flex items-center justify-between">
                                    <h3 className="text-lg font-semibold text-white">{track.name}</h3>
                                    <span className="text-sm text-zinc-400">ID: {track.id}</span>
                                </div>

                                <div className="flex flex-wrap gap-4 text-sm">
                                    <div className="flex items-center gap-1">
                                        <ClockIcon className="w-4 h-4 text-blue-400"/>
                                        <span className="text-zinc-300">Tracking:</span>
                                        <span className="font-medium text-white">
                                        {formatDuration(track.trackingTime)}
                                    </span>
                                    </div>

                                    <div className="flex items-center gap-1">
                                        <ServerIcon className="w-4 h-4 text-purple-400"/>
                                        <span className="text-zinc-300">API Requests:</span>
                                        <span className="font-medium text-white">
                                        {track.apiRequests.toLocaleString()}
                                    </span>
                                    </div>

                                    <div className="flex items-center gap-1">
                                        <CalendarIcon className="w-4 h-4 text-green-400"/>
                                        <span className="text-zinc-300">Last Updated:</span>
                                        <span className="font-medium text-white">
                                        {formatDate(track.lastUpdated)}
                                    </span>
                                    </div>
                                </div>

                                <div className="mt-2 flex flex-wrap gap-3">
                                    <TagPill color="blue" text={track.platform}/>
                                    <TagPill color="green" text={`v${track.version}`}/>
                                    {track.isActive && <TagPill color="emerald" text="Active"/>}
                                    {track.isDeprecated && <TagPill color="amber" text="Deprecated"/>}
                                </div>
                            </div>

                            <div className="w-full sm:w-[300px] sm:min-w-[300px] flex flex-col gap-2">
                                <div className="flex justify-between text-sm">
                                    <span className="text-zinc-400">Success Rate:</span>
                                    <span className="font-medium text-white">
                                    {track.successRate}%
                                </span>
                                </div>
                                <ProgressBar value={track.successRate} max={100} className="h-2"/>

                                <div className="flex justify-between text-sm mt-2">
                                    <span className="text-zinc-400">Avg. Response Time:</span>
                                    <span className="font-medium text-white">
                                    {track.avgResponseTime}ms
                                </span>
                                </div>
                                <ProgressBar
                                    value={track.avgResponseTime}
                                    max={1000}
                                    className="h-2"
                                    colorClass="bg-yellow-500"
                                />
                            </div>
                        </div>
                    )}
                />

                <Pagination
                    pages={10}
                    currentPage={1}
                    onPageChange={() => {
                    }}
                    className="flex gap-1 md:gap-2 justify-end text-sm md:text-md"
                />
            </div>
        </Layout>
    );
};

export default Tracks;