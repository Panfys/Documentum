import styles from './SidePanel.module.css';
import { Settings } from 'lucide-react';
import IconButton from '../buttons/IconButton';
import { useTogglePanel } from './SidePanelHelper';

const SettingsPanel = ({ children }) => {
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
                <IconButton 
                    icon={Settings} 
                    onClick={togglePanel}
                    ref={buttonRef}
                    tabIndex={10}
                />
                <span>Настройки</span>
                <div className={styles.iconPlaceholder}></div>
            </div>
            <div className={styles.sideContent}>
                {children}
            </div>
        </div>
    );
};

export default SettingsPanel;