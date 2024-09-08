package storage

import (
	"context"
	pb "poll-service/genprotos"
)

type StorageI interface {
	User() UserI
	Poll() PollI
	Question() QuestionI
	Result() ResultI
}

type UserI interface {
	Register(ctx context.Context, req *pb.RegisterReq) (*pb.Void, error)
	ConfirmUser(ctx context.Context, req *pb.ConfirmUserReq) (*pb.Void, error)
	Profile(ctx context.Context, req *pb.GetProfileReq) (*pb.GetProfileResp, error)
	UpdatePassword(ctx context.Context, req *pb.UpdatePasswordReq) (*pb.Void, error)
	IsEmailExists(ctx context.Context, req *pb.IsEmailExistsReq) (*pb.IsEmailExistsResp, error)
	GetByID(ctx context.Context, req *pb.GetProfileByIdReq) (*pb.GetProfileByIdResp, error)
	GetByEmail(ctx context.Context, req *pb.ByEmail) (*pb.GetProfileByIdResp, error)
}

type PollI interface {
	Create(ctx context.Context, req *pb.PollCreateReq) (*pb.Void, error)
	GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error)
	Update(ctx context.Context, req *pb.PollUpdateReq) (*pb.Void, error)
	Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error)
	GetAll(ctx context.Context, req *pb.PollGetAllReq) (*pb.PollGetAllRes, error)
}

type QuestionI interface {
	Create(ctx context.Context, req *pb.QuestionCreateReq) (*pb.Void, error)
	GetByID(ctx context.Context, req *pb.ByID) (*pb.QuestionGetByIDRes, error)
	Update(ctx context.Context, req *pb.QuestionUpdateReq) (*pb.Void, error)
	Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error)
	GetAll(ctx context.Context, req *pb.QuestionGetAllReq) (*pb.QuestionGetAllRes, error)
}

type ResultI interface {
	CreateResult(ctx context.Context, req *pb.CreateResultReq) (*pb.CreateResultRes, error)
	SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error)
	GetResultsInExcel(ctx context.Context, req *pb.Void) (*pb.ExcelResultsRes, error)
	GetByIDRes(ctx context.Context, req *pb.ByIDs) (*pb.ByIDResponse, error)
}
