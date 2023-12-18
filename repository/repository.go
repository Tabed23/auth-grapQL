package repository

import (
	"context"

	"github.com/tabed23/auth-graphql/graph/model"
)

type UserRepository interface {
	UserCreate(ctx context.Context, input model.NewUser) (*model.User, error)
	UserGetByID(ctx context.Context, id string) (*model.User, error)
	UserGetByEmail(ctx context.Context, email string) (*model.User, error)
	UserDelete(ctx context.Context, email string) error
	UserUpdate(ctx context.Context, email string, user *model.NewUser) (string, error)
}
