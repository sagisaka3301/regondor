import React, { ReactNode, useState, useEffect } from 'react'
import '../css/guest.css'

interface GuestLayoutProps {
  children: ReactNode
}

const GuestLayout: React.FC<GuestLayoutProps> = ({ children }) => {
  // Initialize isDark state using localStorage value or default to false
  const [isDark, setIsDark] = useState(
    localStorage.getItem('isDark') === 'true' // Parse string from localStorage
  )

  // Update localStorage whenever isDark state changes
  useEffect(() => {
    localStorage.setItem('isDark', isDark.toString()) // Convert boolean to string
  }, [isDark])

  const toggleDark = () => {
    setIsDark((prev) => !prev)
  }

  return (
    <div className="wrap">
      <div className={`guest-container ${isDark ? 'back-dark' : ''}`}>
        <div className="back-toggle">
          <button
            className={`toggle-btn ${isDark ? 'btn-dark' : ''}`}
            onClick={toggleDark}
          >
            {isDark ? 'Bright' : 'Dark'}
          </button>
        </div>
        <div className="ent-title">
          <h1>Regondor</h1>
        </div>
        <main>{children}</main>
      </div>
    </div>
  )
}

export default GuestLayout
