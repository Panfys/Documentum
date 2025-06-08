import React from 'react';
import styles from './Layout.module.css'; // Или используйте styled-components
import SidePanel from '../sidesPanel/SidePanel';
import AuthContainer from '../auth/AuthContainer';

const Layout = ({ }) => {
  return (
    <div className={styles.appContainer}>
    <SidePanel>

    </SidePanel>
    <AuthContainer>

    </AuthContainer>
     <SidePanel>

    </SidePanel>
    </div>
  );
};

export default Layout;