package main

import (
    "context"
    "google.golang.org/grpc"
    pb "grpc-tutorial/03customer/customer"
    "io"
    "log"
)
const (
    address = ":50051"
)

func createCustomer(client pb.CustomerClient, customer *pb.CustomerRequest) {
    resp, err := client.CreateCustomer(context.Background(), customer)
    if err != nil {
        log.Fatalf("Could not create Customer: %v ", err)
    }
    if resp.Success {
        log.Printf("A new Customer has been added with id: %d ", resp.Id)
    }
}

//// getCustomers calls the RPC method GetCustomers of CustomerServer
func getCustomers(client pb.CustomerClient, filter *pb.CustomerFilter) {
    stream, err := client.GetCustomers(context.Background(), filter)
    if err != nil {
        log.Fatalf("Error on get customers:%v", err)
    }
    for {
        customer, err := stream.Recv()
        if err == io.EOF {
            break
        }
        if err != nil {
            log.Fatalf("%v . GetCustomers(_)=_, %v", client, err)
        }
        log.Printf("Customer: %v", customer)
    }
}

func main() {
    conn, err := grpc.Dial(address, grpc.WithInsecure())
    if err != nil {
        log.Fatalf("failed to connect server: %v", err)
    }
    defer conn.Close()

    //creates a new CustomerClient
    client := pb.NewCustomerClient(conn)

    customer := &pb.CustomerRequest{
        Id:                   101,
        Name:                 "shyi val",
        Email:                "shiye@gmail.com",
        Phone:                "123-123-123",
        Addresses:            []*pb.CustomerRequest_Address{
            &pb.CustomerRequest_Address{
                Street:               "i Misson stream",
                City:                 "san city",
                State:                "ca",
                Zip:                  "94123",
                IsShippingAddress:    false,
            },
            &pb.CustomerRequest_Address{
                Street:               "greendfile",
                City:                 "123",
                State:                "CA",
                Zip:                  "98775",
                IsShippingAddress:    false,
            },
        },
    }

    //create a new customer
    createCustomer(client, customer)
    
    customer = &pb.CustomerRequest{
        Id:                   102,
        Name:  "Irene Rose",
        Email: "irene@xyz.com",
        Phone: "732-757-2924",
        Addresses: []*pb.CustomerRequest_Address{
            &pb.CustomerRequest_Address{
                Street:            "1 Mission Street",
                City:              "San Francisco",
                State:             "CA",
                Zip:               "94105",
                IsShippingAddress: true,
            },
        },
    }

    // Create a new customer
    createCustomer(client, customer)
    // Filter with an empty Keyword
    filter := &pb.CustomerFilter{Keyword:""}
    getCustomers(client, filter)
}
