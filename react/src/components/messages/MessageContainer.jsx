import styles from './MessageContainer.module.css'
import IconButton from '../buttons/IconButton';

const MessageContainer = ({}) => {
    return (
        <div className={styles.messageContainer}>
        <IconButton icon={AlertCircle}></IconButton><div>Сообщения об ошибках</div><button>История</button><button>Закрыть</button>
        </div>
    )
}

export default MessageContainer;