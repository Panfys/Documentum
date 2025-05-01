Электронный документооборот для государственных структур, 
построен на основе Инструкции по делопроизводству в ВС РФ.

Для запуска:
1) Загрузите проверк по командер "go get github.com/Panfys/Documentum"
2) Cоздайте в коревой дирректории файл ".env"
с примерыным содержанием:

# MySQL settings
MYSQL_ROOT_PASSWORD=12345678
MYSQL_DATABASE=documentum
MYSQL_USER=user
MYSQL_PASSWORD=12345678

# GO settings
DB_CONNECTION_STRING=user:12345678@tcp(db:3306)/documentum?parseTime=true
DB_ROOT_CONNECTION_STRING=root:12345678@tcp(db:3306)/documentum?parseTime=true
JWT_SECRET_KEY=ASs@$%dasewE123AFSDGf325@&41sdafHAJvs!@#$%^&*

3) Запустите сервер командой "docker compose up"
4) Перейдите в браузере по адресу "http://localhost/" (подставте вместо localhost свое значение)  
