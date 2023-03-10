import NavBar from '@components/nav/nav.bar'

import { Outlet } from 'react-router-dom'

import { ToastContainer } from 'react-toastify'
import 'react-toastify/dist/ReactToastify.css';

export default function Layout(){

  return(
    <main>
      <NavBar/>
      <Outlet/>

      <ToastContainer position='bottom-right'/>
    </main>
  )
}
