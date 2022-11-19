import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    notations: []
}

export const notationSlice = createSlice({
    name: 'notation',
    initialState,
    reducers: {
        addNotation: (state, action) => {
            state.notations = [ ...state.notations, action.payload ]
        },
    }
})

export const { addNotation }  = notationSlice.actions

export default notationSlice.reducer