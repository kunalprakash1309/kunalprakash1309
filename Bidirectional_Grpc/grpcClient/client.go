package main

import (
	"context"
	"io"
	"log"

	pb "github.com/kunalprakash1309/Bidirectional_Grpc/datafiles"
	"google.golang.org/grpc"
)

const (
	address = "localhost:50051"
)

// Recievestream listens to stream contents and use them
func Recievestream(client pb.MoneyTransactionClient, request *pb.TransactionRequest){
	log.Println("Started listening to the server stream!")
	stream, err := client.MakeTransaction(context.Background(), request)
	if err != nil {
		log.Fatalf("%v.MakeTrasaction(_) = _, %v", client, err)
	}

	// Listen to stream of messages
	for {
		response, err := stream.Recv()
		if err == io.EOF{
			// If there are no more messages, get out of the loop
			break
		}
		if err != nil {
			log.Fatalf("%v.MakeTransaction(_) = _, %v", client, err)
		}

		log.Printf("Status: %v, Operation: %v", response.Status, response.Description)
	}
}

func main() {
	// Set up a connection to the server.
	conn, err := grpc.Dial(address, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Did not connect: %v", err)
	}

	defer conn.Close()
	client := pb.NewMoneyTransactionClient(conn)

	//Prepare data. Get this from clients like Front-end
	from := "1234"
	to := "5678"
	amount := float32(1250.75)

	// contact the server and print out its response.
	Recievestream(client, &pb.TransactionRequest{From: from,
		To: to,
		Amount: amount,
	})
}