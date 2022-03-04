## Users Application with gRPC


### Build
    go build ./cmd/users_grpc


### Docker
    docker build -t users_grpc .
    docker run -p 8080:8080 users_grpc
