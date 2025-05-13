import React from 'react'
import { NavLink } from 'react-router-dom'

const Home = () => {
  return (
    <div>
      <NavLink
        to="about"
      >
        About
      </NavLink>
      <button>Contact</button>
    </div>
  )
}

export default Home
