package user

import (
	"context"
)

// Implementation of the repository in this service.
type userService struct {
	userRepository UserRepository
}

// Create a new 'service' or 'use-case' for 'User' entity.
func NewUserService(repo UserRepository) UserService {
	return &userService{
		userRepository: repo,
	}
}

func (service *userService) GetUsers(ctx context.Context) ([]*User, error) {
	return service.userRepository.GetUsers(ctx)
}

func (service *userService) GetUser(ctx context.Context, userID string) (*User, error) {
	return service.userRepository.GetUser(ctx, userID)
}

func (service *userService) CreateUser(ctx context.Context, user *User) (*User, error) {
	return service.userRepository.CreateUser(ctx, user)
}

func (service *userService) UpdateUser(ctx context.Context, userID string, user *User) (*User, error) {
	return service.userRepository.UpdateUser(ctx, userID, user)
}

func (service *userService) DeleteUser(ctx context.Context, userID string) error {
	return service.userRepository.DeleteUser(ctx, userID)
}
