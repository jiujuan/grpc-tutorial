package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    pb "grpc-tutorial/06auth/proto/hello"
    "log"
    "net"
)

const (
    Address = ":50001"
)

type Hello struct {
    pb.UnimplementedHelloServer
}

func (h *Hello) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    log.Printf("server Req name: Server send words: %v", in.Name)
    return &pb.HelloResponse{Message: "Server send words: " + in.Name}, nil
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("failed to listen : %v", err)
    }
    log.Println("Listen on  ", Address)

    s := grpc.NewServer()

    pb.RegisterHelloServer(s, &Hello{})

    if err := s.Serve(listen); err != nil {
        grpclog.Fatalf("failed to serve : %v", err)
    }
}