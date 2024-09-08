package service

import (
	"context"
	pb "poll-service/genprotos"
	st "poll-service/storage"
)

type QuestionService struct {
	storage st.StorageI
	pb.UnimplementedQuestionServiceServer
}


func NewQuestionService(storage st.StorageI) *QuestionService {
	return &QuestionService{storage: storage}
}

func (s *QuestionService) Create(ctx context.Context, req *pb.QuestionCreateReq) (*pb.Void, error) {
	return s.storage.Question().Create(ctx, req)
}

func (s *QuestionService) GetByID(ctx context.Context, req *pb.ByID) (*pb.QuestionGetByIDRes, error) {
	return s.storage.Question().GetByID(ctx, req)
}

func (s *QuestionService) Update(ctx context.Context, req *pb.QuestionUpdateReq) (*pb.Void, error) {
	return s.storage.Question().Update(ctx, req)
}

func (s *QuestionService) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	return s.storage.Question().Delete(ctx, req)
}

func (s *QuestionService) GetAll(ctx context.Context, req *pb.QuestionGetAllReq) (*pb.QuestionGetAllRes, error) {
	return s.storage.Question().GetAll(ctx, req)
}
