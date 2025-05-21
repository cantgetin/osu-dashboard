import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

export interface UserStatisticState {
    stats: UserStatistics | null
    loading: LoadingState
}

const initialState: UserStatisticState = {
    stats: null,
    loading: LoadingState.Idle,
}

export const fetchUserStats = createAsyncThunk(
    'UserStats/fetch',
    async (userID: string): Promise<{ stats: UserStatistics}> => {

        const response = await fetch(`/api/user/statistic/${userID}`);
        const data = await response.json();
        return {stats: data};
    }
);
const userStatsSlice = createSlice({
    name: 'UserStats',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchUserStats.fulfilled, (state, action) => {
            state.stats = action.payload.stats
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchUserStats.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchUserStats.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

export const {} = userStatsSlice.actions

export const selectUserStats = (state: RootState) => state.userStatsSlice.stats
export const selectUserStatsLoading = (state: RootState) => state.userStatsSlice.loading

export default userStatsSlice.reducer