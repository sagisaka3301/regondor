import React from 'react'
import { Link, NavLink } from 'react-router-dom'
// import '../css/common.css'

const Header = () => {
  return (
    <div className="header">
      <div className="header-logo">
        <img src={`${process.env.PUBLIC_URL}/images/header-logo.png`} alt="" />
      </div>
      <nav className="menu-wrap">
        <ul className="menu">
          <li>
            {/* exact:完全一致。/aboutの場合も、/がその一部であるため、activeクラスが付与されてしまうことを防ぐ。 */}
            <Link to="/todo">Home</Link>
          </li>
          <li>
            <Link to="/todo">About</Link>
          </li>
          <li>
            <Link to="/mypage">MyPage</Link>
          </li>
          <li>
            <Link to="/todo">Settings</Link>
          </li>
        </ul>
      </nav>
    </div>
  )
}

export default Header
