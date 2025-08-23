import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';
import {buildQueryParams} from "../utils/utils.ts";

export interface UsersState {
    users: User[] | null
    pages: number
    currentPage: number
    loading: LoadingState
}

const initialState: UsersState = {
    users: null,
    pages: 0,
    currentPage: 0,
    loading: LoadingState.Idle,
}

export interface fetchUsersProps {
    search?: string
    sort?: string
    direction?: string
    page?: number
}

export const fetchUsers = createAsyncThunk(
    'users/fetch',
    async (cmd: fetchUsersProps): Promise<UsersState> => {
        const queryParams = buildQueryParams(cmd)

        const response = await fetch(`/api/user/list${queryParams}`);
        const userData = await response.json();
        return {
            loading: LoadingState.Pending,
            currentPage: userData.current_page,
            users: userData.users,
            pages: userData.pages
        }
    }
)

const usersSlice = createSlice({
    name: 'User',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchUsers.fulfilled, (state, action) => {
            state.users = action.payload.users
            state.pages = action.payload.pages
            state.currentPage = action.payload.currentPage
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchUsers.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchUsers.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

// eslint-disable-next-line no-empty-pattern
export const {} = usersSlice.actions

export const selectUsers = (state: RootState) => state.usersSlice as UsersState

export const selectUsersLoadingState = (state: RootState) => state.usersSlice.loading as LoadingState

export default usersSlice.reducer