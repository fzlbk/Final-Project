FROM golang:1.23-alpine AS builder
WORKDIR /app

# подтягиваем зависимости
COPY go.mod go.sum ./
RUN go mod download

# билдим бинарник
COPY . .
RUN go build -o user-service

# подтягиваем скрипт ожидания
RUN apk update && apk add --no-cache bash wget \
 && wget -O wait-for-it.sh \
      https://raw.githubusercontent.com/vishnubob/wait-for-it/master/wait-for-it.sh \
 && chmod +x wait-for-it.sh

# финальный образ
FROM alpine:3.17

# устанавливаем bash для скрипта
RUN apk update && apk add --no-cache bash

WORKDIR /app

COPY --from=builder /app/user-service    .
COPY --from=builder /app/wait-for-it.sh  .
COPY --from=builder /app/migrations      ./migrations

EXPOSE 8081

# ждём postgres и запускаем сервис
CMD ["bash", "./wait-for-it.sh", "ms-db:5432", "--", "./user-service"]
