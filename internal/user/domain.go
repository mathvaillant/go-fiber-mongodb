package user

import "context"

type User struct {
	ID    string `json:"id,omitempty" bson:"_id,omitempty"`
	Name  string `json:"name" bson:"name"`
	Email string `json:"email" bson:"email"`
}

type UserRepository interface {
	GetUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, userID string, user *User) (*User, error)
	DeleteUser(ctx context.Context, userID string) error
}

type UserService interface {
	GetUsers(ctx context.Context) ([]*User, error)
	GetUser(ctx context.Context, userID string) (*User, error)
	CreateUser(ctx context.Context, user *User) (*User, error)
	UpdateUser(ctx context.Context, userID string, user *User) (*User, error)
	DeleteUser(ctx context.Context, userID string) error
}
