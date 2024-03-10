import {configureStore, ThunkAction, Action} from '@reduxjs/toolkit';
import userCardSlice from './userCardSlice';
import userSlice from "./userSlice.ts";
import mapsetsSlice from "./mapsetsSlice.ts";

export const store = configureStore({
    reducer: {
        userCardSlice: userCardSlice,
        userSlice: userSlice,
        mapsetsSlice: mapsetsSlice
    },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType,
    RootState,
    unknown,
    Action<string>>;