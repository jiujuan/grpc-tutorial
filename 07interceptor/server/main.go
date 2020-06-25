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

func UnaryServerInterceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo,
                      handler grpc.UnaryHandler) (interface{}, error) {
    log.Printf("before handling. Info: %+v", info)
    resp, err := handler(ctx, req)
    log.Println("after handling. resp: ", resp)
    return resp, err
}

func StreamServerInterceptor(srv interface{}, ss grpc.ServerStream, info *grpc.StreamServerInfo,
      handler grpc.StreamHandler) error {
    log.Printf("before handling. Info: %+v", info)
    err := handler(srv, ss)
    log.Println("after handling. err: ", err)
    return err
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("failed to listen : %v", err)
    }
    log.Println("Listen on  ", Address)

    s := grpc.NewServer(grpc.StreamInterceptor(StreamServerInterceptor), grpc.UnaryInterceptor(UnaryServerInterceptor))

    pb.RegisterHelloServer(s, &Hello{})

    if err := s.Serve(listen); err != nil {
        grpclog.Fatalf("failed to serve : %v", err)
    }
}