import NavItem from "@components/nav/nav.item"

import '@components/nav/nav.styles.css'

export default function NavBar(){

  return(
    <nav className="nav">
      <div className="nav__container">
        <NavItem to="" label="Home"/>
      </div>
      <div className="nav__container">
        {/*<NavItem to="auth" label="Start"/>*/}
        <NavItem to="donate" label="Contribute"/>
      </div>
    </nav>
  )
}
