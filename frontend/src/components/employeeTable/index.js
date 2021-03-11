import React, { useEffect, useState } from 'react'
import { TableRow } from '../tableRow'
import { Button } from '../button'
import { EmployeeDialog } from '../employeeDialog'
import { ConfirmDialog } from '../confirmDialog'
import API from '../../services/api'

import Plus from '../../assets/svg/plus.svg'
import './styles.scss'

export function EmployeeTable() {
  const [openEmployeeDialog, setOpenEmployeeDialog] = useState(false)
  const [employees, setEmployees] = useState([])
  const [employeeData, setEmployeeData] = useState(null)
  const [openConfirmDialog, setOpenConfirmDialog] = useState(false)
  const [fetchAgain, setFetchAgain] = useState(true)

  const handleOpenEmployeeDialog = (data) => { 
    if (data && data._id) {
      setEmployeeData(data)
    }
    setOpenEmployeeDialog(true)
  }

  const handleCloseEmployeeDialog = (response) => { 
    setEmployeeData(null)
    setFetchAgain(!!response)
    setOpenEmployeeDialog(false)
  }

  const handleOpenConfirmDialog = (data) => {
    if (data && data._id) {
      setEmployeeData(data)
    }
    setOpenConfirmDialog(true)
  }

  const handleCloseConfirmDialog = async (response) => {
    if (response) {
      try {
        await API.delete(`/employees/${employeeData._id}`)
        setFetchAgain(true)
      } catch(err) {
        setFetchAgain(false)
        throw err
      }
    }

    setEmployeeData(null)
    setOpenConfirmDialog(false)
  }

  const fetchEmployees = async () => {
    try {
      const response = await API.get('/employees')
  
      if (response && response.data) {
        setEmployees(response.data)
      }
    } catch(err) {
      console.log(err)
      throw err
    }
  }

  useEffect(() => {
    console.log(fetchAgain)
    if (fetchAgain) {
      fetchEmployees()
    }
    return () => setFetchAgain(false)
  }, [fetchAgain])

  return (
    <div>
      <EmployeeDialog 
        open={openEmployeeDialog}
        closeDialog={handleCloseEmployeeDialog}
        data={employeeData}
      />
      <ConfirmDialog 
        open={openConfirmDialog}
        closeDialog={handleCloseConfirmDialog}
      />
      <Button onClick={handleOpenEmployeeDialog} text='Add new employee' icon={Plus}/>
      <table className="employeeTable">
        <tbody>
          <tr className="employeeTable__header">
            <th className="rowEdit"></th>
            <th className="rowDelete"></th>
            <th className="rowName">Name</th>
            <th className="rowEmail">Email</th>
            <th className="rowDate">Start Date</th>
            <th className="rowTeam">Team</th>
          </tr>
          { employees && employees.map(data => (
            <TableRow
              key={data._id}
              data={data}
              handleOpenConfirmDialog={handleOpenConfirmDialog}
              handleOpenEmployeeDialog={handleOpenEmployeeDialog}
            />
          )) }
        </tbody>
      </table>
    </div>
  )
}