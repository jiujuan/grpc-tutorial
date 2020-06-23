package hello_controller

import (
    "context"
    "fmt"
    hello2 "grpc-tutorial/01hello/go_server/proto/hello"
)

type HelloController struct{}

func (h *HelloController) SayHello(ctx context.Context, in *hello2.HelloRequest) (*hello2.HelloResponse, error) {
    return &hello2.HelloResponse{Message: fmt.Sprintf("%s", in.Name)}, nil
}

func (h *HelloController) LotsOfReplies(in *hello2.HelloRequest, stream hello2.Hello_LotsOfRepliesServer) error {
    for i := 0; i < 10; i++ {
        _ = stream.Send(&hello2.HelloResponse{Message: fmt.Sprintf("%s %s %d", in.Name, "Reply", i)})
    }
    return nil
}