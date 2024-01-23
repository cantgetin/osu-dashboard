import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { RootState } from './store';
import { LoadingState } from '../interfaces/LoadingState';

interface UserCardState {
    userCard: UserCard | null
    lastTimeFetched: string | null
    loading: LoadingState
}

const initialState: UserCardState = {
    userCard: null,
    lastTimeFetched: null,
    loading: LoadingState.Idle
}

export const fetchUserCard = createAsyncThunk(
    'userCard/fetch',
    async (userId: number): Promise<{
        userCard: UserCard,
        lastTimeFetched: string
    }> => {
        const response = await fetch(`/api/user_card/${userId}`);
        const userData = await response.json();

        return {
            userCard: JSON.parse(JSON.stringify(userData)) as UserCard,
            lastTimeFetched: Date.now().toString()
        }
    }
)

const userCardSlice = createSlice({
    name: 'UserCard',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(fetchUserCard.fulfilled, (state, action) => {
            state.lastTimeFetched = action.payload.lastTimeFetched
            state.userCard = action.payload.userCard
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

export const selectUserCard = (state: RootState) => state.userCardSlice.userCard

export default userCardSlice.reducer