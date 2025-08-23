export function getRemainingPendingTime(expirationTimeStr: string): string {
    const expirationTime = new Date(expirationTimeStr);
    const expirationTime28DaysLater = new Date(expirationTime.getTime() + (28 * 24 * 60 * 60 * 1000));
    const currentTime = new Date();
    const timeDifference = expirationTime28DaysLater.getTime() - currentTime.getTime();

    if (timeDifference <= 0) {
        return "pending";
    }

    const days = Math.floor(timeDifference / (1000 * 60 * 60 * 24));
    const hours = Math.floor((timeDifference % (1000 * 60 * 60 * 24)) / (1000 * 60 * 60));
    const minutes = Math.floor((timeDifference % (1000 * 60 * 60)) / (1000 * 60));

    const result = `pending for ${days}d ${hours}h ${minutes}m`;

    return result;
}

export function convertDataToDayMonth(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');

    // Creating the formatted date string
    return `${day}.${month}`;
}

export function convertDateFormat(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');
    const year = String(dateObj.getUTCFullYear());

    // Creating the formatted date string
    return `${day}.${month}.${year}`;
}

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

// Convert nanoseconds to seconds for duration formatting
export const formatDuration = (nanoseconds: number): string => {
    const seconds = nanoseconds / 1e9;
    if (seconds < 60) return `${seconds.toFixed(1)}s`;

    const minutes = Math.floor(seconds / 60);
    const remainingSeconds = Math.floor(seconds % 60);

    if (minutes < 60) return `${minutes}m ${remainingSeconds}s`;

    const hours = Math.floor(minutes / 60);
    const remainingMinutes = minutes % 60;

    if (hours < 24) return `${hours}h ${remainingMinutes}m`;

    const days = Math.floor(hours / 24);
    const remainingHours = hours % 24;

    return `${days}d ${remainingHours}h`;
};