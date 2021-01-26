package main

import (
	"context"
	"fmt"
	"io"
	"strings"

	// "fmt"
	"log"
	"net"

	"github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/kunalprakash1309/grpc_patterns/datafiles"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

const (
	port = ":50051"
)

var orderMap = make(map[string]pb.Order)

type server struct {
	orderMap map[string]*pb.Order
}

// Simple RPC
func (s *server) GetOrder(ctx context.Context, orderId *wrappers.StringValue) (*pb.Order, error) {
	ord, exists := orderMap[orderId.Value]
	if exists {
		return &ord, status.Error(codes.OK, "")
	}

	return nil , status.Errorf(codes.NotFound, "Order does not exist. :", orderId)
}

// Server-side Streaming RPC
func (s *server) SearchOrders(searchQuery *wrappers.StringValue, stream pb.OrderManagement_SearchOrdersServer) error {

	for key, order := range orderMap {
		log.Print(key, order)

		for _, itemStr := range order.Items {
			log.Print(itemStr)
			if strings.Contains(itemStr, searchQuery.Value) {
				// send the matching orders in a stream
				err := stream.Send(&order)
				if err != nil {
					return fmt.Errorf("Error sending message to stream : %v", err)
				}
				log.Print("Matching order Found : " + key)
				break
			}
		}
	}
	return nil
}

// Client-side Streaming RPC
func (s *server) UpdateOrders(stream pb.OrderManagement_UpdateOrdersServer) error {

	orderStr := "Updated Order IDs: "
	for {
		order, err := stream.Recv()
		if err == io.EOF{
			// Finished reading the order stream
			return stream.SendAndClose(&wrappers.StringValue{Value: "Order processed" + orderStr})
		}

		if err != nil {
			return err
		}

		// Update order
		orderMap[order.Id] = *order

		log.Printf("Order ID: %s - %s", order.Id, "Updated")
		orderStr += order.Id + ", "

	}
}

func main() {

	initSampleData()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterOrderManagementServer(s, &server{})

	fmt.Printf("Server started at %v \n", port)
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

func initSampleData() {
	orderMap["102"] = pb.Order{Id: "102", Items: []string{"Google Pixel 3A", "Mac Book Pro"}, Destination: "Mountain View, CA", Price: 1800.00}
	orderMap["103"] = pb.Order{Id: "103", Items: []string{"Apple Watch S4"}, Destination: "San Jose, CA", Price: 400.00}
	orderMap["104"] = pb.Order{Id: "104", Items: []string{"Google Home Mini", "Google Nest Hub"}, Destination: "Mountain View, CA", Price: 400.00}
	orderMap["105"] = pb.Order{Id: "105", Items: []string{"Amazon Echo"}, Destination: "San Jose, CA", Price: 30.00}
	orderMap["106"] = pb.Order{Id: "106", Items: []string{"Amazon Echo", "Apple iPhone XS"}, Destination: "Mountain View, CA", Price: 300.00}
}
