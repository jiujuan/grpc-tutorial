package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    pb "grpc-tutorial/02fourinteractionmode/proto/simple"
)

const (
    //定义服务器地址
    Address = ":50001"
)
func main() {
    // 建立连接
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect server: %v ", err)
    }
    defer conn.Close()

    // 创建gRPC客户端
    grpcClient := pb.NewUserServiceClient(conn)

    //调用服务端函数
    req := pb.UserRequest{ID: 2}
    resp, err := grpcClient.GetUserInfo(context.Background(), &req)
    if err != nil {
        log.Fatalf("receive response error: %v", err)
    }

    //输出响应
    log.Printf("receive response: %v \n", resp)
}
