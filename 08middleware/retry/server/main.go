package main

import (
    "context"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    pb "grpc-tutorial/08middleware/retry/proto/hello"
    "log"
    "net"
    "sync"
)

const (
    Address = ":50001"
)

type failServer struct {
    pb.UnimplementedHelloServer
    mu sync.Mutex

    reqNum uint
    reqMax uint
}

func (f *failServer) failRequest() error {
    f.mu.Lock()
    defer f.mu.Unlock()
    f.reqNum++
    log.Println("wai: ", f.reqNum)
    if (f.reqMax > 0) && (f.reqNum % f.reqMax == 0) {
        log.Println("reqNum: ", f.reqNum, "reqMax: ", f.reqMax)
        return nil
    }
    return status.Errorf(codes.Unavailable, "failRequest: failing it")
}

func (f failServer) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloResponse, error) {
    if err := f.failRequest(); err != nil {
        log.Println("request failed num: ", f.reqNum)
        return nil, err
    }
    log.Println("request succeeded num: ", f.reqNum)
    return &pb.HelloResponse{Message: fmt.Sprintf("%s", in.Name)}, nil
}

func main() {
    listen, err := net.Listen("tcp", Address)
    if err != nil {
        log.Fatalf("failed to listen: %v ", err)
    }
    fmt.Println("listen on address: ", Address)

    s := grpc.NewServer()

    failService := &failServer{
        reqNum: 0,
        reqMax: 4,
    }
    pb.RegisterHelloServer(s, failService)
    if err := s.Serve(listen); err != nil {
        log.Fatalf("failed to serve: %v", err)
    }
}
