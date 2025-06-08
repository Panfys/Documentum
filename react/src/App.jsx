import { useState } from 'react'
import './App.css'
import Layout from './components/lyouts/MainLyout'
import MessageContainer from './components/messages/MessageContainer'

function App() {
  const [count, setCount] = useState(0)

  return (
    <>
      <Layout>
        
      </Layout>
      <MessageContainer>
        
      </MessageContainer>
    </>
  )
}

export default App
