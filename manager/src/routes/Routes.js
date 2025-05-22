import React from 'react'
import { Route, Routes } from 'react-router-dom'
import Dashboard from '../pages/Dashboard'
import AllEmployees from '../pages/AllEmployees'
import AllDepartments from '../pages/AllDepartments'

const Router = () => {
  return (
    <Routes>
      <Route path="dashboard" element={<Dashboard />} />
      <Route path="employees" element={<AllEmployees />} />
      <Route path="departments" element={<AllDepartments />} />
    </Routes>
  )
}

export default Router
