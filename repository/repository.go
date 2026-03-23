package repository

import (
	"context"

	"github.com/jmoiron/sqlx"
	m "github.com/venwex/weatherbot/models"
)

type Repository interface {
	GetUserCity(ctx context.Context, userID int64) (string, error)
	CreateUser(ctx context.Context, userID int64) error
	UpdateUserCity(ctx context.Context, userID int64, city string) error
	GetUser(ctx context.Context, userID int64) (*m.User, error)
}

type UserRepo struct {
	db *sqlx.DB
}

func NewUserRepo(db *sqlx.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (r *UserRepo) GetUserCity(ctx context.Context, userID int64) (string, error) {
	var city string
	if err := r.db.GetContext(ctx, &city, "select coalesce(city, '') from users where id = $1", userID); err != nil {
		return "", err
	}

	return city, nil
}

func (r *UserRepo) CreateUser(ctx context.Context, userID int64) error {
	if _, err := r.db.ExecContext(ctx, "insert into users (id) values ($1)", userID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) UpdateUserCity(ctx context.Context, userID int64, city string) error {
	if _, err := r.db.ExecContext(ctx, "update users set city = $1 where id = $2;", city, userID); err != nil {
		return err
	}

	return nil
}

func (r *UserRepo) GetUser(ctx context.Context, userID int64) (*m.User, error) {
	var user m.User
	if err := r.db.GetContext(ctx, &user, "select id, coalesce(city, '') as city, created_at from users where id = $1", userID); err != nil {

		return &m.User{}, err
	}

	return &user, nil
}


