package storage

import (
	"database/sql"
	"fmt"

	"poll-service/config"
	"poll-service/storage/managers"

	_ "github.com/lib/pq"
)

type Storage struct {
	PgClient  *sql.DB
	UserS     *managers.UserManager
	QuestionS *managers.QuestionManager
	PollS     *managers.PollManager
	ResultS   *managers.ResultManager
}

func NewPostgresStorage(config config.Config) (*Storage, error) {
	// #################    POSTGRESQL CONNECTION     ###################### //
	conn := fmt.Sprintf("host=%s user=%s dbname=%s password=%s port=%d sslmode=disable",
		config.DB_HOST, config.DB_USER, config.DB_NAME, config.DB_PASSWORD, config.DB_PORT)
	db, err := sql.Open("postgres", conn)
	fmt.Println(conn)
	if err != nil {
		return nil, err
	}
	if err = db.Ping(); err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to database pgsql!")

	um := managers.NewUserManager(db)
	pm := managers.NewPollManager(db)
	qm := managers.NewQuestionManager(db)
	rm := managers.NewResultManager(db)

	return &Storage{
		PgClient:  db,
		UserS:     um,
		PollS:     pm,
		QuestionS: qm,
		ResultS:   rm,
	}, nil
}

func (s *Storage) Poll() PollI {
	if s.PollS == nil {
		s.PollS = managers.NewPollManager(s.PgClient)
	}
	return s.PollS
}

func (s *Storage) User() UserI {
	if s.UserS == nil {
		s.UserS = managers.NewUserManager(s.PgClient)
	}
	return s.UserS
}

func (s *Storage) Question() QuestionI {
	if s.QuestionS == nil {
		s.QuestionS = managers.NewQuestionManager(s.PgClient)
	}
	return s.QuestionS
}

func (s *Storage) Result() ResultI {
	if s.ResultS == nil {
		s.ResultS = managers.NewResultManager(s.PgClient)
	}
	return s.ResultS
}
