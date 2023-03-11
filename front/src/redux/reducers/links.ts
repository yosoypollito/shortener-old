import axios from "axios"
import type { LinkState, LinkResponse, Scope } from '@src/types'
import { createSlice } from '@reduxjs/toolkit'

import store, { RootState } from '@redux/store'
import { PayloadAction } from '@reduxjs/toolkit'

import * as utils from "@utils"

const API_URL = import.meta.env.VITE_API_URL;
const initialState:LinkState = {
  step:0,
  link:""
}

const linkSlice = createSlice({
  name:'link',
  initialState,
  reducers:{
    nextStep:state=>{ 
      state.step += 1;
    },
    prevStep:state=>{ 
      if(state.step == 0) return;
      state.step -= 1 
    },
    setLink:(state, action: PayloadAction<string>)=>{
      state.link = `${API_URL}/${action.payload}`
    }
  }
})

export const selectLink = (state:RootState) => state.link;

export const { nextStep, prevStep, setLink } = linkSlice.actions

export const create = (scope:Scope)=>store.dispatch(async(dispatch)=>{
  try{
    const { data:{ link }} = await axios.post<LinkResponse>(`${API_URL}/link`, {
      scope:scope
    })
    dispatch(nextStep())
    dispatch(setLink(link.id))
  }catch(e:any){
      utils.Error(e);
  }
})

export default linkSlice.reducer;
