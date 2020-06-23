package main

import (
    "context"
    "google.golang.org/grpc"
    "log"
    pb "grpc-tutorial/03customer/customer"
    "net"
    "strings"
)

const (
    port = ":50051"
)

type server struct {
    savedCustomers []*pb.CustomerRequest
}

func (s *server) CreateCustomer(ctx context.Context, in *pb.CustomerRequest) (*pb.CustomerResponse, error) {
    s.savedCustomers = append(s.savedCustomers, in)
    return &pb.CustomerResponse{Id: in.Id, Success: true}, nil
}

func (s *server) GetCustomers(filter *pb.CustomerFilter, stream pb.Customer_GetCustomersServer) error {
    for _, customer := range s.savedCustomers {
        if filter.Keyword != "" {
            if !strings.Contains(customer.Name, filter.Keyword) {
                continue
            }
        }
        if err := stream.Send(customer); err != nil {
            return err
        }
    }
    return nil
}

func main() {
    listener, err := net.Listen("tcp", port)
    if err != nil {
        log.Fatalf("failed to listen: %v", err)
    }

    s := grpc.NewServer()
    pb.RegisterCustomerServer(s, &server{})
    s.Serve(listener)
}
