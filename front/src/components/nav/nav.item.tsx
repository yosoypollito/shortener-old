import type { NavItemProps } from "@src/types"

import { Link } from 'react-router-dom'

export default function NavItem({ to, label }:NavItemProps){

  return(
    <Link to={to}>
      {label}
    </Link>
  )
}
