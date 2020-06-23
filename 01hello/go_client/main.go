package main

import (
    "context"
    "google.golang.org/grpc"
    "grpc-tutorial/01hello/go_client/proto/hello"
    "io"
    "log"
)

const (
    //gRPC 服务地址
    Address = ":9090"
)

func main() {
    //连接到远端的gRPC服务器
    conn, err := grpc.Dial(Address, grpc.WithInsecure())
    if err != nil {
        log.Fatalln(err)
    }
    defer conn.Close()

    //初始化客户端
    c := hello.NewHelloClient(conn)

    //调用SayHello方法
    res, err := c.SayHello(context.Background(), &hello.HelloRequest{Name: "hello world!"})
    if err != nil {
        log.Fatalln(err)
    }
    log.Println(res.Message)

    //调用 LotsOfReplies 方法
    stream, err := c.LotsOfReplies(context.Background(), &hello.HelloRequest{Name: "Hello world!"})
    if err != nil {
        log.Fatalln(err)
    }

    for {
        res, err := stream.Recv()
        if err == io.EOF {
            break
        }

        if err != nil {
            log.Printf("stream.Recv: %v ", err)
        }
        log.Printf("%s", res.Message)
    }
}
