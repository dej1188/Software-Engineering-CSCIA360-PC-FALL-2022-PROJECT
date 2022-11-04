import { createSlice } from '@reduxjs/toolkit'

const initialState = {
    playBackTime: 0,
    videoState: 0,
    currentTime: 0,
}

export const videoSlice = createSlice({
    name: 'video',
    initialState,
    reducers: {
        setPlaybackTime: (state, action) => {
            state.playBackTime = action.payload
            state.currentTime = action.payload
        },
        setCurrentTime: (state, action) => {
            state.currentTime = action.payload
        },
        incrementCurrentTime: (state, action) => {
            state.currentTime += action.payload
        }
    }
})

export const { setPlaybackTime, setCurrentTime, incrementCurrentTime } = videoSlice.actions

export default videoSlice.reducer