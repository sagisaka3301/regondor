import React, { ReactNode } from 'react'
import Header from './Header'
import '../css/common.css'

interface AppLayoutProps {
  children: ReactNode
}

const AppLayout: React.FC<AppLayoutProps> = ({ children }) => {
  return (
    <div>
      <Header />
      <div className="app-container">
        <div className="header-pseudo"></div>
        <main>{children}</main>
      </div>
    </div>
  )
}

export default AppLayout
