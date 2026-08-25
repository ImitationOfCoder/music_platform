# СТАДИЯ 1: Сборка бинарника
FROM golang:1.26.5-bookworm AS builder

# Устанавливаем рабочую директорию внутри контейнера
WORKDIR /app

# Копируем только go.mod и go.sum (для кэширования зависимостей)
COPY go.mod go.sum ./
RUN go mod download

# Копируем остальной код (но при локальной разработке мы это перемонтируем)
COPY . .

# Собираем статический бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/out/service ./cmd/music_platform/main.go

# СТАДИЯ 2: Финальный легкий образ
FROM alpine:3.23

WORKDIR /app

# Копируем бинарник из стадии сборки
COPY --from=builder /app/out/service /app/service

# Открываем порт (например, 8080)
EXPOSE 8080

# Запускаем
CMD ["/app/service"]
