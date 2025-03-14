package repository

import (
	"database/sql"
	"errors"

	"github.com/entonekryzhovnik/user-service/internal/model"
)

type UserRepository interface {
	CreateUser(user model.User) (int64, error) // ✅ Принимаем model.User
	GetUser(id int64) (*model.User, error)
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db: db}
}

func (r *userRepository) CreateUser(user model.User) (int64, error) { // ✅ Принимаем model.User
	result, err := r.db.Exec("INSERT INTO users (email) VALUES (?)", user.Email) // ✅ Берем email из структуры
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *userRepository) GetUser(id int64) (*model.User, error) {
	var user model.User
	err := r.db.QueryRow("SELECT id, email, created_at FROM users WHERE id = ?", id).
		Scan(&user.ID, &user.Email, &user.CreatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}
