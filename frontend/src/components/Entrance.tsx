import GuestLayout from './GuestLayout'
import React, { ReactNode, useState, useEffect } from 'react'
import { Link, NavLink } from 'react-router-dom'

import '../css/guest.css'

export const Entrance = () => {
  const mischief = () => {
    console.log('hello')
  }
  return (
    <GuestLayout>
      <div className="entrance-btns flex">
        <div className="to-auth">
          <Link to="/auth">ENTRANCE</Link>
        </div>

        <div className="center-img">
          <img
            id="mischief-img"
            src={`${process.env.PUBLIC_URL}/images/grim-red.png`}
            alt=""
          />
        </div>
        <div className="to-event">
          <button onClick={mischief}>MISCHIEF</button>
        </div>
      </div>
    </GuestLayout>
  )
}
