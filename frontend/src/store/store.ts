import {configureStore, ThunkAction, Action} from '@reduxjs/toolkit';
import userCardSlice from './userCardSlice';
import userSlice from "./userSlice.ts";
import mapsetsSlice from "./mapsetsSlice.ts";
import usersSlice from "./usersSlice.ts";
import mapsetSlice from "./mapsetSlice.ts";
import systemStatsSlice from "./systemStats.ts";
import userStatsSlice from "./userStatsSlice.ts"

export const store = configureStore({
    reducer: {
        userCardSlice: userCardSlice,
        userSlice: userSlice,
        usersSlice: usersSlice,
        mapsetsSlice: mapsetsSlice,
        mapsetSlice: mapsetSlice,
        systemStatsSlice: systemStatsSlice,
        userStatsSlice: userStatsSlice,
    },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType,
    RootState,
    unknown,
    Action<string>>;