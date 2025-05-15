import React from 'react'
import Sidebar from './Sidebar';
import Content from './Content';
import { Box } from '@mui/material';

const MainLayout = () => {
  return (
    <Box sx={{ display: 'flex', background: 'black',height: '100vh' }}>
      <Sidebar />
      <Content />
    </Box>
  )
}

export default MainLayout;
