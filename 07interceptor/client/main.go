package main

import (
    "context"
    "google.golang.org/grpc"
    "google.golang.org/grpc/grpclog"
    pb "grpc-tutorial/06auth/proto/hello"
    "log"
    "strings"
)

const (
    Address = ":50001"
)

// clientInterceptor，修改客户端发送给服务端的内容， 然后把服务端发送给客户端的内容也修改小
func clientInterceptor(ctx context.Context, method string, req, reply interface{}, cc *grpc.ClientConn,
    invoker grpc.UnaryInvoker, opts ...grpc.CallOption) error {
    // 修改客户端发送给服务端的内容
    if re, ok := req.(*pb.HelloRequest);ok {
        name := strings.Replace(re.Name, "World", "big big world",1)
        req = &pb.HelloRequest{Name: name}
    }

    err := invoker(ctx, method, req, reply, cc, opts...)
    if err != nil {
        log.Println("invoke error: ", method, err)
        return err
    }

    // 修改服务端发送给客户端的内容
    if replyMsg,ok:=reply.(*pb.HelloResponse);ok {
        msg := strings.Replace(replyMsg.Message, "words", "english words", 1)
        replyMsg.Message = msg
    }
    return nil
}

func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithUnaryInterceptor(clientInterceptor))
    if err != nil {
        grpclog.Fatalf("failed to connect : " , err)
    }
    defer conn.Close()

    c := pb.NewHelloClient(conn)
    name := "Hi, World"
    log.Println("client Send name: ", name)
    r, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
    if err != nil {
        grpclog.Fatalln(err)
    }
    log.Println("client Recv name: ", r.Message)
}
