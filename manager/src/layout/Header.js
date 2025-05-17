import { Avatar, Box, Button, InputAdornment, styled, TextField, Typography } from '@mui/material'
import React from 'react'
import SearchIcon from '@mui/icons-material/Search';
import ExpandMoreIcon from '@mui/icons-material/ExpandMore';

const CssTextField = styled(TextField)({
  '& .MuiOutlinedInput-root': {
    paddingTop: '5px',
    paddingBottom: '5px',
    borderRadius: '10px'
  },
});


const Header = () => {
  return (
    <Box sx={{ display: 'flex', justifyContent: 'space-between', alignItems: 'center' }}>
      <Box>
        <Typography variant='h6'>All Employees</Typography>
        <Typography sx={{color: 'gray'}}>All Employees Information</Typography>
      </Box>

      <Box sx={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
        <Box>
          <CssTextField
            size="small"
            placeholder="Search"
            id="input-with-icon-textfield"
            slotProps={{
              
              input: {
                startAdornment: (
                  <InputAdornment position="start" sx={{ height: '40px', borderRadius: '50px' }}>
                    <SearchIcon />
                  </InputAdornment>
                ),
              },
            }}
          />
        </Box>

        <Box sx={{border: '1px solid gray', borderRadius: '10px', padding: '5px'}}>
          <Box sx={{ display: 'flex', alignItems: 'center', gap: '10px' }}>
            <Avatar variant="rounded" />
            <Box>
              <Typography variant='h7'>Robert Allen</Typography>
              <Typography sx={{color: 'gray', fontSize: '12px'}}>HR Manager</Typography>
            </Box>
            <ExpandMoreIcon />
          </Box>
        </Box>
      </Box>
    </Box>
  )
}

export default Header
