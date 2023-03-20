package services

import (
	"context"
	"github.com/hovhannesyan/gsSportBot_UserService/pkg/db"
	"github.com/hovhannesyan/gsSportBot_UserService/pkg/models"
	"github.com/hovhannesyan/gsSportBot_UserService/pkg/pb"
)

type Server struct {
	DbHandler db.Handler
	pb.UnimplementedUserServer
}

func (s *Server) AddUser(ctx context.Context, req *pb.AddUserRequest) (*pb.AddUserResponse, error) {
	user := models.User{
		Id:       req.Id,
		Username: req.Username,
		Name:     &req.Name,
		Phone:    &req.Phone,
		IsAdmin:  req.IsAdmin,
	}

	if err := s.DbHandler.DB.Create(&user).Error; err != nil {
		return &pb.AddUserResponse{
			Message: "smth happened...",
		}, err
	}

	return &pb.AddUserResponse{
		Message: "success",
	}, nil
}

func (s *Server) IsAdmin(ctx context.Context, req *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	var user models.User

	if err := s.DbHandler.DB.Where(&models.User{Id: req.Id}).First(&user).Error; err != nil {
		return nil, err
	}

	return &pb.IsAdminResponse{
		IsAdmin: user.IsAdmin,
	}, nil
}

func (s *Server) GetUsersById(ctx context.Context, req *pb.GetUsersByIdRequest) (*pb.GetUsersByIdResponse, error) {
	var users []models.User

	if err := s.DbHandler.DB.Where("id in (?)", req.Id).Find(&users).Error; err != nil {
		return nil, err
	}

	data := make([]*pb.UserData, len(users))
	for i, user := range users {
		data[i] = &pb.UserData{
			Id:       user.Id,
			Username: user.Username,
			Name:     *user.Name,
			Phone:    *user.Phone,
			IsAdmin:  user.IsAdmin,
		}
	}

	return &pb.GetUsersByIdResponse{
		Data: data,
	}, nil
}
