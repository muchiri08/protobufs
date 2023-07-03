package main

import (
	"encoding/json"
	"fmt"
	"github.com/golang/protobuf/proto"
	pb "gitub.com/muchiri08/protobufs/protofiles"
)

func main() {
	p := &pb.Person{
		Id:    1234,
		Name:  "Muchiri K",
		Email: "muchirikennedy08@gmail.com",
		Phones: []*pb.Person_PhoneNumber{
			{Number: "0788501613", Type: pb.Person_HOME},
		},
	}

	p1 := &pb.Person{}
	body, _ := proto.Marshal(p)
	_ = proto.Unmarshal(body, p1)
	fmt.Println("Original struct loaded from proto file:", p, "\n")
	fmt.Println("Marshaled proto data: ", body, "\n")
	fmt.Println("Unmarshaled struct: ", p1)

	//Returning a json object
	fmt.Println()
	body, _ = json.Marshal(p)
	fmt.Println(string(body))
}
