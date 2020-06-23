package main

import (
    "google.golang.org/grpc"
    "grpc-tutorial/01hello/go_server/controller/hello_controller"
    "grpc-tutorial/01hello/go_server/proto/hello"
    "log"
    "net"
)

const (
    Address = ":9090"
)

func main() {
    //服务端监听服务端口
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("Failed to listen: %v", err)
    }

    //实例化一个gRPC服务器
    s := grpc.NewServer()
    //在gRPC上注册服务
    hello.RegisterHelloServer(s, &hello_controller.HelloController{})

    log.Println("Listen on " + Address)

    //启动服务
    if err := s.Serve(listen); err != nil {
        log.Fatalf("Failed to serve: %v", err)
    }
}
