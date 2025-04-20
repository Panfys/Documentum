-- Создание пользователя (если не создан через environment)
CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY 'qweR1234';

-- Назначение прав
GRANT SELECT, INSERT, UPDATE, DELETE ON `documentum`.* TO 'user'@'%';

FLUSH PRIVILEGES;