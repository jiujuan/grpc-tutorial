package main

import (
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/both_side_stream"
    "io"
    "log"
    "net"
)
//模拟数据
var users = map[int32]pb.UserResponse{
    1: {Name: "Tom one", Age: 30},
    2: {Name: "Jimmy two", Age: 50},
    3: {Name: "Anny three", Age: 35},
}

type bothSideStreamServer struct{}

func (s *bothSideStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
    for {
        req, err := stream.Recv()
        if err == io.EOF {
            return nil
        }
        if err != nil {
            return err
        }
        u := users[req.ID]
        err = stream.Send(&u)
        if err != nil {
            return err
        }
        log.Printf("[RECEVIED REQUEST]: %v\n", req)
    }
    return nil
}

const Address = ":50001"

func main() {
    //指定服务器监听地址
    listener, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("connect server error : %v ", err)
    }

    //创建gRPC服务器实例
    grpcServer := grpc.NewServer()

    //向gRPC服务器主从服务
    pb.RegisterUserServiceServer(grpcServer, &bothSideStreamServer{})

    log.Println("server running...")
    //启动gRPC服务器
    //阻塞等待客户端的调用
    grpcServer.Serve(listener)
}