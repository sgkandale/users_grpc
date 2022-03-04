# syntax=docker/dockerfile:1

FROM golang:1.17-alpine

WORKDIR /app

COPY . .

RUN go mod download

RUN go build ./cmd/users_grpc

EXPOSE 8010

CMD [ "./users_grpc" ]