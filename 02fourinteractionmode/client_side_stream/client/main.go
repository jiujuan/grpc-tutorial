package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/client_side_stream"
    "log"
)

const (
    Address = ":50001"
)
func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect server: %v", err)
    }
    defer conn.Close()

    //创建gRPC客户端实例
    grpcClient := pb.NewUserServiceClient(conn)

    //向服务端发送数据流
    stream, err := grpcClient.GetUserInfo(context.Background())
    //模拟数据库中有3条数据， ID 分别为 1,2,3
    var i int32
    for i = 1; i < 4; i++ {
        err := stream.Send(&pb.UserRequest{ID: i})
        if err != nil {
            log.Fatalf("send error: %v", err)
        }
    }

    //接收服务端响应
    resp, err := stream.CloseAndRecv()
    if err != nil {
        log.Fatalf("receive response error: %v", err)
    }
    //输出响应
    log.Printf("receive response: %v \n", resp)
}