-- Создание пользователя (если не создан через environment)
CREATE USER IF NOT EXISTS 'user'@'%' IDENTIFIED BY '${MYSQL_PASSWORD}';

-- Назначение прав
GRANT SELECT, INSERT, UPDATE, DELETE ON `documentum`.* TO 'user'@'%';

FLUSH PRIVILEGES;