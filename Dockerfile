# Используем официальный образ Go для сборки нашего приложения
FROM golang:1.22 AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальные файлы проекта
COPY . .

# Сборка нашего Go-приложения
RUN go build -o service.music cmd/service.music/main.go

# Используем официальный образ Debian для запуска нашего приложения
FROM debian:latest

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /root/

# Копируем скомпилированное приложение из предыдущего контейнера
COPY --from=builder /app/service.music .

# Устанавливаем make, так как это может быть необходимо для дальнейшей автоматизации
RUN apt-get update && apt-get install -y make

# Устанавливаем переменные окружения для PostgreSQL
ENV POSTGRES_DB=musicdb
ENV POSTGRES_USER=musicuser
ENV POSTGRES_PASSWORD=musicpassword

# Устанавливаем переменные окружения для Redis
ENV REDIS_HOST=redis
ENV REDIS_PORT=6379

# Запускаем наше приложение
CMD ["./service.music"]
