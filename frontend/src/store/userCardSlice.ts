import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { RootState } from './store';
import { LoadingState } from '../interfaces/LoadingState';

interface UserCardState {
    userCard: UserCard | null
    loading: LoadingState
    page: number
}

const initialState: UserCardState = {
    userCard: null,
    loading: LoadingState.Idle,
    page: 0
}

interface fetchUserCardCmd {
    userId: number
    page: number
}

export const fetchUserCard = createAsyncThunk(
    'userCard/fetch',
    async (cmd : fetchUserCardCmd): Promise<{userCard: UserCard, page: number}> => {
        const response = await fetch(`/api/user_card/${cmd.userId}?page=${cmd.page}`);
        const userData = await response.json();
        return {userCard: userData, page: cmd.page}
    }
)

const userCardSlice = createSlice({
    name: 'UserCard',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(fetchUserCard.fulfilled, (state, action) => {
            state.userCard = action.payload.userCard
            state.page = action.payload.page
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchUserCard.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchUserCard.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

export const { } = userCardSlice.actions

export const selectUserCard = (state: RootState) => state.userCardSlice.userCard as UserCard
export const selectLoadingState = (state: RootState) => state.userCardSlice.loading as LoadingState
export const selectUserCardPage = (state: RootState) => state.userCardSlice.page as number

export default userCardSlice.reducer