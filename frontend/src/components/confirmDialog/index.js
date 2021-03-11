import React from 'react'
import PropTypes from 'prop-types'
import { Button } from '../button'

import './styles.scss'

export function ConfirmDialog({ open, closeDialog, title, text }) {

  const handleConfirm = () => closeDialog(true)
  const handleDecline = () => closeDialog(false)

  if (open) {
    return (
      <div className='confirmDialog'>
        <div className='confirmDialog__content'>
          <h2 className='confirmDialog__content__title'>{title ? title : 'Are your sure?'}</h2>
          <p>{text ? text : 'This action is irreversible'}</p>
          <div className="wrapper">
            <Button onClick={handleConfirm} text='Yes'/>
            <Button onClick={handleDecline} text='Cancel'/>
          </div>
        </div>
      </div>
    )
  } else {
    return <></>
  }
}

ConfirmDialog.propTypes = {
  open: PropTypes.bool,
  closeDialog: PropTypes.func,
  title: PropTypes.string,
  text: PropTypes.string,
}