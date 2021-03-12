import React from 'react'
import PropTypes from 'prop-types'
import PencilSVG from '../../assets/svg/pencil.svg'
import TrashSVG from '../../assets/svg/trash.svg'
import { Icon } from '../icon'
import moment from 'moment'

import './styles.scss'

export function TableRow(props) {
  const {
    handleOpenConfirmDialog,
    handleOpenEmployeeDialog,
  } = props

  const { 
    name, 
    email,
    startDate,
    team,
  } = props.data

  return (
    <>
      <tr className="tableRow">
        <td className="tableRow__item tableRow--edit">
          <Icon src={PencilSVG} handleClick={() => handleOpenEmployeeDialog(props.data)}/>
        </td>
        <td className="tableRow__item tableRow--delete">
          <Icon src={TrashSVG} handleClick={() => handleOpenConfirmDialog(props.data)}/>
        </td>
        <td className="tableRow__item">{name}</td>
        <td className="tableRow__item">{email}</td>
        <td className="tableRow__item">{moment(startDate).format('DD/MM/YYYY')}</td>
        <td className="tableRow__item">{team}</td>
      </tr>
    </>
  )
}

TableRow.propTypes = {
  data: PropTypes.object,
  handleOpenConfirmDialog: PropTypes.func,
  handleOpenEmployeeDialog: PropTypes.func
}