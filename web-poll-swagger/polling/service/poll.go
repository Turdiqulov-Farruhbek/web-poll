package service

import (
	"context"
	pb "poll-service/genprotos"
	st "poll-service/storage"
)

type PollService struct {
	storage st.StorageI
	pb.UnimplementedPollServiceServer
}

func NewPollService(storage st.StorageI) *PollService {
	return &PollService{storage: storage}
}

func (s *PollService) Create(ctx context.Context, req *pb.PollCreateReq) (*pb.Void, error) {
	return s.storage.Poll().Create(ctx, req)
}

func (s *PollService) GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error) {
	return s.storage.Poll().GetByID(ctx, req)
}

func (s *PollService) Update(ctx context.Context, req *pb.PollUpdateReq) (*pb.Void, error) {
	return s.storage.Poll().Update(ctx, req)
}

func (s *PollService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Poll().Delete(ctx, req)
}

func (s *PollService) GetAll(ctx context.Context, req *pb.PollGetAllReq) (*pb.PollGetAllRes, error) {
	return s.storage.Poll().GetAll(ctx, req)
}
