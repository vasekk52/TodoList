# Этап сборки
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Копируем зависимости
COPY go.mod go.sum ./
RUN go mod download

# Копируем исходный код
COPY . .

# Собираем бинарник
RUN CGO_ENABLED=0 GOOS=linux go build -o todo ./cmd/todo

# Финальный образ
FROM alpine:3.19

WORKDIR /app

# Копируем бинарник из этапа сборки
COPY --from=builder /app/todo .

# Для безопасности запускаем не от root
RUN adduser -D -g '' appuser
USER appuser

EXPOSE 8080

CMD ["./todo"]