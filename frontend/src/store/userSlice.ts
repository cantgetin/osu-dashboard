import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

interface UserState {
    user: User | null
    loading: LoadingState
}

const initialState: UserState = {
    user: null,
    loading: LoadingState.Idle,
}

export const fetchUser = createAsyncThunk(
    'user/fetch',
    async (userId: number): Promise<{ user: User }> => {
        const response = await fetch(`/api/user/${userId}`);
        const userData = await response.json();
        return {user: userData}
    }
)

const userSlice = createSlice({
    name: 'User',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchUser.fulfilled, (state, action) => {
            state.user = action.payload.user
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchUser.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchUser.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

// eslint-disable-next-line no-empty-pattern
export const {} = userSlice.actions

export const selectUser = (state: RootState) => state.userSlice.user as User

export const selectUserLoadingState = (state: RootState) => state.userSlice.loading as LoadingState

export default userSlice.reducer