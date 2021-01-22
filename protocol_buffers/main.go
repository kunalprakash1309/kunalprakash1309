package main

import (
	"encoding/json"
	"fmt"

	_ "github.com/golang/protobuf/proto"
	pb "github.com/kunalprakash1309/protocol_buffers/protofiles"
)

func main(){
	// p := &pb.Person{
	// 	Id: 1234,
	// 	Name: "Kunal P",
	// 	Email: "kunalprakash1309@gmail.com",
	// 	Phones: []*pb.Person_PhoneNumber{
	// 		{Number: "555-4321", Type: pb.Person_HOME},
	// 	},
	// }

	// p1 := &pb.Person{}

	// body, _ := proto.Marshal(p)
	// _ = proto.Unmarshal(body, p1)

	// fmt.Println("Original struct loaded from proto files:", p, "\n")
	// fmt.Println("Marshaled proto data, ", body, "\n")
	// fmt.Println("Unmarshaled struct: ", p1)

	p := &pb.Person{
		Id: 1234,
		Name: "Roger F",
		Email: "kunalprakash1309@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "555-4321", Type: pb.Person_HOME},
		},
	}

	body, _ := json.Marshal(p)
	fmt.Println(string(body))
}