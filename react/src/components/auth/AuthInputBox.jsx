import styles from './AuthContainer.module.css'
import { CircleX } from 'lucide-react';

const AuthInputBox = ({ label, name, tabIndex, type, value, onChange, placeholder, error }) => {

  return (
    <>
      <div className={styles.authInputBox}>
        <label htmlFor={name}>{label}</label>
        <input
          id={name}
          type={type}
          tabIndex={tabIndex}
          placeholder={placeholder}
          name={name}
          value={value || ''} // Добавляем fallback для value
          onChange={onChange} // Обязательно передаем onChange
          className={styles.authInput}
        />
      </div>
      <div className={styles.authErrorBox}>
        {error && <CircleX className={styles.authErrorIcon}/>}
        <p>{error}</p>
      </div >
    </>

  );
};

export default AuthInputBox;