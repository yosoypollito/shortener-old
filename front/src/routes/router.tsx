import { createBrowserRouter } from 'react-router-dom'


//There goes the pages
import Layout from '@src/layout'

import AuthLayout from '@components/auth/auth.layout'

import Home from '@pages/home'
import Donation from "@pages/donation"

import Login from "@components/auth/auth.login"
import Register from "@components/auth/auth.register"

import Error from '@pages/error'

const router = createBrowserRouter([
  {
    path:'/app/',
    element:<Layout/>,
    errorElement:<Error/>,
    children:[
      {
        path:'',
        element:<Home/>
      },
      {
        path:'donate',
        element:<Donation/>
      },
      {
        path:'auth',
        element:<AuthLayout/>,
        errorElement:<Error/>,
        children:[
          {
            path:'',
            element:<Login/>
          },
          {
            path:'register',
            element:<Register/>
          },
        ]
      },
      {
        path:'*',
        element:<Error/>,
      }
    ]
  }
])

export default router;
