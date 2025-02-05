import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import Layout from './layout'
import Login from './pages/login/page'

createRoot(document.getElementById('root')!).render(
  <StrictMode>
    <Layout>
      <Login />
    </Layout>
  </StrictMode>,
)
