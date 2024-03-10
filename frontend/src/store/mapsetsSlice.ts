import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { RootState } from './store';
import { LoadingState } from '../interfaces/LoadingState';

interface MapsetState {
    mapsets: Mapset[] | null
    loading: LoadingState
}

const initialState: MapsetState = {
    mapsets: null,
    loading: LoadingState.Idle,
}

export const fetchMapsetsForUser = createAsyncThunk(
    'Mapsets/fetch',
    async (userId: number): Promise<{mapsets: Mapset[]}> => {
        const response = await fetch(`/api/beatmapset/list_for_user/${userId}`);
        const mapsetsData = await response.json();
        return {mapsets: mapsetsData}
    }
)

const mapsetSlice = createSlice({
    name: 'Mapsets',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(fetchMapsetsForUser.fulfilled, (state, action) => {
            state.mapsets = action.payload.mapsets
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchMapsetsForUser.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchMapsetsForUser.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

export const { } = mapsetSlice.actions

export const selectMapset = (state: RootState) => state.mapsetsSlice.mapsets as Mapset[]
export const selectMapsetLoading = (state: RootState) => state.mapsetsSlice.loading

export default mapsetSlice.reducer