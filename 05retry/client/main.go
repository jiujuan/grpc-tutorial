package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-tutorial/05retry/proto/hello"
    "log"
    "time"
)

const (
    Address = ":50001"
)

var (
    retryPolicy = `{
       "methodConfig": [{
          "name": [{"service": "grpc-tutorial.05retry.proto.hello"}],
          "waitForReady": true,
          "retryPolicy": {
               "MaxAttempts": 4,
               "InitialBackoff": ".01s",
               "MaxBackoff": ".01s",
               "BackoffMultiplier": 1.0,
                "RetryableStatusCodes": ["UNAVAILABLE"]
           }
       }]}`
)

func main() {
    conn, err := grpc.Dial(Address, grpc.WithInsecure(), grpc.WithDefaultServiceConfig(retryPolicy))
    if err != nil {
        log.Fatalf("can not connect : %v", err)
    }
    defer func() {
        if e := conn.Close(); e != nil {
            log.Printf("failed to close connection : %v", e)
        }
    }()

    client := pb.NewHelloClient(conn)

    ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
    defer cancel()

    reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "TryTry"})
    if err != nil {
       log.Fatalf("sayhello err: %v", err)
    }
    log.Printf("sayhello reply: %v", reply)

    //for range time.Tick(time.Second) {
    //    ctx, cancel := context.WithTimeout(context.Background(), 1 * time.Second)
    //    reply, err := client.SayHello(ctx, &pb.HelloRequest{Name: "trytry"})
    //    if err != nil {
    //        log.Printf("sayhello err: %v", err)
    //    } else {
    //        log.Printf("message: %s", reply.Message)
    //    }
    //    cancel()
    //}
}
