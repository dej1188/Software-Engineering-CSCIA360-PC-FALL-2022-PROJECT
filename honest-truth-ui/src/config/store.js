import {configureStore} from '@reduxjs/toolkit'
import videoReducer from '../components/video/slice'
import notationsReducer from '../components/notations/slice'

export const store = configureStore({
    reducer: {
        video: videoReducer,
        notations: notationsReducer
    }
})

export default store