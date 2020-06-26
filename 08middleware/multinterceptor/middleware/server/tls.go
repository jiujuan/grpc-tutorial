package server

import (
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
)

func TLSServerInterceptor() grpc.ServerOption {
    creds, err := credentials.NewServerTLSFromFile("./keys/server.pem", "./keys/server.key")
    if err != nil {
        grpclog.Fatalf("failed to generate credentials %v", err)
    }
    return grpc.Creds(creds)
}
