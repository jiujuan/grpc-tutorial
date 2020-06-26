package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/credentials"
    "google.golang.org/grpc/grpclog"
    pb "grpc-tutorial/06TLSauth/proto/hello"
    "log"
)

const (
    Address = ":50001"
)

func main() {
    // TLS连接
    creds, err := credentials.NewClientTLSFromFile("../keys/server.pem", "gprc-auth-name")
    if err != nil {
        grpclog.Fatalf("failed to create TLS credentials %v ", err)
    }

    conn, err := grpc.Dial(Address, grpc.WithTransportCredentials(creds))
    if err != nil {
        grpclog.Fatalf("failed to connect : " , err)
    }
    defer conn.Close()

    c := pb.NewHelloClient(conn)
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: " World, auth! "})
    if err != nil {
        grpclog.Fatalln(err)
    }
    log.Println(r.Message)
}
