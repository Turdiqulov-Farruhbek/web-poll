package service

import (
	"context"
	"encoding/json"
	"poll-service/storage"
	pb "poll-service/genprotos"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type PollService struct {
	storage storage.StorageI
	pb.UnimplementedPollServiceServer
}

func NewPollService(storage storage.StorageI) *PollService {
	return &PollService{storage: storage}
}

func (s *PollService) Create(ctx context.Context, req *pb.PollCreateReq) (*pb.Void, error) {
	_, err := s.storage.Poll().Create(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, s.jsonError("Failed to create poll", err))
	}
	return &pb.Void{}, nil
}

func (s *PollService) GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error) {
	res, err := s.storage.Poll().GetByID(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.NotFound, s.jsonError("Poll not found", err))
	}
	return res, nil
}

func (s *PollService) Update(ctx context.Context, req *pb.PollUpdateReq) (*pb.Void, error) {
	_, err := s.storage.Poll().Update(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, s.jsonError("Failed to update poll", err))
	}
	return &pb.Void{}, nil
}

func (s *PollService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	_, err := s.storage.Poll().Delete(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, s.jsonError("Failed to delete poll", err))
	}
	return &pb.Void{}, nil
}

func (s *PollService) GetAll(ctx context.Context, req *pb.PollGetAllReq) (*pb.PollGetAllRes, error) {
	res, err := s.storage.Poll().GetAll(ctx, req)
	if err != nil {
		return nil, status.Errorf(codes.Internal, s.jsonError("Failed to get all polls", err))
	}
	return res, nil
}

func (s *PollService) jsonError(message string, err error) string {
	errorMessage := struct {
		Message string `json:"message"`
		Error   string `json:"error"`
	}{
		Message: message,
		Error:   err.Error(),
	}

	errorJSON, _ := json.Marshal(errorMessage)
	return string(errorJSON)
}
