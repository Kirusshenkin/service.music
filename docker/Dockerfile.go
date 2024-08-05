# Указываем базовый образ
FROM golang:1.21

# Устанавливаем рабочую директорию
WORKDIR /app

# Копируем go.mod и go.sum для установки зависимостей
COPY go-server/go.mod go-server/go.sum ./
RUN go mod download

# Копируем остальные файлы из go-server в рабочую директорию
COPY go-server/ ./

# Устанавливаем зависимости и собираем приложение
WORKDIR /app/cmd
RUN go build -o /app/server main.go

# Указываем команду запуска
CMD ["/app/server"]
