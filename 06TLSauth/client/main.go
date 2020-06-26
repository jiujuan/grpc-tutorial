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
    // credentials.NewClientTLSFromFile函数是从文件为服务器构造证书对象，
    // 然后通过grpc.WithTransportCredentials(creds)函数将证书包装为选项后作为参数传入grpc.Dial函数
    // 在客户端基于服务器的证书和服务器名字就可以对服务器进行验证
    creds, err := credentials.NewClientTLSFromFile("./keys/server.pem", "gprc-auth-name")
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
