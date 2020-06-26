package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    pb "grpc-tutorial/06TLSauth/proto/hello"
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
    return &pb.HelloResponse{Message: "Hello " + in.Name}, nil
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        grpclog.Fatalf("failed to listen : %v", err)
    }
    log.Println("Listen on  ", Address)
    //TLS
    creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
    if err != nil {
        grpclog.Fatalf("failed to generate credentials %v", err)
    }

    s := grpc.NewServer(grpc.Creds(creds))

    pb.RegisterHelloServer(s, &Hello{})

    grpclog.Infof("Listen on ", Address, " with TLS")
    if err := s.Serve(listen); err != nil {
        grpclog.Fatalf("failed to serve : %v", err)
    }
}