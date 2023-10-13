import { configureStore } from "@reduxjs/toolkit";

//Reducers
import linkSlice from "@redux/reducers/links"

const store = configureStore({
  reducer:{
    link:linkSlice
  }
});

export type RootState = ReturnType<typeof store.getState>

export type AppDispatch = typeof store.dispatch

export default store;
