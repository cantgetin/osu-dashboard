import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

export interface SystemStatsState {
    stats: SystemStats | null
    loading: LoadingState
}

const initialState: SystemStatsState = {
    stats: null,
    loading: LoadingState.Idle,
}

export const fetchSystemStats = createAsyncThunk(
    'SystemStats/fetch',
    async (): Promise<{ stats: SystemStats }> => {
        const response = await fetch(`api/system/statistic`);
        const data = await response.json();
        return {stats: data};
    }
);

const systemStatsSlice = createSlice({
    name: 'SystemStats',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchSystemStats.fulfilled, (state, action) => {
            state.stats = action.payload.stats
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchSystemStats.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchSystemStats.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

// eslint-disable-next-line no-empty-pattern
export const {} = systemStatsSlice.actions

export const selectSystemStatsState = (state: RootState) => state.systemStatsSlice
export const selectSystemStats = (state: RootState) => state.systemStatsSlice.stats as SystemStats
export const selectSystemStatsLoading = (state: RootState) => state.systemStatsSlice.loading

export default systemStatsSlice.reducer