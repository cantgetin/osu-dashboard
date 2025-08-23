import axios from "axios";

const CLIENT_ID = import.meta.env.VITE_OSU_API_CLIENT_ID
const REDIRECT_URI = import.meta.env.VITE_REDIRECT_URI

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

export function buildQueryParams(cmd: any): string {
    const params = new URLSearchParams();

    const validParams = ['search', 'status', 'sort', 'direction', 'page'];

    validParams.forEach(param => {
        if (cmd[param] != null && cmd[param] !== '') {
            params.append(param, cmd[param]);
        }
    });

    const queryString = params.toString();
    return queryString ? `?${queryString}` : '';
}

export async function handleOsuSiteRedirect(state: string, code: string) {
    console.log(`redirect state: ${state} local state: ${localStorage.getItem('state')}, all good`)
    if (state == localStorage.getItem('state')) {
        localStorage.setItem('code', code?.toString())
        console.log('set the code to local storage, now exchange code for token')

        await axios.post(`/api/following/create/${code}`);
    }
}

export const formatNanosToMilliseconds = (nanoseconds: number): string => {
    const milliseconds = nanoseconds / 1e6;
    return `${milliseconds.toFixed(2)}ms`;
};