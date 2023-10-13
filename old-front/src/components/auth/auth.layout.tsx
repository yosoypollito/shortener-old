import { Outlet } from "react-router-dom"

import styles from "@components/auth/auth.module.css"

export default function AuthLayout(){

  return(
    <div className={styles.authContainer}>
      <Outlet/>
    </div>
  )
}
