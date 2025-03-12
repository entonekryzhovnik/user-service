package repository

import (
	"database/sql"
	"errors"
	"user-service/gen/go/userpb"
)

type UserRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{db: db}
}

func (r *UserRepository) CreateUser(email string) (int64, error) {
	result, err := r.db.Exec("INSERT INTO users (email) VALUES (?)", email)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *UserRepository) GetUser(id int64) (*userpb.User, error) {
	var user userpb.User
	err := r.db.QueryRow("SELECT id, email, created_at FROM users WHERE id = ?", id).
		Scan(&user.Id, &user.Email, &user.CreateAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
