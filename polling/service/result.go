package service

import (
	"context"
	pb "poll-service/genprotos"
	st "poll-service/storage"
)

type ResultService struct {
	storage st.StorageI
	pb.UnimplementedResultServiceServer
}

func NewResultService(storage st.StorageI) *ResultService {
	return &ResultService{storage: storage}
}

func (s *ResultService) CreateResult(ctx context.Context, req *pb.CreateResultReq) (*pb.CreateResultRes, error) {
	return s.storage.Result().CreateResult(ctx, req)
}

func (s *ResultService) SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error) {
	return s.storage.Result().SavePollAnswer(ctx, req)
}

func (s *ResultService) GetResultsInExcel(ctx context.Context, req *pb.Void) (*pb.ExcelResultsRes, error) {
	return s.storage.Result().GetResultsInExcel(ctx, req)
}
