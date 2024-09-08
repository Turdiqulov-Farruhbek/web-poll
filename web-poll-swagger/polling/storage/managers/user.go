package managers

import (
	"context"
	"database/sql"
	"errors"
	pb "poll-service/genprotos"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
)

type UserManager struct {
	Conn *sql.DB
}

func NewUserManager(conn *sql.DB) *UserManager {
	return &UserManager{Conn: conn}
}

func (m *UserManager) Register(context context.Context, req *pb.RegisterReq) (*pb.Void, error) {
	var gender string
	if req.Gender == 0 {
		gender = "male"
	} else {
		gender = "female"
	}

	query := "INSERT INTO users (id, name, surname, gender, email, password, phone_number, working_experience, level_type) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)"
	_, err := m.Conn.Exec(query, uuid.NewString(), req.Name, req.Surname, gender, req.Email, req.Password, req.PhoneNumber, req.WorkingExperience, req.LevelType)
	return nil, err
}

func (m *UserManager) ConfirmUser(context context.Context, req *pb.ConfirmUserReq) (*pb.Void, error) {
	query := "UPDATE users SET is_confirmed = true, confirmed_at = $1 WHERE email = $2"
	_, err := m.Conn.Exec(query, time.Now(), req.Email)
	return nil, err
}

func (m *UserManager) Profile(context context.Context, req *pb.GetProfileReq) (*pb.GetProfileResp, error) {
	query := "SELECT id, email, password, role, is_confirmed FROM users WHERE email = $1"
	row := m.Conn.QueryRow(query, req.Email)
	var user pb.GetProfileResp
	err := row.Scan(&user.Id, &user.Email, &user.Password, &user.Role, &user.IsConfirmed)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (m *UserManager) UpdatePassword(context context.Context, req *pb.UpdatePasswordReq) (*pb.Void, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	query := "UPDATE users SET password = $1 WHERE email = $2"
	_, err = m.Conn.Exec(query, string(hashedPassword), req.Email)
	return nil, err
}

func (m *UserManager) IsEmailExists(context context.Context, email *pb.IsEmailExistsReq) (*pb.IsEmailExistsResp, error) {
	query := "SELECT COUNT(*) FROM users WHERE email = $1"
	var count int
	err := m.Conn.QueryRow(query, email.Email).Scan(&count)
	if err != nil {
		return nil, errors.New("server error: " + err.Error())
	}
	if count > 0 {
		return &pb.IsEmailExistsResp{Exists: true}, nil
	}
	return &pb.IsEmailExistsResp{Exists: false}, nil
}

func (m *UserManager) GetByID(context context.Context, id *pb.GetProfileByIdReq) (*pb.GetProfileByIdResp, error) {
	query := "SELECT id, email, role FROM users WHERE id = $1"
	user := &pb.GetProfileByIdResp{}
	err := m.Conn.QueryRow(query, id.Id).Scan(&user.Id, &user, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}

func (m *UserManager) GetByEmail(context context.Context, id *pb.ByEmail) (*pb.GetProfileByIdResp, error) {
	query := "SELECT id, email, role FROM users WHERE email = $1"
	user := &pb.GetProfileByIdResp{}
	err := m.Conn.QueryRow(query, id.Email).Scan(&user.Id, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return user, nil
}
