package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    pb "grpc-tutorial/02fourinteractionmode/proto/both_side_stream"
    "time"
)

const Address = ":50001"

func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect server: %v ", err)
    }
    defer conn.Close()

    //创建gRPC客户端实例
    grpcClient := pb.NewUserServiceClient(conn)
    stream, err := grpcClient.GetUserInfo(context.Background())
    if err != nil {
        log.Fatalf("receive stream error: %v", err)
    }

    //向服务端发送数据流，并处理响应流
    var i int32
    for i = 1; i < 4; i++ {
        stream.Send(&pb.UserRequest{ID: i})
        time.Sleep(time.Second * 1)
        resp, err := stream.Recv()
        if err != nil {
            log.Fatalf("response error: %v", err)
        }
        //输出响应
        log.Printf("receive response: %v \n", resp)
    }
}
