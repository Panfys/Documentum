# Documentum
**Электронный документооборот для государственных структур**      
*Разработан на основе "Инструкции по делопроизводству в ВС РФ"*

# 🚀 Быстрый старт
**1. Установка зависимостей**
```bash
go get github.com/Panfys/Documentum
```
**2. Настройка окружения**

Создайте файл .env в корне проекта с содержимым:
```ini
MYSQL_ROOT_PASSWORD=12345678
MYSQL_DATABASE=documentum
MYSQL_USER=user
MYSQL_PASSWORD=12345678
DB_CONNECTION_STRING=user:12345678@tcp(db:3306)/documentum?parseTime=true
DB_ROOT_CONNECTION_STRING=root:12345678@tcp(db:3306)/documentum?parseTime=true
JWT_SECRET_KEY=ASs@$%dasewE123AFSDGf325@&41sdafHAJvs!@
```
**3. Запуск через Docker**
```bash
docker compose up
```
**4. Доступ к системе**    

Откройте в браузере: http://localhost/     
*(Замените localhost на ваш адрес, если требуется)*
  
# 📌 Особенности
- Соответствует требованиям делопроизводства ВС РФ
- Простая настройка через Docker
- JWT-аутентификация
    
# 🛠 Технологии
- Backend: Go
- База данных: MySQL
- Деплой: Docker

