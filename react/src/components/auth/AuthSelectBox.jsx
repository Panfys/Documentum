import { useState, useRef, useEffect } from 'react';
import styles from './AuthContainer.module.css';
import { ChevronDown, CircleX } from 'lucide-react';

const AuthSelectBox = ({ label, name, tabIndex, value, onChange, options, error }) => {
  const [isOpen, setIsOpen] = useState(false);
  const [focusedIndex, setFocusedIndex] = useState(-1);
  const selectRef = useRef(null);
  const optionsRef = useRef(null);

  useEffect(() => {
    const handleClickOutside = (e) => {
      if (selectRef.current && !selectRef.current.contains(e.target)) {
        setIsOpen(false);
        setFocusedIndex(-1);
      }
    };

    document.addEventListener('mousedown', handleClickOutside);
    return () => document.removeEventListener('mousedown', handleClickOutside);
  }, []);

  useEffect(() => {
    if (isOpen && optionsRef.current) {
      const selectedOption = options.findIndex(option => option.value === value);
      setFocusedIndex(selectedOption >= 0 ? selectedOption : 0);
    }
  }, [isOpen, options, value]);

  const handleSelect = (selectedValue) => {
    if (onChange) {
      onChange({
        target: {
          name: name,
          value: selectedValue
        }
      });
    }
    setIsOpen(false);
    setFocusedIndex(-1);
  };

  const handleKeyDown = (e) => {
    if (!isOpen) {
      // Открытие селекта
      if (e.key === 'Enter' || e.key === ' ' || e.key === 'ArrowDown' || e.key === 'ArrowUp') {
        e.preventDefault();
        setIsOpen(true);
      }
    } else {
      // Навигация по опциям когда селект открыт
      switch (e.key) {
        case 'Enter':
          e.preventDefault();
          if (focusedIndex >= 0 && focusedIndex < options.length) {
            handleSelect(options[focusedIndex].value);
          }
          break;
        
        case ' ':
          e.preventDefault();
          if (focusedIndex >= 0 && focusedIndex < options.length) {
            handleSelect(options[focusedIndex].value);
          }
          break;
        
        case 'Escape':
          e.preventDefault();
          setIsOpen(false);
          setFocusedIndex(-1);
          break;
        
        case 'ArrowDown':
          e.preventDefault();
          setFocusedIndex(prev => 
            prev < options.length - 1 ? prev + 1 : 0
          );
          break;
        
        case 'ArrowUp':
          e.preventDefault();
          setFocusedIndex(prev => 
            prev > 0 ? prev - 1 : options.length - 1
          );
          break;
        
        case 'Tab':
          if (focusedIndex >= 0 && focusedIndex < options.length) {
            e.preventDefault();
            handleSelect(options[focusedIndex].value);
          }
          setIsOpen(false);
          setFocusedIndex(-1);
          break;
        
        case 'Home':
          e.preventDefault();
          setFocusedIndex(0);
          break;
        
        case 'End':
          e.preventDefault();
          setFocusedIndex(options.length - 1);
          break;
        
        default:
          // Поиск по первой букве
          if (e.key.length === 1 && e.key.match(/[a-z0-9]/i)) {
            const foundIndex = options.findIndex(option => 
              option.label.toLowerCase().startsWith(e.key.toLowerCase())
            );
            if (foundIndex >= 0) {
              setFocusedIndex(foundIndex);
            }
          }
          break;
      }
    }
  };

  const selectedOption = options.find(option => option.value === value);
  const displayValue = selectedOption ? selectedOption.label : '';

  return (
    <>
      <div className={styles.authInputBox} ref={selectRef}>
        <label htmlFor={`${name}-hidden`}>{label}</label>
        <div 
          className={`${styles.authInput} ${styles.customSelect} ${isOpen ? styles.selectOpen : ''}`}
          tabIndex={tabIndex}
          onClick={() => {
            setIsOpen(!isOpen);
            setFocusedIndex(-1);
          }}
          onKeyDown={handleKeyDown}
          role="combobox"
          aria-haspopup="listbox"
          aria-expanded={isOpen}
          aria-controls={`${name}-options`}
          aria-activedescendant={isOpen && focusedIndex >= 0 ? `${name}-option-${focusedIndex}` : undefined}
        >
          <span className={styles.selectValue}>{displayValue}</span>
          <ChevronDown 
            className={`${styles.arrow} ${isOpen ? styles.arrowOpen : ''}`} 
            size={16}
          />
        </div>
        
        <div 
          ref={optionsRef}
          className={`${styles.optionsContainer} ${isOpen ? styles.optionsOpen : ''}`}
          id={`${name}-options`}
          role="listbox"
          aria-label="Выберите опцию"
        >
          {options.map((option, index) => (
            <div
              key={option.value}
              id={`${name}-option-${index}`}
              className={`${styles.option} ${value === option.value ? styles.selected : ''} ${focusedIndex === index ? styles.focused : ''}`}
              onClick={() => handleSelect(option.value)}
              onMouseEnter={() => setFocusedIndex(index)}
              role="option"
              aria-selected={value === option.value}
            >
              {option.label}
            </div>
          ))}
        </div>
        
        {/* Скрытый нативный select для формы */}
        <select 
          id={`${name}-hidden`}
          className={styles.hiddenSelect}
          name={name}
          value={value || ''}
          onChange={onChange}
          tabIndex={-1}
        >
          <option value=""></option>
          {options.map(option => (
            <option key={option.value} value={option.value}>
              {option.label}
            </option>
          ))}
        </select>
      </div>
      <div className={styles.authErrorBox}>
        {error && <CircleX className={styles.authErrorIcon}/>}
        <p>{error}</p>
      </div>
    </>
  );
};

export default AuthSelectBox;