import React, { useState } from 'react'
import { Box, Link, Typography } from '@mui/material'
import logo from '../assets/images/logo.png'
import DashboardIcon from '@mui/icons-material/Dashboard'

const Sidebar = () => {
  const [sidebarActive, setSidebarActive] = useState('employees')
  console.log(' sidebarActive:', sidebarActive)

  return (
    <Box sx={{ padding: '20px' }}>
      <Box
        sx={{
          height: '100%',
          display: 'block',
          flexDirection: 'column',
          justifyContent: 'space-between',
          width: 240,
          background: 'rgba(162, 161, 168, 0.3)',
          borderRadius: '20px',
        }}
      >
        <Box sx={{ padding: '30px' }}>
          {/* logo */}
          <Box sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', paddingBottom: '30px' }}>
            <Box
              sx={{
                display: 'flex',
                justifyContent: 'center',
                alignItems: 'center',
                height: '24px',
                width: '24px',
                padding: '10px',
                borderRadius: '50px',
                background: '#8D75F5',
              }}
            >
              <img src={logo} alt="logo" />
            </Box>
            <Typography variant="h5" sx={{ color: 'white' }}>
              HRMS
            </Typography>
          </Box>

          {/* navbar */}
          <Box>
            <NavItem
              title="Dashboard"
              icon={DashboardIcon}
              active={sidebarActive === 'dashboard'}
              onClick={() => {
                setSidebarActive('dashboard')
              }}
            />

            <NavItem
              title="All Employees"
              icon={DashboardIcon}
              active={sidebarActive === 'employees'}
              onClick={() => {
                setSidebarActive('employees')
              }}
            />
          </Box>
        </Box>
      </Box>
    </Box>
  )
}

const NavItem = ({ title, icon, active, onClick }) => {
  const NavIcon = icon

  return (
    <Box
      sx={{
        display: 'flex',
        gap: '10px',
        padding: '10px',
        borderLeft: active ? '4px solid #7152F3' : '4px solid transparent',
        cursor: 'pointer',
        background: active ? 'rgba(83, 83, 83, 0.3)' : 'transparent',
        borderRadius: '0 8px 8px 0',
      }}
      onClick={onClick}
    >
      <NavIcon sx={{ color: active ? '#7152F3' : 'white' }} />
      <Link href="#" sx={{ color: active ? '#7152F3' : 'white', textDecoration: 'none', fontWeight: '500' }}>
        {title}
      </Link>
    </Box>
  )
}

export default Sidebar
