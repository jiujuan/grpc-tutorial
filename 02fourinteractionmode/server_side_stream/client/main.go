package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/server_side_stream"
    "io"
    "log"
)

const (
    Address = ":50001"
)

func main() {
    //建立连接
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect error: %v", err)
    }
    defer conn.Close()

    //创建gRPC客户端实例
    grpcClient := pb.NewUserServiceClient(conn)

    //调用服务端函数
    req := pb.UserRequest{Id: 1}
    stream, err := grpcClient.GetUserInfo(context.Background(), &req)
    if err != nil {
        log.Fatalf("receive response error: %v", err)
    }

    //接收数据流
    for {
        resp, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("receive error: %v", err)
        }
        log.Printf("receive response: %v\n", resp)
    }
}
