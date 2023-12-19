package store

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/tabed23/auth-graphql/graph/model"
	"github.com/tabed23/auth-graphql/utils"
	"gorm.io/gorm"

	"github.com/tabed23/auth-graphql/repository"
)

type UserStore struct {
	store *gorm.DB
}

func NewUserStore(db *gorm.DB) repository.UserRepository {
	return &UserStore{store: db}
}

// UserCreate implements repository.UserRepository.
func (u *UserStore) UserCreate(ctx context.Context, input model.NewUser) (*model.User, error) {
	input.Password = utils.HashPassword(input.Password)
	usr := model.User{
		ID:       uuid.New().String(),
		Name:     input.Name,
		Email:    input.Email,
		Password: input.Password,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	tx := u.store.Create(usr).Error
	if tx != nil {
		log.Fatal(tx.Error())
		return nil, tx
	}

	return &usr, nil

}

// UserGetByEmail implements repository.UserRepository.
func (u *UserStore) UserGetByEmail(ctx context.Context, email string) (*model.User, error) {
	usr := model.User{}
	if err := u.store.Where("email = ?", email).Find(&usr).Error; err != nil {
		return nil, err
	}

	return &usr, nil
}

// UserGetByID implements repository.UserRepository.
func (u *UserStore) UserGetByID(ctx context.Context, id string) (*model.User, error) {
	usr := model.User{}
	if err := u.store.Where("id = ?", id).Find(&usr).Error; err != nil {
		return nil, err
	}

	return &usr, nil
}

// UserDelete implements repository.UserRepository.
func (u *UserStore) UserDelete(ctx context.Context, email string) error {
	user, err := u.UserGetByEmail(ctx, email)
	if err != nil {
		return err
	}

	if err := u.store.Delete(&model.User{},"email = ?", user.Email).Error; err != nil {
		return err
	}

	return nil
}

// UserUpdate implements repository.UserRepository.
func (u *UserStore) UserUpdate(ctx context.Context, email string, user *model.NewUser) (string, error) {
	updateTime :=  time.Now()
	
	usr := model.NewUser{
		Name:     user.Name,
		Email:    user.Email,
		Password: user.Password,
		UpdatedAt: &updateTime,
	}
	foundUser, err := u.UserGetByEmail(ctx, email)
	if err != nil {
		return "", err
	}

	if err := u.store.Where("id = ?", foundUser.ID).Updates(&usr).Error; err != nil {
		return fmt.Sprintf("Cannot update the user with email %s", email), nil
	}

	return fmt.Sprintf("User with email %s as been updated", email), nil
}
