import React from 'react'
import ReactDOM from 'react-dom/client'

import '@styles/globals.css'

import store from '@redux/store'
import { Provider } from 'react-redux'

import { RouterProvider } from 'react-router-dom'
import router from '@routes/router'

ReactDOM.createRoot(document.getElementById('root') as HTMLElement).render(
  <React.StrictMode>
    <Provider store={store}>
      <RouterProvider router={router}/>
    </Provider>
  </React.StrictMode>,
)
