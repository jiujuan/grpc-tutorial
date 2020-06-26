package main

import (
    "context"
    grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
    grpc_auth "github.com/grpc-ecosystem/go-grpc-middleware/auth"
    grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    servermid "grpc-tutorial/08middleware/multinterceptor/middleware/server"
    pb "grpc-tutorial/08middleware/multinterceptor/proto/hello"
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

    s := grpc.NewServer(servermid.TLSServerInterceptor(),
        grpc.StreamInterceptor(grpc_middleware.ChainStreamServer(
            grpc_auth.StreamServerInterceptor(servermid.AuthInterceptor),
            grpc_recovery.StreamServerInterceptor(servermid.RecoveryInterceptor()),
        )),
        grpc.UnaryInterceptor(grpc_middleware.ChainUnaryServer(
            grpc_auth.UnaryServerInterceptor(servermid.AuthInterceptor),
            grpc_recovery.UnaryServerInterceptor(servermid.RecoveryInterceptor()),
         )),
    )

    pb.RegisterHelloServer(s, &Hello{})

    log.Println("Listen on ", Address, " with TLS")
    if err := s.Serve(listen); err != nil {
        grpclog.Fatalf("failed to serve : %v", err)
    }
}