import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';
import {buildQueryParams} from "../utils/utils.ts";

export interface LogsState {
    logs: Log[] | null
    pages: number
    currentPage: number
    loading: LoadingState
}

const initialState: LogsState = {
    logs: null,
    pages: 0,
    currentPage: 0,
    loading: LoadingState.Idle,
}

export interface fetchLogsProps {
    search?: string
    sort?: string
    direction?: string
    page?: number
}

export const fetchLogs = createAsyncThunk(
    'logs/fetch',
    async (cmd: fetchLogsProps): Promise<LogsState> => {
        const queryParams = buildQueryParams(cmd)

        const response = await fetch(`/api/log/list${queryParams}`);
        const userData = await response.json();
        return {
            loading: LoadingState.Succeeded,
            currentPage: userData.current_page,
            logs: userData.logs,
            pages: userData.pages
        }
    }
)

const logsSlice = createSlice({
    name: 'Logs',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchLogs.fulfilled, (state, action) => {
            state.logs = action.payload.logs
            state.pages = action.payload.pages
            state.currentPage = action.payload.currentPage
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchLogs.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchLogs.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

// eslint-disable-next-line no-empty-pattern
export const {} = logsSlice.actions

export const selectLogs = (state: RootState) => state.logsSlice as LogsState

export const selectLogsLoadingState = (state: RootState) => state.usersSlice.loading as LoadingState

export default logsSlice.reducer