import React, { useEffect } from 'react'

import { useDispatch, useSelector } from 'react-redux'
import {incrementCurrentTime, setCurrentTime} from './slice'

import YouTube from "react-youtube";

let player;
let internalTimer;

const onPlayerReady = (event) => {
    player = event.target;
}

const Video = ({video}) => {
    const dispatch = useDispatch();
    const playBackTime = useSelector((state) => state.video.playBackTime)

    const onStateChange = (event) => {
        switch(event.data) {
            case -1: // unstarted
                break;
            case 0: // ended
                clearInterval(internalTimer)
                internalTimer = null
                break;
            case 1: // playing
                clearInterval(internalTimer)
                internalTimer = setInterval(() => dispatch(incrementCurrentTime(1)), 1000)
                break;
            case 2: // paused
                clearInterval(internalTimer)
                internalTimer = null
                dispatch(setCurrentTime(player.getCurrentTime()))
                break;
            case 3: // buffering
                break;
            case 5: // video cued
                break;
            default:
        }
    }

    useEffect(() => {
        if(player === undefined) return
        if(playBackTime > 0) {
            player.seekTo(playBackTime)
        }
    }, [playBackTime])

    return(
        <div className="video-wrapper">
            <YouTube videoId={video.id}
                     onReady={onPlayerReady}
                     onStateChange={onStateChange}
            />
        </div>
    )
}

export default Video