import React from 'react'
import PropTypes from 'prop-types'
import './styles.scss'

export function Icon({ src, handleClick, style }) {
  return (
    <img className="icon" onClick={handleClick} style={style} src={src} alt="icon" />
  )
}

Icon.propTypes = {
  src: PropTypes.any,
  handleClick: PropTypes.func,
  style: PropTypes.object,
}