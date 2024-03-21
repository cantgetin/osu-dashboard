import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

interface UsersState {
    users: User[] | null
    loading: LoadingState
}

const initialState: UsersState = {
    users: null,
    loading: LoadingState.Idle,
}

export const fetchUsers = createAsyncThunk(
    'users/fetch',
    async (): Promise<{ users: User[] }> => {
        const response = await fetch(`/api/user/list`);
        const userData = await response.json();
        return {users: userData}
    }
)

const usersSlice = createSlice({
    name: 'User',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchUsers.fulfilled, (state, action) => {
            state.users = action.payload.users
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

export const {} = usersSlice.actions

export const selectUsers = (state: RootState) => state.usersSlice.users as User[]

export const selectUsersLoadingState = (state: RootState) => state.usersSlice.loading as LoadingState

export default usersSlice.reducer