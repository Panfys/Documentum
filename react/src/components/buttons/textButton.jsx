import styles from './IconButton.module.css';

const TextButton = ({
    icon: Icon,
    onClick,
    className = '',
    iconClassName = '',
    ...props
}) => {
    return (
        <button
            className={`${styles['textButton']} ${className}`}
            onClick={onClick}
            aria-label="button"
            {...props}
        />
    );
};

export default TextButton;