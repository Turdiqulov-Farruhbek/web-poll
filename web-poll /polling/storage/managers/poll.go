package managers

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	pb "poll-service/genprotos"
)

type PollManager struct {
	Conn *sql.DB
}

func NewPollManager(conn *sql.DB) *PollManager {
	return &PollManager{Conn: conn}
}

func (m *PollManager) Create(ctx context.Context, poll *pb.PollCreateReq) (*pb.Void, error) {
	query := "SELECT insert_poll($1, $2, $3, $4)"
	optionsJSON, err := json.Marshal(poll.Options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal options to JSON: %w", err)
	}
	feedbackJSON, err := json.Marshal(poll.Feedbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal feedback to JSON: %w", err)
	}

	_, err = m.Conn.ExecContext(ctx, query, poll.Title, poll.Subtitle, optionsJSON, feedbackJSON)
	if err != nil {
		return nil, fmt.Errorf("failed to create poll: %w", err)
	}
	return &pb.Void{}, nil
}

func (m *PollManager) Update(ctx context.Context, poll *pb.PollUpdateReq) (*pb.Void, error) {
	// Eski pollni bazadan olish
	var existingPoll struct {
		Title     string          `db:"title"`
		Subtitle  string          `db:"subtitle"`
		Options   json.RawMessage `db:"options"`
		Feedbacks json.RawMessage `db:"feedbacks"`
	}
	err := m.Conn.QueryRowContext(ctx, "SELECT title, subtitle, options, feedbacks FROM polls WHERE id = $1", poll.Id).Scan(&existingPoll.Title, &existingPoll.Subtitle, &existingPoll.Options, &existingPoll.Feedbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to get existing poll: %w", err)
	}
	// Title yangilash yoki eski title saqlash
	title := poll.Title
	if title == nil {
		title = &existingPoll.Title
	}

	// Subtitle yangilash yoki eski subtitle saqlash
	subtitle := poll.Subtitle
	if subtitle == nil {
		subtitle = &existingPoll.Subtitle
	}

	// Options yangilash yoki eski options saqlash
	options := poll.Options
	if len(options) == 0 {
		options = []*pb.Option{}
		json.Unmarshal(existingPoll.Options, &options)
	}

	// Feedbacks yangilash yoki eski feedbacks saqlash
	feedbacks := poll.Feedbacks
	if len(feedbacks) == 0 {
		feedbacks = []*pb.Feedback{}
		json.Unmarshal(existingPoll.Feedbacks, &feedbacks)
	}

	// Yangilangan options va feedbacklarni JSON formatida saqlash
	optionsJSON, err := json.Marshal(options)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal options: %w", err)
	}

	feedbacksJSON, err := json.Marshal(feedbacks)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal feedbacks: %w", err)
	}

	// Yangilanishni amalga oshirish
	query := `
		UPDATE polls
		SET title = $1, subtitle = $2, options = $3, feedbacks = $4
		WHERE id = $5
	`
	_, err = m.Conn.ExecContext(ctx, query, title, subtitle, optionsJSON, feedbacksJSON, poll.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to update poll: %w", err)
	}

	return &pb.Void{}, nil
}

func (m *PollManager) GetByID(ctx context.Context, req *pb.ByID) (*pb.PollGetByIDRes, error) {
	query := "SELECT id, poll_num, title, subtitle, options, feedbacks FROM polls WHERE id = $1"
	row := m.Conn.QueryRowContext(ctx, query, req.Id)
	var (
		id        string
		pollNum   int32
		title     string
		subtitle  string
		options   []byte
		feedbacks []byte
	)
	
	err := row.Scan(&id, &pollNum, &title, &subtitle, &options, &feedbacks)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, fmt.Errorf("poll not found: %w", err)
		}
		return nil, fmt.Errorf("failed to get poll by id: %w", err)
	}

	// Unmarshal Options JSON
	var optionList []*pb.Option
	if len(options) > 0 {
		var rawOptions []map[string]interface{}
		if err := json.Unmarshal(options, &rawOptions); err != nil {
			return nil, fmt.Errorf("failed to unmarshal options: %s", err.Error())
		}

		for _, rawOption := range rawOptions {
			option := &pb.Option{}
			if variant, ok := rawOption["variant"].(string); ok {
				option.Variant = &variant
			}
			if ball, ok := rawOption["ball"].(float64); ok {
				ballInt := int32(ball)
				option.Ball = &ballInt
			}
			optionList = append(optionList, option)
		}
	}

	// Unmarshal Feedbacks JSON
	var feedbackList []*pb.Feedback
	if len(feedbacks) > 0 {
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

	return &pb.PollGetByIDRes{
		Id:       &id,
		PollNum:  &pollNum,
		Title:    &title,
		Subtitle: &subtitle,
		Options:  optionList,
		Feedback: feedbackList,
	}, nil
}

func (m *PollManager) Delete(ctx context.Context, req *pb.ByID) (*pb.Void, error) {
	query := "DELETE FROM polls WHERE id = $1"
	_, err := m.Conn.ExecContext(ctx, query, req.Id)
	if err != nil {
		return nil, fmt.Errorf("failed to delete poll: %w", err)
	}
	return &pb.Void{}, nil
}

// PollGetAll method with fixed JSON unmarshaling.
func (m *PollManager) GetAll(ctx context.Context, req *pb.PollGetAllReq) (*pb.PollGetAllRes, error) {
	query := "SELECT id, poll_num, title, subtitle, options, feedbacks FROM polls WHERE 1 = 1"
	var args []interface{}
	paramIndex := 1

	// Agar user_id mavjud bo'lsa, shart qo'shish
	if req.UserId != nil && *req.UserId != "" {
		query += fmt.Sprintf(" AND user_id = $%d", paramIndex)
		args = append(args, *req.UserId)
		paramIndex++
	}

	// So'rovni bajarish
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("failed to get all polls: %w", err)
	}
	defer rows.Close()

	var polls []*pb.PollGetByIDRes
	for rows.Next() {
		var (
			id        string
			pollNum   int32
			title     string
			subtitle  string
			options   []byte
			feedbacks []byte
		)

		// Ma'lumotlarni o'qib olish
		err := rows.Scan(&id, &pollNum, &title, &subtitle, &options, &feedbacks)
		if err != nil {
			return nil, fmt.Errorf("failed to scan poll: %s", err.Error())
		}

		// Optionsni JSON ob'ekti sifatida unmarshaling qilish
		var optionList []*pb.Option
		if len(options) > 0 {
			if err := json.Unmarshal(options, &optionList); err != nil {
				return nil, fmt.Errorf("failed to unmarshal options: %s", err.Error())
			}
		}

		// Feedbacksni JSON ob'ekti sifatida unmarshaling qilish
		var feedbackList []*pb.Feedback
		if len(feedbacks) > 0 {
			if err := json.Unmarshal(feedbacks, &feedbackList); err != nil {
				return nil, fmt.Errorf("failed to unmarshal feedbacks: %s", err.Error())
			}
		}

		// Har bir poll uchun javobni yig'ish
		polls = append(polls, &pb.PollGetByIDRes{
			Id:       &id,
			PollNum:  &pollNum,
			Title:    &title,
			Subtitle: &subtitle,
			Options:  optionList,
			Feedback: feedbackList,
		})
	}

	// Xatoliklarni tekshirish
	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %s", err.Error())
	}

	// Javobni qaytarish
	return &pb.PollGetAllRes{
		Poll: polls,
	}, nil
}

