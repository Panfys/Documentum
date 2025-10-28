import React from 'react';
import styles from './SidePanel.module.css';
import IconButton from '../buttons/IconButton';
import { User } from 'lucide-react';
import { useTogglePanel } from './SidePanelHelper'; // Импортируем тот же хук

const AccountPanel = ({ children }) => {
    const {
        isExpanded,
        togglePanel,
        panelRef,
        buttonRef
    } = useTogglePanel(false);

    return (
        <div 
            ref={panelRef}
            className={`${styles.sidePanel} ${isExpanded ? styles.expanded : styles.collapsed}`}
        >
            <div className={styles.sideTitle}>
                <div className={styles.iconPlaceholder}></div>
                <span>Логин</span>
                <IconButton 
                    icon={User} 
                    onClick={togglePanel}
                    ref={buttonRef}
                />
            </div>
            <div className={styles.sideContent}>
                {children}
            </div>
        </div>
    );
};

export default AccountPanel;