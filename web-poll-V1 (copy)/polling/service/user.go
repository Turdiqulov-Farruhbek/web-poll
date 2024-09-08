package service

import (
	"context"
	pb "poll-service/genprotos"
	st "poll-service/storage"
)

type UserService struct {
	storage st.StorageI
	pb.UnimplementedUserServiceServer
}

func NewUserService(storage st.StorageI) *UserService {
	return &UserService{storage: storage}
}

func (s *UserService) Register(ctx context.Context, req *pb.RegisterReq) (*pb.Void, error) {
	return s.storage.User().Register(ctx, req)
}

func (s *UserService) ConfirmUser(ctx context.Context, req *pb.ConfirmUserReq) (*pb.Void, error) {
	return s.storage.User().ConfirmUser(ctx, req)
}

func (s *UserService) Profile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileResp, error) {
	return s.storage.User().Profile(ctx, req)
}

func (s *UserService) UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) (*pb.Void, error) {
	return s.storage.User().UpdatePassword(ctx, req)
}

func (s *UserService) IsEmailExists(ctx context.Context, req *pb.IsEmailExistsReq) (*pb.IsEmailExistsResp, error) {
	return s.storage.User().IsEmailExists(ctx, req)
}

func (s *UserService) GetByID(ctx context.Context, req *pb.GetProfileByIdReq) (*pb.GetProfileByIdResp, error) {
	return s.storage.User().GetByID(ctx, req)
}

func (s *UserService) GetByEmail(ctx context.Context, req *pb.ByEmail) (*pb.GetProfileByIdResp, error) {
	return s.storage.User().GetByEmail(ctx, req)
}
