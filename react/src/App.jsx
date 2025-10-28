import { useState, useEffect } from 'react';
import SettingsPanel from './components/sidesPanel/SettingsPanel'
import MainContainer from './components/mainContainer/MainContainer'
import AuthContainer from './components/auth/AuthContainer'
import DocContainer from './components/documents/documentsContainer';
import MessageContainer from './components/messages/MessageContainer'
import { checkToken } from './api/authAPI';
import AccountPanel from './components/sidesPanel/AccountPanel';
import styles from './components/sidesPanel/SidePanel.module.css';

function App({ onAuthSuccess }) {
  const [isAuthenticated, setIsAuthenticated] = useState(false);

  useEffect(() => {
    const verifyToken = async () => {
      const { isAuthenticated } = await checkToken();
      setIsAuthenticated(isAuthenticated);
      onAuthSuccess?.(isAuthenticated);
    };
    verifyToken();
  }, [onAuthSuccess]);

  const handleAuthSuccess = () => {
    setIsAuthenticated(true);
    onAuthSuccess?.(true);
  };

  return (
    <>
      <SettingsPanel>

      </SettingsPanel>
      <MainContainer>
        {isAuthenticated ?
          <DocContainer /> : 
          <AuthContainer onSuccess={handleAuthSuccess} />}
        <MessageContainer />
      </MainContainer>
      {isAuthenticated ? <AccountPanel /> : <div className={styles.sidePanelPseudo} ></div>}
    </>
  )
}

export default App
