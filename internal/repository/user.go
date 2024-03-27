package repository

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
	"github.com/xxlifestyle/auth_service/internal/entity"
)

type UserRepository interface {
	Create(user *entity.User) error
	FindById(id string) (*entity.User, error)
}

type userPostgresRepository struct {
	db *pgx.Conn
}

func (ur userPostgresRepository) Create(user *entity.User) error {
	sql := fmt.Sprintf("INSERT INTO \"user\" (username, password_hash, birthday_date, email, is_email_confirmed) VALUES ('%s','%s','%s','%s',%t)", user.Username, user.PasswordHash, user.BirthdayDate, user.Email, user.IsEmailConfirmed)
	fmt.Print(sql)
	_, err := ur.db.Query(context.Background(), sql)

	if err != nil {
		return err
	}
	return nil
}
func (ur userPostgresRepository) FindById(id string) (*entity.User, error) {
	_, err := ur.db.Query(context.Background(), "SELECT id,username FROM user WHERE id = $1", id)
	if err != nil {
		return nil, err
	}
	return &entity.User{}, nil
}

func NewUserPostgresRepository(db *pgx.Conn) UserRepository {
	return userPostgresRepository{
		db: db,
	}
}
