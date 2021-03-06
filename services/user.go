package services

import (
	"context"
	"fmt"

	"github.com/jeffersonkleyton/fc-grpc/pb"
)

/* type UserServiceServer interface {
	AddUser(context.Context, *User) (*User, error)
	mustEmbedUnimplementedUserServiceServer()
} */

type UserService struct {
	pb.UnimplementedUserServiceServer
}

func NewUserService() *UserService {
	return &UserService{}
}

func (*UserService) AddUser(ctx context.Context, req *pb.User) (*pb.User, error) {
	//INSERT DATABASE

	fmt.Println(req.Name)
	return &pb.User{
		Id:    "123",
		Name:  req.GetName(),
		Email: req.GetEmail(),
	}, nil
}
