import styles from './MessageContainer.module.css'
import IconButton from '../buttons/IconButton';
import { History, MessageCircleWarning, X } from 'lucide-react';

const MessageContainer = ({}) => {
    return (
        <div className={styles.messageContainer}>
        <MessageCircleWarning className={styles.messageIcon}/>
       <p>Сообщения об ошибках!!!!</p>
       <IconButton icon={History}/>
       <IconButton icon={X}/>
        </div>
    )
}

export default MessageContainer;