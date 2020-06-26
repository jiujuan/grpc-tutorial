package main

import (
    "context"
    grpc_retry "github.com/grpc-ecosystem/go-grpc-middleware/retry"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    pb "grpc-tutorial/08middleware/retry/proto/hello"
    "log"
    "time"
)

const (
    Address = ":50001"
)
func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure(),
        // 添加拦截器
        grpc.WithUnaryInterceptor(grpc_retry.UnaryClientInterceptor(
            grpc_retry.WithCodes(codes.Canceled, codes.Unavailable, codes.NotFound),
            grpc_retry.WithMax(4),
            grpc_retry.WithPerRetryTimeout(time.Second * 1),
        )),
    )
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer conn.Close()

    client := pb.NewHelloClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), time.Second * 1)
    defer cancel()

    resp, err := client.SayHello(ctx, &pb.HelloRequest{Name: "world"})
    if err != nil {
        log.Fatalf("failed to call func: （sayhello: %v", err)
    }
    log.Fatalln("response: ", resp.Message)
}
