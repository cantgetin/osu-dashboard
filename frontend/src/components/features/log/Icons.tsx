import {FaCalendarAlt, FaClock, FaExclamationTriangle, FaServer} from "react-icons/fa";

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

// Type Icon
export const TypeIcon = ({className = 'w-4 h-4'}) => (
    <FaExclamationTriangle className={`text-emerald-400 ${className}`}/>
);