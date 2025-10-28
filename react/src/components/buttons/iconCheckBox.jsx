import { useState } from 'react';
import { Circle, CircleCheckBig } from 'lucide-react';
import styles from './IconButton.module.css';

const IconCheckbox = ({ text, name, tabIndex, checBoxClass}) => {
  const [isChecked, setIsChecked] = useState(false);

  const handleKeyDown = (e) => {
    if (e.key === 'Enter' || e.key === ' ') {
      e.preventDefault();
      setIsChecked(!isChecked);
    }
  };

  return (
    <div className={checBoxClass}>
      <input
        id="icon-checkbox"
        type="checkbox"
        tabIndex={tabIndex}
        checked={isChecked}
        onChange={() => setIsChecked(!isChecked)}
        onKeyDown={handleKeyDown}
        className={styles.hiddenCheckbox}
        name={name}
      />
      <label 
        htmlFor="icon-checkbox"
        className={styles.checkboxLabel}
      >
        <div className={styles.iconContainer}>
          <Circle 
            className={`${styles.circleIcon} ${isChecked ? styles.hidden : ''}`}
          />
          <CircleCheckBig 
            className={`${styles.checkIcon} ${!isChecked ? styles.hidden : ''}`}
          /> 
          <span>{text}</span>
        </div>
      </label>
    </div>
  );
};

export default IconCheckbox;