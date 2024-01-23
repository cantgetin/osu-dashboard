import {configureStore, ThunkAction, Action} from '@reduxjs/toolkit';
import userCardSlice from './userCardSlice';

export const store = configureStore({
    reducer: {
        userCardSlice: userCardSlice
    },
});

export type AppDispatch = typeof store.dispatch;
export type RootState = ReturnType<typeof store.getState>;
export type AppThunk<ReturnType = void> = ThunkAction<ReturnType,
    RootState,
    unknown,
    Action<string>>;