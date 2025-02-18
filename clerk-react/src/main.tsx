import { StrictMode } from 'react'
import { createRoot } from 'react-dom/client'
import './index.css'
import App from './App.tsx'
import { ClerkProvider } from '@clerk/clerk-react'
import { Session } from './pages/Session.tsx'
import { BrowserRouter, Routes, Route, Navigate } from 'react-router-dom';
import Home from './pages/Home.tsx'


const PUBLISHABLE_KEY = import.meta.env.VITE_CLERK_PUBLISHABLE_KEY

if (!PUBLISHABLE_KEY) {
  throw new Error("Missing Publishable Key")
}

createRoot(document.getElementById('root')!).render(
  <StrictMode>
     <ClerkProvider publishableKey={PUBLISHABLE_KEY} afterSignOutUrl="/">
     <BrowserRouter>
    <Routes>
      <Route path="/" element={<App />}>
        <Route path="search" element={<Session />} />
        <Route path="home" element={<Home />} />
        <Route path="" element={<Navigate to="/home" replace />} />
      </Route>
    </Routes>
  </BrowserRouter>
     
     </ClerkProvider>
  </StrictMode>,
)
