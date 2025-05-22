import React, { use, useState } from 'react'
import { Box, Link, Typography } from '@mui/material'
import logo from '../assets/images/logo.png'
import DashboardIcon from '@mui/icons-material/Dashboard'
import GroupIcon from '@mui/icons-material/Group'
import AutoModeIcon from '@mui/icons-material/AutoMode'
import EventAvailableIcon from '@mui/icons-material/EventAvailable'
import SettingsIcon from '@mui/icons-material/Settings'
import { NavLink, useLocation } from 'react-router-dom'

const Sidebar = () => {
  const [sidebarActive, setSidebarActive] = useState('dashboard')
  const location = useLocation()
  console.log(' location:', location.pathname.startsWith('/dashboard'))

  const sidebarList = [
    { id: 'dashboard', title: 'Dashboard', icon: DashboardIcon },
    { id: 'employees', title: 'All Employees', icon: GroupIcon },
    { id: 'departments', title: 'All Departments', icon: AutoModeIcon },
    { id: 'attendance', title: 'Attendance', icon: EventAvailableIcon },
    { id: 'settings', title: 'Settings', icon: SettingsIcon },
  ]

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
          <Box
            sx={{ display: 'flex', justifyContent: 'center', alignItems: 'center', paddingBottom: '30px', gap: '10px' }}
          >
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
            <Typography variant="h5">HRMS</Typography>
          </Box>

          {/* navbar */}
          <Box>
            {sidebarList.map((item, index) => (
              <NavItem
                key={index}
                id={item.id}
                title={item.title}
                icon={item.icon}
                active={location.pathname.startsWith(`/${item.id}`)}
              />
            ))}
          </Box>
        </Box>
      </Box>
    </Box>
  )
}

const NavItem = ({ title, icon, active, id }) => {
  const NavIcon = icon

  return (
    <NavLink
      to={id}
      style={{
        display: 'flex',
        gap: '10px',
        padding: '10px',
        borderLeft: active ? '4px solid #7152F3' : '4px solid transparent',
        cursor: 'pointer',
        background: active ? 'rgba(83, 83, 83, 0.3)' : 'transparent',
        borderRadius: '0 8px 8px 0',
        textDecoration: 'none',
      }}
    >
      <NavIcon sx={{ color: active ? '#7152F3' : 'black' }} />
      <Box sx={{ color: active ? '#7152F3' : 'black', fontWeight: '500' }}>{title}</Box>
    </NavLink>
  )
}

export default Sidebar
