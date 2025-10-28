import React from 'react';
import styles from './IconButton.module.css';

const IconButton = ({ 
  icon: Icon, 
  onClick, 
  className = '',
  iconClassName = '',
  ...props 
}) => {
  return (
    <button
      className={`${styles['icon-button']} ${className}`}
      onClick={onClick}
      aria-label="Icon button"
      {...props}
    >
      <Icon 
        className={`${styles['icon-button__icon']} ${iconClassName}`}
      />
    </button>
  );
};

export default IconButton;