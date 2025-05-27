package userservice

import (
	usersv1 "github.com/wrtgvr/grpc-microblog/gen/users/users"
)

type UserService interface {
	GetUser(*usersv1.GetUserRequest) (*usersv1.UserResponse, error)
	CreateUser(*usersv1.CreateUserRequest) (*usersv1.UserResponse, error)
	UpdateUser(*usersv1.UpdateUserRequest) (*usersv1.UserResponse, error)
	DeleteUser(*usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error)
}

type defaultUserService struct {
	usersv1.UnimplementedUsersServer
}

func NewPostgresUserService() *defaultUserService {
	return &defaultUserService{}
}

func (s *defaultUserService) GetUser(*usersv1.GetUserRequest) (*usersv1.UserResponse, error) {
	panic("not implemented")
}

func (s *defaultUserService) CreateUser(*usersv1.CreateUserRequest) (*usersv1.UserResponse, error) {
	panic("not implemented")
}

func (s *defaultUserService) UpdateUser(*usersv1.UpdateUserRequest) (*usersv1.UserResponse, error) {
	panic("not implemented")
}

func (s *defaultUserService) DeleteUser(*usersv1.DeleteUserRequest) (*usersv1.DeleteUserResponse, error) {
	panic("not implemented")
}
