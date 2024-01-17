export function convertDateFormat(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');
    const year = String(dateObj.getUTCFullYear());

    // Creating the formatted date string
    return `${day}.${month}.${year}`;
}

export function convertDataToDayMonth(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');

    // Creating the formatted date string
    return `${day}.${month}`;
}

export function mapUserStatsToArray(userStats: UserStats): UserStatsDataset[] {
    return Object.keys(userStats).map((timestamp) => ({
        timestamp: timestamp,
        ...userStats[timestamp],
    }));
}

export function formatDateDiff(startDateString: string, endDateString: string): string {
    const startDate = new Date(startDateString);
    const endDate = new Date(endDateString);

    const timeDiff = Math.abs(endDate.getTime() - startDate.getTime());
    const seconds = Math.floor(timeDiff / 1000);
    const minutes = Math.floor(seconds / 60);
    const hours = Math.floor(minutes / 60);
    const days = Math.floor(hours / 24);
    const months = Math.floor(days / 30);

    if (months > 0) {
        return months === 1 ? '1 month' : `${months} months`;
    } else if (days > 0) {
        return days === 1 ? '1 day' : `${days} days`;
    } else if (hours > 0) {
        return hours === 1 ? '1 hour' : `${hours} hours`;
    } else {
        return minutes === 1 ? '1 minute' : `${minutes} minutes`;
    }
}

