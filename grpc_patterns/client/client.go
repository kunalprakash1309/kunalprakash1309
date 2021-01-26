package main

import (
	"context"
	// "io"
	"log"
	"time"

	// "github.com/golang/protobuf/ptypes/wrappers"
	pb "github.com/kunalprakash1309/grpc_patterns/datafiles"
	"google.golang.org/grpc"
)


const (
	address = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect : %v", err)
	}
	defer conn.Close()
	client := pb.NewOrderManagementClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	// // Get Order
	// retrievedOrder, err := client.GetOrder(ctx,
	// 	&wrappers.StringValue{Value: "106"})
	// log.Print("GetOrder Response --> :", retrievedOrder)

	// // Search Order : Server streaming scenario
	// searchStream, _ := client.SearchOrders(ctx, &wrappers.StringValue{Value: "Google"})
	
	// for {
	// 	searchOrder, err := searchStream.Recv()
	// 	if err == io.EOF {
	// 		log.Print("EOF")
	// 		break
	// 	}

	// 	if err == nil {
	// 		log.Print("Search Result : ", searchOrder)
	// 	}
	// }

	// Update Orders : Client streaming scenario
	updOrder1 := pb.Order{Id: "102", Items:[]string{"Kunal Pixel 3A", "Google Pixel Book"}, Destination:"Mountain View, CA", Price:1100.00}
	updOrder2 := pb.Order{Id: "103", Items:[]string{"Apple Watch S4", "Mac Book Pro", "iPad Pro"}, Destination:"San Jose, CA", Price:2800.00}
	updOrder3 := pb.Order{Id: "104", Items:[]string{"Google Home Mini", "Google Nest Hub", "iPad Mini"}, Destination:"Mountain View, CA", Price:2200.00}

	updateStream, err := client.UpdateOrders(ctx)

	if err != nil {
		log.Fatalf("%v.updateOrders(_) = _, %v", client, err)
	}

	// Updating order 1
	if err := updateStream.Send(&updOrder1); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder1, err)
	}
	
	// // Updating order 2
	if err := updateStream.Send(&updOrder2); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder2, err)
	}

	// // Updating order 3
	if err := updateStream.Send(&updOrder3); err != nil {
		log.Fatalf("%v.Send(%v) = %v", updateStream, updOrder3, err)
	}

	updateRes, err := updateStream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", updateStream, err, nil)
	}

	log.Printf("Update Orders Res: %s", updateRes)
}