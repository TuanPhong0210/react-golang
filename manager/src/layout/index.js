import React from 'react'
import Sidebar from './Sidebar'
import Content from './Content'
import { Box } from '@mui/material'
import Header from './Header'

const MainLayout = () => {
  return (
    <Box sx={{ display: 'flex', height: '100vh' }}>
      <Sidebar />
      <Box sx={{ padding: '20px', width: '100%' }}>
        <Header />
        <Content />
      </Box>
    </Box>
  )
}

export default MainLayout
