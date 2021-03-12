import React, { useState, useEffect } from 'react'
import swal from 'sweetalert'
import PropTypes from 'prop-types'
import { Button } from '../button'
import Checked from '../../assets/svg/checked.svg'
import CloseSVG from '../../assets/svg/close.svg'
import { Icon } from '../icon'
import DateFnsUtils from '@date-io/moment'
import { TextField } from '@material-ui/core'
import { makeStyles } from '@material-ui/core/styles'
import InputLabel from '@material-ui/core/InputLabel'
import FormControl from '@material-ui/core/FormControl'
import Select from '@material-ui/core/Select'
import MenuItem from '@material-ui/core/MenuItem'

import API from '../../services/api'

import { 
  KeyboardDatePicker, 
  MuiPickersUtilsProvider
} from '@material-ui/pickers'

import './styles.scss'

const useStyles = makeStyles((theme) => ({
  formControl: {
    margin: theme.spacing(1),
    minWidth: 120,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
}))

const CloseIconStyle = {
  width: '16px',
  height: '16px',
  position: 'absolute',
  top: '18px',
  right: '18px'
}

export function EmployeeDialog({ open, closeDialog, data }) {
  const [gender, setGender] = useState('')
  const [email, setEmail] = useState('')
  const [team, setTeam] = useState('')
  const [name, setName] = useState(data ? data.name : '')
  const [cpf, setCpf] = useState('')
  const [birthDate, setBirthDate] = useState(new Date())
  const [startDate, setStartDate] = useState(new Date())

  const classes = useStyles()

  useEffect(() => {
    if (data && data.id) {
      setGender(data.gender ? data.gender : '')
      setEmail(data.email ? data.email : '')
      setTeam(data.team ? data.team : '')
      setName(data.name ? data.name : '')
      setCpf(data.cpf ? data.cpf : '')
      setBirthDate(data.birthDate ? data.birthDate : new Date())
      setStartDate(data.startDate ? data.startDate : new Date())
    } else {
      setGender('')
      setEmail('')
      setTeam('')
      setName('')
      setCpf('')
      setBirthDate(new Date())
      setStartDate(new Date())
    }
  }, [data])

  const handleBirthDateChange = birthDate => setBirthDate(birthDate)
  const handleStartDateChange = startDate => setStartDate(startDate)
  const handleNameChange = event => setName(event.target.value)
  const handleEmailChange = event => setEmail(event.target.value)
  const handleCpfChange = event => setCpf(event.target.value)
  const handleTeamChange = event => setTeam(event.target.value)
  const handleGenderChange = event => setGender(event.target.value)

  const handleSubmit = async (e) => {
    e.preventDefault()

    const payload = {
      name: name,
      gender,
      email,
      cpf,
      team,
      startDate,
      birthDate,
    }

    try {
      if (data && data.id) {
        const response = await API.put(`/employees/${data.id}`, payload)
        if (response) {
          closeDialog(true)
        }
      } else {
        const response = await API.post('/employees', payload)
        if (response) {
          closeDialog(true)
        }
      }
    } catch(err) {
      const message = err.response.data.error 
      if (message) {
        swal({ title: message, icon: 'warning'})
      }
    }
  }

  if (open) {
    return (
      <div className="employeeDialog">
        <div className="employeeDialog__content">
          <Icon 
            src={CloseSVG}
            style={CloseIconStyle}
            handleClick={() => closeDialog(false)}
          />
          <h2 className="employeeDialog__content__title">{ data && data.id ? 'Update employee' : 'New employee' }</h2>
          <form onSubmit={ handleSubmit } className="employeeDialog__content__form" action="">
            <TextField
              required
              onChange={handleNameChange}
              value={name}
              // error={!name}
              name="name"
              autoFocus
              variant="outlined"
              size="small"
              margin="dense"
              id="name"
              label="Name"
              fullWidth
              style={{ marginBottom: '6px' }}
            />
            <TextField
              required
              onChange={handleEmailChange}
              value={email}
              type="email"
              name="email"
              autoFocus
              variant="outlined"
              size="small"
              margin="dense"
              id="email"
              label="Email"
              fullWidth
              style={{ marginBottom: '6px' }}
            />
            <TextField
              required
              onChange={handleCpfChange}
              name="cpf"
              autoFocus
              value={cpf}
              variant="outlined"
              size="small"
              margin="dense"
              id="cpf"
              label="CPF"
              fullWidth
              style={{ marginBottom: '6px' }}
            />
            <FormControl fullWidth variant="outlined" size='small' className={classes.formControl}>
              <InputLabel id="gender-select-label">Gender</InputLabel>
              <Select
                required
                labelId="gender-select-label"
                id="gender-select"
                value={gender}
                onChange={handleGenderChange}
              >
                <MenuItem value={'Male'}>Male</MenuItem>
                <MenuItem value={'Female'}>Female</MenuItem>
                <MenuItem value={'Other'}>Other</MenuItem>
              </Select>
            </FormControl>
            <FormControl fullWidth variant="outlined" size='small' className={classes.formControl}>
              <InputLabel id="team-select-label">Team</InputLabel>
              <Select
                labelId="team-select-label"
                id="team-select"
                value={team}
                onChange={handleTeamChange}
              >
                <MenuItem value="">None</MenuItem>
                <MenuItem value='Mobile'>Mobile</MenuItem>
                <MenuItem value='Frontend'>Frontend</MenuItem>
                <MenuItem value='Backend'>Backend</MenuItem>
              </Select>
            </FormControl>
            <div className="wrapper">
              <MuiPickersUtilsProvider utils={DateFnsUtils}>
                <KeyboardDatePicker
                  required
                  name="birthDate"
                  style={{ maxWidth: '140px', marginRight: '18px', marginTop: 0, marginBottom: 0 }}
                  disableToolbar
                  variant="inline"
                  format="MM/DD/YYYY"
                  maxDate={new Date()}
                  margin="normal"
                  id="birthdate-picker-inline"
                  label="Birth date"
                  value={birthDate}
                  onChange={handleBirthDateChange}
                  KeyboardButtonProps={{
                    'aria-label': 'change date',
                  }}
                />
                <KeyboardDatePicker
                  required
                  maxDate={new Date()}
                  name="startDate"
                  format="MM/YYYY"
                  style={{ maxWidth: '140px', marginTop: 0, marginBottom: 0 }}
                  views={['year', 'month']}
                  variant="inline"
                  id="startdate-picker-inline"
                  label="Start date"
                  value={startDate}
                  onChange={handleStartDateChange}
                  KeyboardButtonProps={{
                    'aria-label': 'change date',
                  }}
                />
              </MuiPickersUtilsProvider>
              <Button type='submit' text="Save" icon={Checked} />
            </div>
          </form>
        </div>
      </div>
    )
  } else {
    return <></>
  }
}

EmployeeDialog.propTypes = {
  open: PropTypes.bool,
  closeDialog: PropTypes.func,
  data: PropTypes.object,
}