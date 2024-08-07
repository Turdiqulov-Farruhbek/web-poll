package managers

import (
	"context"
	"database/sql"
	"fmt"
	pb "poll-service/genprotos"

	"github.com/google/uuid"
)

type QuestionManager struct {
	Conn *sql.DB
}

func NewQuestionManager(conn *sql.DB) *QuestionManager {
	return &QuestionManager{Conn: conn}
}

func (m *QuestionManager) Create(ctx context.Context, question *pb.QuestionCreateReq) (*pb.Void, error) {
	query := "INSERT INTO questions (id, num, content, poll_id) VALUES ($1, $2, $3, $4)"
	_, err := m.Conn.Exec(query, uuid.NewString(), question.Num, question.Content, question.PollId)
	if err != nil {
		return nil, fmt.Errorf("failed to create question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) GetByID(ctx context.Context, req *pb.ByID) (*pb.QuestionGetByIDRes, error) {
	query := "SELECT id, num, content, poll_id FROM questions WHERE id = $1"
	row := m.Conn.QueryRowContext(ctx, query, req.Id)

	var res pb.QuestionGetByIDRes
	err := row.Scan(&res.Id, &res.Num, &res.Content, &res.PollId)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("question not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get question by id: %w", err)
	}
	return &res, nil
}

func (m *QuestionManager) Update(ctx context.Context, question *pb.QuestionUpdateReq) (*pb.Void, error) {
	query := "UPDATE questions SET num = $1, content = $2, poll_id = $3 WHERE id = $4"
	_, err := m.Conn.ExecContext(ctx, query, question.Num, question.Content, question.PollId, question.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM questions WHERE id = $1"
	_, err := m.Conn.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete question: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *QuestionManager) GetAll(ctx context.Context, req *pb.QuestionGetAllReq) (*pb.QuestionGetAllRes, error) {
	query := "SELECT id, num, content, poll_id FROM questions"
	var args []interface{}
	paramIndex := 1
	if req.PollId != "" {
		query += fmt.Sprintf(" WHERE poll_id = $%d", paramIndex)
		args = append(args, req.PollId)
		paramIndex++
	}
	if req.Pagination.Limit != 0 {
		query += fmt.Sprintf(" LIMIT $%d", paramIndex)
		args = append(args, req.Pagination.Limit)
		paramIndex++
	}
	if req.Pagination.Offset != 0 {
		query += fmt.Sprintf(" OFFSET $%d", paramIndex)
		args = append(args, req.Pagination.Offset)
		paramIndex++
	}

	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all questions: %w", err)
	}
	defer rows.Close()

	var questions []*pb.Question
	for rows.Next() {
		var question pb.Question
		err := rows.Scan(&question.Id, &question.Num, &question.Content, &question.PollId)
		if err != nil {
			return nil, fmt.Errorf("failed to scan question: %w", err)
		}
		questions = append(questions, &question)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return &pb.QuestionGetAllRes{
		Question: questions,
	}, nil
}
