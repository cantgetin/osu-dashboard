import {createSlice, PayloadAction} from "@reduxjs/toolkit";
import {RootState} from "./store";

export interface AuthUser {
    id: number;
    username: string;
    avatar_url: string;
}

export interface AuthState {
    user: AuthUser | null;
}

const initialState: AuthState = {
    user: null,
};

const authSlice = createSlice({
    name: "auth",
    initialState,
    reducers: {
        setUser(state, action: PayloadAction<AuthUser | null>) {
            state.user = action.payload;
        },
        clearUser(state) {
            state.user = null;
        },
    },
});

export const {setUser, clearUser} = authSlice.actions;

export const selectAuthUser = (state: RootState) => state.auth.user;

export default authSlice.reducer;

