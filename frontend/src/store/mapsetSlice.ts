import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

export interface MapsetState {
    mapset: Mapset | null
    loading: LoadingState
}

const initialState: MapsetState = {
    mapset: null,
    loading: LoadingState.Idle,
}

export const fetchMapset = createAsyncThunk(
    'Mapset/fetch',
    async (mapsetID: string): Promise<{ mapset: Mapset}> => {

        const response = await fetch(`/api/beatmapset/${mapsetID}`);
        const mapsetsData = await response.json();
        return {mapset: mapsetsData};
    }
);
const mapsetSlice = createSlice({
    name: 'Mapsets',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchMapset.fulfilled, (state, action) => {
            state.mapset = action.payload.mapset
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchMapset.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchMapset.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

export const {} = mapsetSlice.actions

export const selectMapsetState = (state: RootState) => state.mapsetSlice
export const selectMapset = (state: RootState) => state.mapsetSlice.mapset as Mapset
export const selectMapsetLoading = (state: RootState) => state.mapsetSlice.loading

export default mapsetSlice.reducer