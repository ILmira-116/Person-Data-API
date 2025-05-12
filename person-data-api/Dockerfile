# Официальный образ Go нужной версии
FROM golang:1.23.4

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем go.mod и go.sum, чтобы сначала подтянуть зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем всё остальное
COPY . .

# Собираем бинарник
RUN go build -o main ./cmd/main.go

# Указываем команду по умолчанию
CMD ["./main"]