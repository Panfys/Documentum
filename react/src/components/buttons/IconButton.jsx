import React from 'react';
import styles from './IconButton.module.css';

const IconButton = ({ 
  icon: Icon, 
  onClick, 
  className = '',
  iconClassName = '',
  color = "currentColor", // Дефолтный цвет
  hoverColor = "red", // Цвет при наведении
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
        stroke={color} // Передаём цвет в иконку
        data-hover-color={hoverColor} // Для CSS
      />
    </button>
  );
};

export default IconButton;