import {createAsyncThunk, createSlice} from '@reduxjs/toolkit'
import {RootState} from './store';
import {LoadingState} from '../interfaces/LoadingState';

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
        let queryParams = '?';

        if (cmd.search != null) {
            queryParams += `&search=${cmd.search}`;
        }
        if (cmd.status != null) {
            queryParams += `&status=${cmd.status}`;
        }
        if (cmd.sort != null) {
            queryParams += `&sort=${cmd.sort}`;
        }
        if (cmd.direction != null) {
            queryParams += `&direction=${cmd.direction}`;
        }
        if (cmd.page != null) {
            queryParams += `&page=${cmd.page}`;
        }
        if (queryParams === '?') {
            queryParams = '';
        }

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

export const {} = mapsetsSlice.actions

export const selectMapsetsState = (state: RootState) => state.mapsetsSlice
export const selectMapsets = (state: RootState) => state.mapsetsSlice.mapsets as Mapset[]
export const selectMapsetsCurrentPage = (state: RootState) => state.mapsetsSlice.currentPage
export const selectMapsetsPages = (state: RootState) => state.mapsetsSlice.pages
export const selectMapsetsLoading = (state: RootState) => state.mapsetsSlice.loading

export default mapsetsSlice.reducer