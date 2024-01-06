function convertDateFormat(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');
    const year = String(dateObj.getUTCFullYear());

    // Creating the formatted date string
    return `${day}.${month}.${year}`;
}

function convertDataToDayMonth(inputDate: string): string {
    const dateObj = new Date(inputDate);

    // Extracting date components
    const day = String(dateObj.getUTCDate()).padStart(2, '0');
    const month = String(dateObj.getUTCMonth() + 1).padStart(2, '0');

    // Creating the formatted date string
    return `${day}.${month}`;
}

function mapUserStatsToArray(userStats: UserStats): UserStatsDataset[] {
    return Object.keys(userStats).map((timestamp) => ({
        timestamp: convertDataToDayMonth(timestamp),
        ...userStats[timestamp],
    }));
}

export {convertDateFormat, mapUserStatsToArray};

