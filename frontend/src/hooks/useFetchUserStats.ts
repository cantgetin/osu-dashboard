import { useEffect } from 'react';
import { fetchUserStats } from "../store/userStatsSlice.ts";
import { useAppDispatch } from "../store/hooks.ts";

const fetchCache = new Set<string>();

// idk why im doing this i should work on redux store instead
export function useFetchUserStatsOnce(userId: string) {
    const dispatch = useAppDispatch();

    useEffect(() => {
        if (!userId || fetchCache.has(userId)) return;

        fetchCache.add(userId);
        dispatch(fetchUserStats(userId));

        return () => {};
    }, [dispatch, userId]);
}