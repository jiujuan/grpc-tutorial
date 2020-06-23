package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/simple"
    "log"
    "net"
)

//模拟数据
var users = map[int32]pb.UserResponse{
    1: {Name: "Tom one", Age: 30},
    2: {Name: "Jimmy two", Age: 50},
    3: {Name: "Anny three", Age: 35},
}

type simpleServer struct{}

// simpleServer 实现了 user.pb.go 中的 UserServiceServer 接口
func (s *simpleServer) GetUserInfo(ctx context.Context, req *pb.UserRequest) (resp *pb.UserResponse, err error) {
    if user, ok := users[req.ID]; ok {
        resp = &user
    }
    log.Printf("receive request: %v\n", req)
    return
}

const (
    Address = ":50001"
)

func main() {
    listener, err := net.Listen("tcp", Address)
    if err != nil {
        log. Fatalf("listener error : %v", err)
    }

    //创建gRPC服务器实例
    grpcServer := grpc.NewServer()

    //向gRPC注册服务
    pb.RegisterUserServiceServer(grpcServer, &simpleServer{})

    //启动gRPC服务器
    //阻塞等待客户端调用结果
    grpcServer.Serve(listener)
}

