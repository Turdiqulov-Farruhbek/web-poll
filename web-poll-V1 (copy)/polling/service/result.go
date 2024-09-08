package service

import (
	"context"
	"fmt"
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

func (s *ResultService) GetPollResults(ctx context.Context, req *pb.ByIDs) (*pb.ByIDResponse, error) {
	var a, b int32

	resAnswer, err := s.storage.Result().GetByIDRes(ctx, req)
	if err != nil {
		return nil, err
	}
	// fmt.Println(">>",resAnswer.Feed)
	var extrovert, nevrotizm, total int32
	feed := ""
	poll, err := s.storage.Poll().GetByID(ctx, &pb.ByID{Id: *req.PollId})
	if err != nil {
		return nil, err
	}
	if *poll.PollNum == 1 {
		for _, v := range resAnswer.Answers {
			switch {
			case isIn(*v.Num, []int32{1, 3, 8, 10, 13, 17, 22, 25, 27, 39, 44, 46, 49, 53, 56}) && *v.AnswerPoint == int32(1):
				extrovert += 1
			case isIn(*v.Num, []int32{5, 15, 29, 32, 34, 37, 41, 51}) && *v.AnswerPoint == int32(0):
				extrovert += 1
			case isIn(*v.Num, []int32{2, 4, 7, 9, 11, 14, 16, 19, 21, 23, 26, 28, 31, 33, 35, 38, 40, 43, 45, 47, 50, 52, 55, 57}) && *v.AnswerPoint == int32(1):
				nevrotizm += 1
			default:
				continue
			}
		}
		feed += "Siz "

		if extrovert > 12 && extrovert > nevrotizm {
			feed = "Extrovert"
		} else if extrovert < 12 && nevrotizm < 12 {
			feed = "Introvert"
		} else if extrovert == nevrotizm && extrovert > 12 && nevrotizm < 12 {
			feed = "Extrovert va Nevrotizm"
		} else {
			feed = "Nevrotizm"
		}
		resAnswer.Feed = []*pb.Feedback{{From: &a, To: &a, Text: &feed}}

	} else {
		for _, v := range resAnswer.Answers {
			total += *v.AnswerPoint
		}
		for _, v := range poll.Feedback {
			if total >= *v.From && total < *v.To {
				feed = *v.Text
				break
			}
		}
		for _, v := range resAnswer.Feed {
			if *v.From >= total && *v.To <= total {
				feed = *v.Text
				a = *v.From
				b = *v.To
				break
			}
		}
		fmt.Println(total)
		fmt.Println(feed)

		resAnswer.Feed = []*pb.Feedback{{From: &a, To: &b, Text: &feed}}
	}

	return resAnswer, nil
}

func isIn(num int32, ls []int32) bool {
	for _, v := range ls {
		if v == num {
			return true
		}
	}
	return false
}
