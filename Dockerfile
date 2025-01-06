# Build stage
FROM golang:1.23.4-alpine3.21 AS build
WORKDIR /app
COPY . .
RUN go build -o main main.go
RUN apk add curl
RUN curl -L https://github.com/golang-migrate/migrate/releases/download/v4.18.1/migrate.linux-amd64.tar.gz | tar xvz

# Run stage
FROM alpine:3.21
WORKDIR /app
COPY --from=build /app/main .
COPY --from=build /app/migrate ./migrate
COPY db/migration ./migration
COPY app.env .
COPY init.sh .
COPY wait-for.sh .

EXPOSE 8080