package repository

import (
	"context"

	"github.com/dyaksa/boilerplate-go-clean-arhictecture/domain"
	"github.com/dyaksa/boilerplate-go-clean-arhictecture/infrastructure/pqsql"
)

type useRepository struct {
	database pqsql.Client
}

func NewUserRepository(db pqsql.Client) domain.UserRepository {
	return &useRepository{
		database: db,
	}
}

func (ur *useRepository) GetUserByEmail(ctx context.Context, email_bidx string) (*domain.User, error) {
	tx, err := ur.database.Begin()
	if err != nil {
		return nil, err
	}
	query := `SELECT id, name, email FROM users WHERE email_bidx = $1`
	res, err := ur.database.Database().QueryRow(ctx, tx, query, nil, &domain.User{}, email_bidx)
	if err != nil {
		return nil, err
	}

	return res.(*domain.User), nil
}
