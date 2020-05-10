FROM golang:latest as builder

ENV GO111MODULE=auto
ENV SERVER_PORT=:50051

WORKDIR /go/src/github.com/maei/golang_grpc_sum_servcer

COPY go.mod .
COPY go.sum .

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o grpc-server /go/src/github.com/maei/golang_grpc_sum_servcer/src/server/server.go
EXPOSE 50051

# CMD ["./grpc-server"]

FROM alpine
COPY --from=builder /go/src/github.com/maei/golang_grpc_sum_servcer/grpc-server /
CMD ["./grpc-server"]