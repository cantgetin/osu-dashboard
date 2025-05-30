import axios from "axios";

const CLIENT_ID = import.meta.env.VITE_OSU_API_CLIENT_ID
const REDIRECT_URI = import.meta.env.VITE_REDIRECT_URI

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

export function mapMapsetStatsToArray(mapsetStats: MapsetStats): MapsetStatsDataset[] {
    return Object.keys(mapsetStats).map((timestamp) => ({
        timestamp: timestamp,
        ...mapsetStats[timestamp],
    }));
}

export function extractUserMapsCountFromStats(userStats: UserStats): number {
    const arr = mapUserStatsToArray(userStats);
    if (arr.length === 0) {
        return 0
    }

    const lastElement = arr[arr.length - 1];
    return lastElement.map_count;
}

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

function generateRandomString() {
    return "x".repeat(5)
        .replace(/./g, _ =>
            "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"[Math.floor(Math.random() * 62)]);
}

export function redirectToAuthorize() {
    const url = new URL(
        "https://osu.ppy.sh/oauth/authorize"
    );

    const randomString: string = generateRandomString()
    localStorage.setItem('state', randomString)

    const params: any = {
        "client_id": CLIENT_ID,
        "redirect_uri": REDIRECT_URI,
        "response_type": "code",
        "scope": "public identify",
        "state": randomString,
    };
    Object.keys(params)
        .forEach(key => url.searchParams.append(key, params[key]));

    window.location.href = url.toString();
}

export async function handleOsuSiteRedirect(state: string, code: string) {
    console.log(`redirect state: ${state} local state: ${localStorage.getItem('state')}, all good`)
    if (state == localStorage.getItem('state')) {
        localStorage.setItem('code', code?.toString())
        console.log('set the code to local storage, now exchange code for token')

        await axios.post(`/api/following/create/${code}`);
    }
}