package managers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"

	pb "poll-service/genprotos"

	"github.com/google/uuid"
)

type ResultManager struct {
	Conn *sql.DB
}

func NewResultManager(conn *sql.DB) *ResultManager {
	return &ResultManager{Conn: conn}
}

func (m *ResultManager) CreateResult(ctx context.Context, req *pb.CreateResultReq) (*pb.CreateResultRes, error) {
	id := uuid.NewString()
	query := `INSERT INTO results (id, user_id, poll_id) VALUES ($1, $2, $3)`
	_, err := m.Conn.ExecContext(ctx, query, id, req.UserId, req.PollId)
	if err != nil {
		return nil, err
	}
	return &pb.CreateResultRes{ResultId: &id}, nil
}

func (m *ResultManager) SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error) {
	query := "INSERT INTO poll_answers (id, result_id, question_id, answer) VALUES ($1, $2, $3, $4)"
	_, err := m.Conn.ExecContext(ctx, query, uuid.NewString(), req.ResultId, req.QuestionId, req.Answer)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *ResultManager) GetResultsInExcel(ctx context.Context, req *pb.Void) (*pb.ExcelResultsRes, error) {
	query := `
		SELECT 
			u.name, u.surname, u.gender, u.age, u.nation, u.email, u.phone_number, u.working_experience, u.level_type,
			p.poll_num,
			q.num AS question_num, q.id, pa.answer
		FROM 
			results r
		JOIN 
			users u ON r.user_id = u.id
		JOIN 
			poll_answers pa ON r.id = pa.result_id
		JOIN 
			questions q ON pa.question_id = q.id
		JOIN 
			polls p ON r.poll_id = p.id
		ORDER BY 
			u.id, p.poll_num, q.num;
	`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resultsMap := make(map[string]*pb.ResultRes)
	for rows.Next() {
		var name, surname, gender, age, nation, email, phoneNumber, level_type, questionId string
		var workingExperience, pollNum, questionNum, answer int32

		err := rows.Scan(&name, &surname, &gender, &age, &nation, &email, &phoneNumber, &workingExperience, &level_type, &pollNum, &questionNum, &questionId, &answer)
		if err != nil {
			return nil, err
		}

		user := &pb.UserExcelResp{
			Name:              name,
			Surname:           surname,
			Gender:            gender,
			Email:             email,
			Age:               age,
			Nation:            nation,
			PhoneNumber:       phoneNumber,
			WorkingExperience: workingExperience,
			LevelType:         level_type,
		}

		resultKey := name + surname + string(pollNum)
		if result, exists := resultsMap[resultKey]; exists {
			result.Answers = append(result.Answers, &pb.IncomingAnswer{
				QuestionId:  &questionId,
				AnswerPoint: &answer,
			})
		} else {
			resultsMap[resultKey] = &pb.ResultRes{
				User:    user,
				PollNum: &pollNum,
				Answers: []*pb.IncomingAnswer{
					{
						QuestionId:  &questionId,
						AnswerPoint: &answer,
					},
				},
			}
		}
	}

	var results []*pb.ResultRes
	for _, result := range resultsMap {
		results = append(results, result)
	}

	return &pb.ExcelResultsRes{Results: results}, nil
}

func (m *ResultManager) GetByIDRes(ctx context.Context, req *pb.ByIDs) (*pb.ByIDResponse, error) {
	query := `
	SELECT pa.question_id, q.num, pa.answer, p.feedbacks
	FROM poll_answers pa
	JOIN questions q ON pa.question_id = q.id
	JOIN results r ON pa.result_id = r.id
	JOIN polls p on r.poll_id = p.id
	WHERE result_id = $1
	`
	var feedbacks []byte
	rows, err := m.Conn.Query(query, req.ResultId)
	var res pb.ByIDResponse
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var questionId string
		var answer, num int32
		err := rows.Scan(&questionId, &num, &answer, &feedbacks)
		if err != nil {
			return nil, err
		}
		res.Answers = append(res.Answers, &pb.IncomingAnswer{QuestionId: &questionId, AnswerPoint: &answer, Num: &num})
	}
	var feedbackList []*pb.Feedback
	if len(feedbacks) > 0 {
		fmt.Println(string(feedbacks))
		var rawFeedbacks []map[string]interface{}
		if err := json.Unmarshal(feedbacks, &rawFeedbacks); err != nil {
			return nil, fmt.Errorf("failed to unmarshal feedbacks: %s", err.Error())
		}

		for _, rawFeedback := range rawFeedbacks {
			feedback := &pb.Feedback{}
			if from, ok := rawFeedback["from"].(float64); ok {
				fromInt := int32(from)
				feedback.From = &fromInt
			}
			if to, ok := rawFeedback["to"].(float64); ok {
				toInt := int32(to)
				feedback.To = &toInt
			}
			if text, ok := rawFeedback["text"].(string); ok {
				feedback.Text = &text
			}
			feedbackList = append(feedbackList, feedback)
		}
	}
	res.Feed = feedbackList
	return &res, nil
}
