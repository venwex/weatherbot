package service

import (
	"context"
	"database/sql"
	"errors"

	m "github.com/venwex/weatherbot/models"
	"github.com/venwex/weatherbot/repository"
)

type Service struct {
	repo repository.Repository
}

func NewService(repo repository.Repository) *Service {
	return &Service{repo: repo}
}

func (s *Service) GetUserCity(ctx context.Context, userID int64) (string, error) {
	return s.repo.GetUserCity(ctx, userID)
}

func (s *Service) CreateUser(ctx context.Context, userID int64) error {
	return s.repo.CreateUser(ctx, userID)
}

func (s *Service) UpdateUserCity(ctx context.Context, userID int64, city string) error {
	return s.repo.UpdateUserCity(ctx, userID, city)
}

func (s *Service) GetUser(ctx context.Context, id int64) (*m.User, error) {
	user, err := s.repo.GetUser(ctx, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return user, nil
}
