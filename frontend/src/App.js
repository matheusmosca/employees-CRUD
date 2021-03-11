import React from 'react'
import { HeroBackground } from './components/heroBackground'
import { EmployeeTable } from './components/employeeTable'
import './app.scss'

function App() {

  return (
    <main className="homepage">
      <HeroBackground />
      <div className="homepage__tableWrapper">
        <EmployeeTable />
      </div>
    </main>
  )
}

export default App
