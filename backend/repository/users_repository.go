package repository

import (
	"asset-tracking-system/dto"
	"asset-tracking-system/entity"
	"context"
	"database/sql"
)

type UserRepository interface {
}

type userRepository struct {
	db *sql.DB
}

func NewUserRepository(db *sql.DB) *userRepository {
	return &userRepository{
		db: db,
	}
}

func (r *userRepository) AddUser(newUser *dto.UserRegisterRequest, pwd []byte) (*entity.User, error) {
	registeredUser := entity.User{}

	return &registeredUser, nil
}

func (r *userRepository) UpdateUser(ctx context.Context) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) LoginUser(userEmail string) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) IsUserIdExist(id int) bool {
	return false
}

func (r *userRepository) IsUserExist(email string) bool {
	return false
}

func (r *userRepository) UserList() (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) GetUserbyId(id int) (*entity.User, error) {
	return nil, nil
}

func (r *userRepository) ResetPassword(email string, pwd []string) error {
	return nil
}

func (r *userRepository) DeactivateUser(id int) error {
	return nil
}
