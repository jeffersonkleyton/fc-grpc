package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/jeffersonkleyton/fc-grpc/pb"
	"google.golang.org/grpc"
)

func main() {
	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect to gRPC Server: %v", err)

	}
	defer connection.Close()

	client := pb.NewUserServiceClient(connection)
	//AddUser(client)
	//AddServerVerbose(client)
	AddUsers(client)
}

func AddUser(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jefferson",
		Email: "p.kleyton71@gmail.com",
	}

	res, err := client.AddUser(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}
	fmt.Println(res)

}

func AddServerVerbose(client pb.UserServiceClient) {
	req := &pb.User{
		Id:    "0",
		Name:  "Jefferson",
		Email: "p.kleyton71@gmail.com",
	}

	responseStream, err := client.AddUserVerbose(context.Background(), req)

	if err != nil {
		log.Fatalf("Could not make gRPC request: %v", err)
	}

	for {
		stream, err := responseStream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatalf("Could not receive the msg: %v", err)
		}

		fmt.Println("Status: ", stream.Status, " - ", stream.GetUser())
	}
}

func AddUsers(client pb.UserServiceClient) {
	reqs := []*pb.User{
		&pb.User{
			Id:    "w1",
			Name:  "Jefferson-0",
			Email: "j0@mail.com",
		},
		&pb.User{
			Id:    "w2",
			Name:  "Jefferson-1",
			Email: "j1@mail.com",
		},
		&pb.User{
			Id:    "w3",
			Name:  "Jefferson-2",
			Email: "j2@mail.com",
		},
		&pb.User{
			Id:    "w4",
			Name:  "Jefferson-3",
			Email: "j3@mail.com",
		},
		&pb.User{
			Id:    "w1",
			Name:  "Jefferson",
			Email: "j@mail.com",
		},
	}

	stream, err := client.AddUsers(context.Background())

	if err != nil {
		log.Fatalf("error creating request: %v", err)
	}

	for _, req := range reqs {
		stream.Send(req)
		time.Sleep(time.Second * 3)
	}

	res, err := stream.CloseAndRecv()

	if err != nil {
		log.Fatalf("error receiving response: %v", err)
	}

	fmt.Println(res)
}
