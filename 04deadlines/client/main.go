package main

import(
    "context"
    "flag"
    "fmt"
    "google.golang.org/grpc"
    "google.golang.org/grpc/codes"
    "google.golang.org/grpc/status"
    pb "grpc-tutorial/04deadlines/proto/echo"
    "log"
    "time"
)

var addr = flag.String("addr", "localhost:50052", "the address to connect to")

// unaryCall 不是stream的请求
func unaryCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    req := &pb.EchoRequest{Message: message}

    resp, err := c.UnaryEcho(ctx, req)
    fmt.Println("Resp: ", resp)
    got := status.Code(err)
    fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
}

// streamingCall，2端都是stream
func streamingCall(c pb.EchoClient, requestID int, message string, want codes.Code) {
    ctx, cancel := context.WithTimeout(context.Background(), time.Second)
    defer cancel()

    stream, err := c.BidirectionalStreamingEcho(ctx)//双向stream
    if err != nil {
        log.Printf("Send error : %v", err)
        return
    }

    err = stream.Send(&pb.EchoRequest{Message: message})//发送
    if err != nil {
        log.Printf("Send error : %v", err)
        return
    }

    _, err = stream.Recv() //接收
    //fmt.Println(err)
    got := status.Code(err)
    fmt.Printf("[%v] wanted = %v, got = %v\n", requestID, want, got)
}

func main() {
    flag.Parse()

    conn, err := grpc.Dial(*addr, grpc.WithInsecure(), grpc.WithBlock())
    if err != nil {
        log.Fatalf("did not connect : %v ", err)
    }
    defer conn.Close()

    c :=pb.NewEchoClient(conn)

    // 成功请求
    unaryCall(c, 1, "word", codes.OK)
    // 超时 deadline
    unaryCall(c, 2, "delay", codes.DeadlineExceeded)
    // A successful request with propagated deadline
    unaryCall(c, 3, "[propagate me]world", codes.OK)
    // Exceeds propagated deadline
    unaryCall(c, 4, "[propagate me][propagate me]world", codes.DeadlineExceeded)
    // Receives a response from the stream successfully.
    streamingCall(c, 5, "[propagate me]world", codes.OK)
    // Exceeds propagated deadline before receiving a response
    streamingCall(c, 6, "[propagate me][propagate me]world", codes.DeadlineExceeded)
}
