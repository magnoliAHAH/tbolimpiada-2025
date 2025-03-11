# Используем официальный образ Go для сборки
FROM golang:1.22 AS builder

# Собираем game
WORKDIR /app/game
COPY game/ .
RUN go mod tidy && go build -o /bin/game

# Собираем server
WORKDIR /app/server
COPY server/ .
RUN go mod tidy && go build -o /bin/server

# Финальный образ
FROM debian:bookworm-slim

# Копируем бинарники из builder
COPY --from=builder /bin/game /bin/game
COPY --from=builder /bin/server /bin/server

# Создаём директорию для временных данных
RUN mkdir /maps
WORKDIR /maps

# Указываем, какой порт слушает сервер
EXPOSE 8080

# Запуск API
CMD ["/bin/server"]
