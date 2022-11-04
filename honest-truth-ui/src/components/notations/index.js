import React from 'react'
import { useDispatch, useSelector } from 'react-redux'
import classNames from 'classnames'

import { setPlaybackTime } from "../video/slice";

const NotationItem = ({item}) => {
    const dispatch = useDispatch();
    const videoPlaybackTime = useSelector((state) => state.video.currentTime)

    const itemClasses = classNames('item', {
        active: videoPlaybackTime >= item.timeStamp
    })

    return(
        <div className={itemClasses} onClick={() => dispatch(setPlaybackTime(item.timeStamp))}>
            <span className="meta-time">{item.timeStamp}</span>
            {item.description}
        </div>
    )
}

const Index = ({notations, onItemClick}) => {

    return(
        <div className="notation-wrapper">
            <h3 className="notation-header">
                Notations
            </h3>
            {notations.map((item) => <NotationItem key={item.id} item={item} onItemClick={onItemClick} />)}
        </div>
    )
}

export default Index;