package managers

import (
	"context"
	"database/sql"

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
	return &pb.CreateResultRes{ResultId: id}, nil
}

func (m *ResultManager) SavePollAnswer(ctx context.Context, req *pb.SavePollAnswerReq) (*pb.Void, error) {

	query := "INSERT INTO poll_answers (id, result_id, question_num, answer) VALUES ($1, $2, $3, $4)"
	_, err := m.Conn.ExecContext(ctx, query, uuid.NewString(), req.ResultId, req.QuestionNum, req.Answer)
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func (m *ResultManager) GetResultsInExcel(ctx context.Context, req *pb.Void) (*pb.ExcelResultsRes, error) {
	query := `
		SELECT 
			u.name, u.surname, u.gender, u.email, u.phone_number, u.working_experience, u.level_type,
			p.poll_num,
			pa.question_num, pa.answer
		FROM 
			results r
		JOIN 
			users u ON r.user_id = u.id
		JOIN 
			poll_answers pa ON r.id = pa.result_id
		JOIN 
			polls p ON r.poll_id = p.id
		ORDER BY 
			u.id, p.poll_num, pa.question_num;
	`

	rows, err := m.Conn.QueryContext(ctx, query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	resultsMap := make(map[string]*pb.ResultRes)
	for rows.Next() {
		var name, surname, gender, email, phoneNumber, level_type string
		var workingExperience int32
		var pollNum, questionNum, answer int32

		err := rows.Scan(&name, &surname, &gender, &email, &phoneNumber, &workingExperience, &level_type, &pollNum, &questionNum, &answer)
		if err != nil {
			return nil, err
		}

		user := &pb.UserExcelResp{
			Name:              name,
			Surname:           surname,
			Gender:            gender,
			Email:             email,
			PhoneNumber:       phoneNumber,
			WorkingExperience: workingExperience,
			LevelType:             level_type,
		}

		resultKey := name + surname + string(pollNum)
		if result, exists := resultsMap[resultKey]; exists {
			result.Answers = append(result.Answers, &pb.IncomingAnswer{
				Num:         questionNum,
				AnswerPoint: answer,
			})
		} else {
			resultsMap[resultKey] = &pb.ResultRes{
				User:    user,
				PollNum: pollNum,
				Answers: []*pb.IncomingAnswer{
					{
						Num:         questionNum,
						AnswerPoint: answer,
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
