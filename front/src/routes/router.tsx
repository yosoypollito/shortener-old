import { createBrowserRouter } from 'react-router-dom'


//There goes the pages
import Layout from '@src/layout'
import Home from '@pages/home'
import Donation from "@pages/donation"

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
        path:'*',
        element:<Error/>
      }
    ]
  }
])

export default router;
