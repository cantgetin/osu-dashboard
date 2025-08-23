import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';
import {buildQueryParams} from "../utils/utils.ts";

export interface MapsetsState {
    mapsets: Mapset[] | null
    pages: number
    currentPage: number
    loading: LoadingState
}

const initialState: MapsetsState = {
    mapsets: null,
    pages: 0,
    currentPage: 0,
    loading: LoadingState.Idle,
}

export interface fetchMapsetsProps {
    search?: string
    status?: string
    sort?: string
    direction?: string
    forUser?: boolean
    userId?: number
    page?: number
}

export const fetchMapsets = createAsyncThunk(
    'Mapsets/fetch',
    async (cmd: fetchMapsetsProps): Promise<{ mapsets: Mapset[], pages: number, currentPage: number }> => {
        const queryParams = buildQueryParams(cmd)

        if (cmd.forUser) {
            const response = await fetch(`/api/beatmapset/list_for_user/${cmd.userId}${queryParams}`);
            const mapsetsData = await response.json();
            return {
                mapsets: mapsetsData.mapsets,
                pages: mapsetsData.pages,
                currentPage: mapsetsData.current_page,
            };
        } else {
            const response = await fetch(`/api/beatmapset/list${queryParams}`);
            const mapsetsData = await response.json();
            return {
                mapsets: mapsetsData.mapsets,
                pages: mapsetsData.pages,
                currentPage: mapsetsData.current_page,
            };
        }
    }
);
const mapsetsSlice = createSlice({
    name: 'Mapsets',
    initialState,
    reducers: {},
    extraReducers: (builder) => {
        builder.addCase(fetchMapsets.fulfilled, (state, action) => {
            state.mapsets = action.payload.mapsets
            state.pages = action.payload.pages
            state.currentPage = action.payload.currentPage
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

// eslint-disable-next-line no-empty-pattern
export const {} = mapsetsSlice.actions

export const selectMapsetsState = (state: RootState) => state.mapsetsSlice
export const selectMapsets = (state: RootState) => state.mapsetsSlice.mapsets as Mapset[]
export const selectMapsetsCurrentPage = (state: RootState) => state.mapsetsSlice.currentPage
export const selectMapsetsPages = (state: RootState) => state.mapsetsSlice.pages
export const selectMapsetsLoading = (state: RootState) => state.mapsetsSlice.loading

export default mapsetsSlice.reducer