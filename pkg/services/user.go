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

func (s *Server) AddName(ctx context.Context, req *pb.AddNameRequest) (*pb.AddNameResponse, error) {
	user := models.User{
		ChatID:   req.ChatId,
		Username: req.Username,
		Name:     &req.Name,
	}

	if err := s.DbHandler.DB.Create(&user).Error; err != nil {
		return &pb.AddNameResponse{
			Message: "smth happened...",
		}, err
	}

	return &pb.AddNameResponse{
		Message: "success",
	}, nil
}

func (s *Server) AddPhone(ctx context.Context, req *pb.AddPhoneRequest) (*pb.AddPhoneResponse, error) {
	user := models.User{
		ChatID: req.ChatId,
		Phone:  &req.Phone,
	}

	if err := s.DbHandler.DB.Where(&models.User{ChatID: user.ChatID}).Updates(&user).Error; err != nil {
		return &pb.AddPhoneResponse{
			Message: "smth happened...",
		}, err
	}

	return &pb.AddPhoneResponse{
		Message: "success",
	}, nil
}

func (s *Server) IsAdmin(ctx context.Context, req *pb.IsAdminRequest) (*pb.IsAdminResponse, error) {
	var user models.User

	if err := s.DbHandler.DB.Where(&models.User{ChatID: req.ChatId}).First(&user).Error; err != nil {
		return nil, err
	}

	return &pb.IsAdminResponse{
		IsAdmin: user.IsAdmin,
	}, nil
}

func (s *Server) GetUsersByChatId(ctx context.Context, req *pb.GetUsersByChatIdRequest) (*pb.GetUsersByChatIdResponse, error) {
	var users []models.User

	if err := s.DbHandler.DB.Where("chat_id in (?)", req.ChatId).Find(&users).Error; err != nil {
		return nil, err
	}

	data := make([]*pb.UserData, len(users))
	for i, user := range users {
		data[i] = &pb.UserData{
			Username: user.Username,
			Name:     *user.Name,
			Phone:    *user.Phone,
		}
	}

	return &pb.GetUsersByChatIdResponse{
		Data: data,
	}, nil
}
