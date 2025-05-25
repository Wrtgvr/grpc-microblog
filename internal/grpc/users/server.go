package usergrpc

import (
	"context"
	"fmt"

	pb "github.com/wrtgvr/grpc-microblog/gen/users/users"
	userservice "github.com/wrtgvr/grpc-microblog/internal/services/users"
	"google.golang.org/grpc"
)

type serverAPI struct {
	pb.UnimplementedUsersServer
	service userservice.UserService
}

func Register(gRPCServer *grpc.Server, srv userservice.UserService) {
	pb.RegisterUsersServer(gRPCServer, &serverAPI{service: srv})
}

func (s *serverAPI) GetUser(ctx context.Context, in *pb.GetUserRequest) (*pb.UserResponse, error) {
	fmt.Println("Got 'GetUser' request")

	res, err := s.service.GetUser(in)
	if err != nil {
		return nil, err
	}

	return res, nil
}
