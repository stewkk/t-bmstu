# syntax=docker/dockerfile:1

FROM golang:1.20 AS build

WORKDIR /app

COPY go.mod go.sum ./
RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o main ./cmd/server/main.go

FROM alpine:latest

WORKDIR /app
COPY --from=build /app/main ./main
COPY --from=build /app/web/templates/* /app/web/templates/
RUN chmod +x ./main
EXPOSE 8080

CMD ["./main"]
