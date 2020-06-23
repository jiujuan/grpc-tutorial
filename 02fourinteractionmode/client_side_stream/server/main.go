package main

import (
    "google.golang.org/grpc"
    pb "grpc-tutorial/02fourinteractionmode/proto/client_side_stream"
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

type clientSideStreamServer struct{}

func (s *clientSideStreamServer) GetUserInfo(stream pb.UserService_GetUserInfoServer) error {
    var lastID int32
    for {
        req, err := stream.Recv()
        //客户端数据发送完毕
        if err == io.EOF {
            //返回最后一个ID信息
            if u, ok := users[lastID]; ok {
                _ = stream.SendAndClose(&u)
                return nil
            }
        }
        lastID = req.ID
        log.Printf("receive request: %v \n", req)
    }
    return nil
}

const (
    Address = ":50001"
)

func main() {
    listener, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("listen error : %v", err)
    }

    //创建gRPC实例
    grpcServer := grpc.NewServer()

    //向gRPC服务器注册服务
    pb.RegisterUserServiceServer(grpcServer, &clientSideStreamServer{})

    log.Println("server running ... ")
    //启动gRPC服务器
    //阻塞等待客户端调用
    _ = grpcServer.Serve(listener)

}