package main

import (
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/server_side_stream"
    "log"
    "net"
)
//模拟数据
var users = map[int32]pb.UserResponse{
    1: {Name: "Tom one", Age: 30},
    2: {Name: "Jimmy two", Age: 50},
    3: {Name: "Anny three", Age: 35},
}

type serverSideStreamServer struct{}

// serverSideStreamServer 实现了 user.pb.go 中的 UserServiceServer 接口
func (s *serverSideStreamServer) GetUserInfo(req *pb.UserRequest, stream pb.UserService_GetUserInfoServer) error {
    //响应流数据
    for _, user := range users {
        stream.Send(&user)
    }
    log.Printf("receive request: %s \n", req)
    return nil
}

const (
    Address = ":50001"
)

func main() {
    listener, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("failed to connect server: %v", err)
    }

    //创建gRPC服务器实例
    grpcServer := grpc.NewServer()

    //向 gRPC服务器主从实例
    pb.RegisterUserServiceServer(grpcServer, &serverSideStreamServer{})

    log.Println("server running...")
    // 启动 gRPC 服务器
    // 阻塞等待客户端的调用
    grpcServer.Serve(listener)
}

