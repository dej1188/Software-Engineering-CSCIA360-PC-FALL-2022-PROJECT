import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    notations: [
        { id: 0, timeStamp: 30, description: "This is a note for 30 seconds"},
        { id: 1, timeStamp: 60, description: "This is a note for 60 seconds"},
        { id: 2, timeStamp: 90, description: "This is a note for 90 seconds"},
        { id: 3, timeStamp: 120, description: "This is a note for 120 seconds"},
    ]
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