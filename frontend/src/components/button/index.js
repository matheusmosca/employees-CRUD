import React from 'react'
import PropTypes from 'prop-types'

import './styles.scss'

export function Button({ text, icon, onClick }) {
  if (icon) {
    return (
      <button onClick={onClick} className="button">
        <img src={icon} alt="" /> {text}
      </button>
    )
  } 

  return <button onClick={onClick} className="button">{text}</button>
}

Button.propTypes = {
  text: PropTypes.string,
  onClick: PropTypes.func,
  icon: PropTypes.any,
}