import { createAsyncThunk, createSlice } from '@reduxjs/toolkit'
import { RootState } from './store';
import { LoadingState } from '../interfaces/LoadingState';

interface MapsetState {
    mapsets: Mapset[] | null
    count: number
    page: number
    loading: LoadingState
}

const initialState: MapsetState = {
    mapsets: null,
    count: 0,
    page: 0,
    loading: LoadingState.Idle,
}

interface fetchMapsetsCommand {
    search?: string
    status?: string
    sort?: boolean
    direction?: string
    forUser?: boolean
    userId?: number
    page?: number
}

export const fetchMapsets = createAsyncThunk(
    'Mapsets/fetch',
    async (cmd: fetchMapsetsCommand): Promise<{mapsets: Mapset[]}> => {
        let queryParams = `?search=${cmd.search}&status=${cmd.status}&sort=${cmd.sort}&direction=${cmd.direction}`

        if (cmd.forUser) {
            const response = await fetch(`/api/beatmapset/list_for_user/${cmd.userId}`+queryParams);
            const mapsetsData = await response.json();
            return {mapsets: mapsetsData}
        } else {
            const response = await fetch(`/api/beatmapset/list`+queryParams);
            const mapsetsData = await response.json();
            return {mapsets: mapsetsData}
        }
    }
)

const mapsetSlice = createSlice({
    name: 'Mapsets',
    initialState,
    reducers: {

    },
    extraReducers: (builder) => {
        builder.addCase(fetchMapsets.fulfilled, (state, action) => {
            state.mapsets = action.payload.mapsets
            state.loading = LoadingState.Succeeded
        })
        builder.addCase(fetchMapsets.pending, (state) => {
            state.loading = LoadingState.Pending
        })
        builder.addCase(fetchMapsets.rejected, (state) => {
            state.loading = LoadingState.Failed
        })
    },
})

export const { } = mapsetSlice.actions

export const selectMapsets = (state: RootState) => state.mapsetsSlice.mapsets as Mapset[]
export const selectMapsetsCount = (state: RootState) => state.mapsetsSlice.count
export const selectMapsetsPage = (state: RootState) => state.mapsetsSlice.page
export const selectMapsetsLoading = (state: RootState) => state.mapsetsSlice.loading

export default mapsetSlice.reducer