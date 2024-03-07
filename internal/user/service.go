package user

import (
	"context"
	"github.com/go_fundaments/internal/domain"
	"log"
)

type (
	Service interface {
		Create(ctx context.Context, firstName, lastName, email string) (*domain.User, error)
		GetAll(ctx context.Context) ([]domain.User, error)
		Get(ctx context.Context, id uint64) (*domain.User, error)
		Update(ctx context.Context, id uint64, firstName, lastName, email *string) error
	}

	service struct {
		log  *log.Logger
		repo Repository
	}
)

func NewService(l *log.Logger, repo Repository) Service {
	return &service{
		log:  l,
		repo: repo,
	}
}

// Method create
func (s service) Create(ctx context.Context, first_name, last_name, email string) (*domain.User, error) {

	user := &domain.User{
		FirstName: first_name,
		LastName:  last_name,
		Email:     email,
	}

	if err := s.repo.Create(ctx, user); err != nil {
		return nil, err
	}
	return user, nil
}

// Method GetAll
func (s service) GetAll(ctx context.Context) ([]domain.User, error) {
	users, err := s.repo.GetAll(ctx)
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (s service) Get(ctx context.Context, id uint64) (*domain.User, error) {
	return s.repo.Get(ctx, id)
}

func (s service) Update(ctx context.Context, id uint64, firstName, lastName, email *string) error {
	if err := s.repo.Update(ctx, id, firstName, lastName, email); err != nil {
		return err
	}
	return nil
}