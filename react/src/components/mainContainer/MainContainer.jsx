
import styles from "./MainContainer.module.css";

const MainContainer = ({ children }) => {
    return (
        <div className={styles.mainPanel}>
            {children}
        </div>

    );
};

export default MainContainer;