import React, { useState } from 'react'

import { useDispatch, useSelector } from 'react-redux'
import classNames from 'classnames'

import { setPlaybackTime } from "../video/slice";
import { addNotation } from "./slice";

const AddNotationForm = () => {
    const dispatch = useDispatch();

    const formInitialState = {
        timeStamp: 0,
        description: ''
    }

    const [notationValues, setNotationValues] = useState(formInitialState)

    const handleInputChange = (event) => {
        event.persist()
        setNotationValues((values) => ({
            ...values,
            [event.target.name]: event.target.value
        }))
    }

    const handleSubmit = (event) => {
        event.preventDefault()

        const newNotation = {
            id: Math.floor(Math.random() * 1000),
            timeStamp: parseInt(notationValues.timeStamp),
            description: notationValues.description
        }

        dispatch(addNotation(newNotation))
        clearForm()
    }

    const clearForm = () => {
        setNotationValues(() => formInitialState)
    }

    return(
        <div className="add-notation-wrapper">
            Time: <input type="text" value={notationValues.timeStamp} name="timeStamp" className="notation-time" onChange={handleInputChange} />
            <br />
            <span>Note:</span><br />
            <textarea name="description" className="notation-description" onChange={handleInputChange} value={notationValues.description} />
            <input type="submit" name="notation-add" onClick={handleSubmit} />
        </div>
    )
}

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

    const sortedNotations = notations.slice().sort((a, b) => {
        return ((a.timeStamp < b.timeStamp) ? -1 : ((a.timeStamp > b.timeStamp) ? 1 : 0))
    })

    return(
        <div className="notation-wrapper">
            <h3 className="notation-header">
                Notes
            </h3>
            {sortedNotations.length === 0
                ? <p>No Notes</p>
                : sortedNotations.map((item) => <NotationItem key={item.id} item={item} onItemClick={onItemClick} />)}
            <AddNotationForm />
        </div>
    )
}

export default Index;