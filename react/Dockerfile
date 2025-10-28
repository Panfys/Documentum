FROM node:20
WORKDIR /app

# Копируем только package файлы для кэширования
COPY package*.json ./

# Очищаем кэш и устанавливаем зависимости
RUN npm cache clean --force
RUN npm install --silent

# Копируем исходный код
COPY . .

EXPOSE 5173
CMD ["npm", "run", "dev", "--host"]