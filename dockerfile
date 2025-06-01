# Development stage with air
FROM golang:1.24.3-alpine

# Устанавливаем зависимости и air
RUN apk add --no-cache curl git bash \
    && curl -sSfL https://raw.githubusercontent.com/cosmtrek/air/master/install.sh | sh -s \
    && mkdir -p /app

WORKDIR /app

# Копируем только необходимые файлы (оптимизация слоев)
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы
COPY . .

# Устанавливаем air в PATH
ENV PATH="/go/bin:${PATH}"

# Проверяем, что air установился
RUN air -v

# Запускаем air
CMD ["air"]